package user

import "log"

type Service interface {
	Create(fistName, lastName, email, phone string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Create(fistName, lastName, email, phone string) error {
	log.Println("Create user")
	s.repository.Create(&User{})
	return nil
}
