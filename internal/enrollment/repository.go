package enrollment

import (
	"log"

	"github.com/SanGameDev/GoWebPractice/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, log *log.Logger) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (repo *repo) Create(enroll *domain.Enrollment) error {

	if err := repo.db.Create(enroll).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}

	repo.log.Println("Enrollment created")
	return nil
}
