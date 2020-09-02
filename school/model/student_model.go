package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	FirsName string `gorm:"column:firs_name; type:varchar(50); not null"`
	LastName string `gorm:"column:last_name; type:varchar(50); not null"`
	Email    string `gorm:"column:email; unique; not null"`
}

func (e Student) TableName() string {
	return "student"
}
