package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GORMModel
	UserID uint `json:"user_id"`
	// If the Title field is empty, the validation will fail and return an error message "Title is required".
	//validate:something-output when fail
	Title       string `json:"title" validate:"required-Title is required"`
	Description string `json:"description" validate:"required-Description is required"`
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Product) BeforeDelete(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
