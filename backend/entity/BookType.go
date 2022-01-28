package entity

import (
	"gorm.io/gorm"

)

type BookType struct {
	gorm.Model

	Type string 

	//BookInformation []BookInformation `gorm:"foreignKey:BookTypeID"`
	BookOrder []BookOrder `gorm:"foreignKey:BookTypeID"`	
}

type Member struct {
	gorm.Model
	Name string
	Email string
	Password string
	BookOrder []BookOrder `gorm:"foreignKey:MemberID"`
}