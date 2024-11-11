package enrollment

import (
	"errors"
	"log"

	"github.com/SanGameDev/GoWebPractice/internal/course"
	"github.com/SanGameDev/GoWebPractice/internal/domain"
	"github.com/SanGameDev/GoWebPractice/internal/user"
)

type (
	Service interface {
		Create(courseID, userID string) (*domain.Enrollment, error)
	}
	service struct {
		log       *log.Logger
		userSrv   user.Service
		courseSrv course.Service
		repo      Repository
	}
)

func NewService(log *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log:       log,
		userSrv:   userSrv,
		courseSrv: courseSrv,
		repo:      repo,
	}
}

func (s *service) Create(userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "pending",
	}

	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		return nil, errors.New("user not found")
	}

	if _, err := s.courseSrv.Get(enroll.CourseID); err != nil {
		return nil, errors.New("course not found")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return enroll, nil
}
