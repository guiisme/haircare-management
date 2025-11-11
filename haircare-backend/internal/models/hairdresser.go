package models

import "gorm.io/gorm"

type Hairdresser struct {
    gorm.Model
	//ID	   uint   `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name" gorm:"not null"`
    Email    string `json:"email" gorm:"unique;not null"`
    CPF      string `json:"cpf" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
    Clients  []Client
}
