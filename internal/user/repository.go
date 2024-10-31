package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	fmt.Println("user created" + user.FirstName)
	return nil
}
