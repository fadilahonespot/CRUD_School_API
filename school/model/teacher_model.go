package model

import "github.com/jinzhu/gorm"

type Teacher struct {
	gorm.Model
	FirstName string `gorm:"column:first_name; type:varchar(50); not null"`
	LastName  string `gorm:"column:last_name; type:varchar(50); not null"`
	Email     string `gorm:"column:email; unique; not null"`
}

func (e Teacher) TableName() string {
	return "teacher"
}
