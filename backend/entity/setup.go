package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("bookorder.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&BookOrder{}, &BookType{}, &Company{}, &OrderStatus{}, &Member{})

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Member{}).Create(&Member{
    Name: "บุญฑิตา ปวงสันเทียะ",
    Email: "boontita@mail.com",
    Password: string(password),
	})

	done := OrderStatus{
		Status: "สั่งซื้อสำเร็จ",
	}
	db.Model(&OrderStatus{}).Create(&done)

	waitforapproval := OrderStatus{
		Status: "รอการอนุมัติ",
	}
	db.Model(&OrderStatus{}).Create(&waitforapproval)

	request := OrderStatus{
		Status: "เสนอสั่งซื้อ",
	}
	db.Model(&OrderStatus{}).Create(&request)

	approval := OrderStatus{
		Status: "อนุมัติ",
	}
	db.Model(&OrderStatus{}).Create(&approval)

	A := Company{
		NameThai:    "บริษัท เอ จำกัด",
		NameEng:     "A Company Limited",
		Address:     "12 ถนน สาม ตำบอลสังคม อำเภอเมือง จังหวัดสมุทรปราการ 10270",
		PhoneNumber: "0123456789",
		Email:       "a4586sa@mail.com",
		Website:     "A Co.Ltd.com",
	}
	db.Model(&Company{}).Create(&A)

	B := Company{
		NameThai:    "บริษัท สมชายการพิมพ์ จำกัด",
		NameEng:     "Somchai Company Limited",
		Address:     "12 ถนน สาม ตำบอลสังคม อำเภอเมือง จังหวัดสมุทรปราการ 10270",
		PhoneNumber: "0123456789",
		Email:       "somchai_typing@mail.com",
		Website:     "SomChai.com",
	}
	db.Model(&Company{}).Create(&B)

	C := Company{
		NameThai:    "สำนักพิมพ์ดีต่อใจ",
		NameEng:     "deetorjaibooks",
		Address:     "12 ถนน สาม ตำบอลสังคม อำเภอเมือง จังหวัดสมุทรปราการ 10270",
		PhoneNumber: "0123456789",
		Email:       "deetorjaibooks@mail.com",
		Website:     "deetorjaibooks.com",
	}
	db.Model(&Company{}).Create(&C)

	novel := BookType{
		Type: "นิยาย",
	}
	db.Model(&BookType{}).Create(&novel)

	documentary := BookType{
		Type: "สารคดี",
	}
	db.Model(&BookType{}).Create(&documentary)

	chidren := BookType{
		Type: "หนังสือสำหรับเด็กและเยาวชน",
	}
	db.Model(&BookType{}).Create(&chidren)

	treatise := BookType{
		Type: "ตำรา",
	}
	db.Model(&BookType{}).Create(&treatise)

	order1 := BookOrder{
		BookTitle:   "คณิตศาสตร์",
		Author:      "หญิงสาม ใจดี",
		BookType:    treatise,
		Company:     B,
		OrderAmount: 5,
		Price:       2052.50,
		OrderStatus: request,
	}
	db.Model(&BookOrder{}).Create(&order1)
}
