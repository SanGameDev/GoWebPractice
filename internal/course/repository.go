package course

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SanGameDev/GoWebPractice/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		Delete(id string) error
		Update(id string, name *string, startDate, endDate *time.Time) error
		Count(filters Filters) (int, error)
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(course *domain.Course) error {
	if err := r.db.Create(course).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("course created with id: ", course.ID)
	return nil
}

func (r *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var c []domain.Course

	tx := r.db.Model(&c)
	tx = tx.Offset(offset).Limit(limit)
	result := tx.Order("created_at desc").Find(&c)

	if result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

func (r *repo) Get(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}

	if err := r.db.First(&course).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *repo) Delete(id string) error {
	course := domain.Course{ID: id}

	if err := r.db.Delete(&course).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(id string, name *string, startDate, endDate *time.Time) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if startDate != nil {
		values["start_date"] = *startDate
	}

	if endDate != nil {
		values["end_date"] = *endDate
	}

	if err := r.db.Model(&domain.Course{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("LOWER(name) LIKE ?", filters.Name)
	}

	return tx
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(&domain.Course{})
	tx = applyFilters(tx, filters)

	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
