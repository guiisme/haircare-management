package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	HairdresserID uint   `json:"hairdresser_id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Description   string `json:"description" gorm:"type:varchar(300)"`
}
