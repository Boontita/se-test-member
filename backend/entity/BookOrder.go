package entity

import (
	"time"

	"gorm.io/gorm"

)
type Company struct {
	gorm.Model
	NameThai string
	NameEng  string
	Address	string
	PhoneNumber	string
	Email     string
	Website		string

	BookOrder []BookOrder `gorm:"foreignKey:CompanyID"`
}

type OrderStatus struct {
	gorm.Model
	Status	string

	BookOrder []BookOrder `gorm:"foreignKey:OrderStatusID"`
}

type BookOrder struct {
	gorm.Model
	BookTitle   string
	Author      string
	OrderAmount uint
	Price       float32
	OrderDate   time.Time

	//Company ทำหน้าที่เป็น FK
	CompanyID *uint
	Company   Company

	//Company ทำหน้าที่เป็น FK
	OrderStatusID *uint
	OrderStatus   OrderStatus

	//BookType ทำหน้าที่เป็น FK
	BookTypeID *uint
	BookType   BookType

	//Librarian
	MemberID *uint
	Member		Member
	
}
