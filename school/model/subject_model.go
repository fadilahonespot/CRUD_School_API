package model

import "github.com/jinzhu/gorm"

type Subject struct {
	gorm.Model
	SubjectName string `gorm: "column: subject_name; not null"`
}

func (e Subject) TableName() string {
	return "subject"
}