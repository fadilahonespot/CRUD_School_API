package repository

import (
	"school/model"
	"school/subject"
	"fmt"

	"github.com/jinzhu/gorm"
)

type SubjectRepoImpl struct {
	DB *gorm.DB
}

func CreateSubjectRepo(DB *gorm.DB) subject.SubjectRepo {
	return &SubjectRepoImpl{DB}
}

func (e *SubjectRepoImpl) ViewAll() (*[]model.Subject, error) {
	var subject []model.Subject
	err := e.DB.Find(&subject).Error
	if err != nil {
		fmt.Println("[SubjectRepoImpl.ViewAll] Error When Execute Query")
		return nil, fmt.Errorf("Opps Server Somting Wrong")
	}
	return &subject, nil
}

func (e *SubjectRepoImpl) ViewByid(id int) (*model.Subject, error) {
	var subject = new(model.Subject)
	err := e.DB.Table("subject").Where("ID = ?", id).First(&subject).Error
	if err != nil {
		fmt.Println("[SubjectRepoImpl.ViewById] Error When Execute Query")
		return nil, fmt.Errorf("id subject is not exist")
	}
	return subject, nil
}

func (e *SubjectRepoImpl) Insert(subject *model.Subject) (*model.Subject, error) {
	err := e.DB.Save(&subject).Error
	if err != nil {
		fmt.Println("[SubjectRepoImpl.Insert] Error When Execute Query")
		return nil, fmt.Errorf("failed insert data subject")
	}
	return subject, nil
}

func (e *SubjectRepoImpl) Update(id int, subject *model.Subject) (*model.Subject, error) {
	var newSubject = new(model.Subject)
	err := e.DB.Table("subject").Where("ID = ?", id).First(&newSubject).Update(&subject).Error
	if err != nil {
		fmt.Println("[SubjectRepoImpl.Update] Error When Execute Query")
		return nil, fmt.Errorf("failed update data subject")
	}
	return newSubject, nil
}

func (e *SubjectRepoImpl) Delete(id int) error {
	var subject = model.Subject{}
	err := e.DB.Table("subject").Where("ID = ?", id).Delete(subject).Error
	if err != nil {
		fmt.Println("[SubjectRepoImpl.Delete] Error When Execute Query")
		return fmt.Errorf("failed delete data subject")
	}
	return nil
}