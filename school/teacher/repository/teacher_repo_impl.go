package repository

import (
	"fmt"
	"school/model"
	"school/teacher"

	"github.com/jinzhu/gorm"
)

type TeacherRepoImpl struct {
	DB *gorm.DB
}

func CreateTeacherRepo(DB *gorm.DB) teacher.TeacherRepo {
	return &TeacherRepoImpl{DB}
}

func (e *TeacherRepoImpl) ViewAll() (*[]model.Teacher, error) {
	var teacher []model.Teacher
	err := e.DB.Find(&teacher).Error
	if err != nil {
		fmt.Println("[TeacherRepoImpl.ViewAll] Error When Execute Query")
		return nil, fmt.Errorf("Opps Server Somting Wrong")
	}
	return &teacher, nil
}

func (e *TeacherRepoImpl) ViewById(id int) (*model.Teacher, error) {
	var teacher = new(model.Teacher)
	err := e.DB.Table("teacher").Where("ID = ?", id).First(&teacher).Error
	if err != nil {
		fmt.Println("[TeacherRepoImpl.ViewById] Error When Execute Query")
		return nil, fmt.Errorf("Teacher id is not exist")
	}
	return teacher, nil
}

func (e *TeacherRepoImpl) ViewByName(name string) (*model.Teacher, error) {
	var teacher = new(model.Teacher)
	err := e.DB.Table("teacher").Where("FirstName = ?", name).First(&teacher).Error
	if err != nil {
		fmt.Println("[TeacherRepoImpl.ViewById] Error When Execute Query")
		return nil, fmt.Errorf("Teacher name is not exist")
	}
	return teacher, nil
}

func (e *TeacherRepoImpl) Insert(teacher *model.Teacher)(*model.Teacher, error) {
	err := e.DB.Save(teacher).Error
	if err != nil {
		fmt.Println("[TeacherRepoImpl.Insert] Error When Execute Query")
		return nil, fmt.Errorf("inser data teacher is failed")
	}
	return teacher, nil
}

func (e *TeacherRepoImpl) Update(id int, teacher *model.Teacher)(*model.Teacher, error) {
	var newTeacher = new(model.Teacher)
	err := e.DB.Table("teacher").Where("ID = ?", id).First(&newTeacher).Update(&teacher).Error
	if err != nil {
		fmt.Println("[TeacherRepoImpl.Update] Error When Execute Query")
		return nil, fmt.Errorf("Update data teacher is failed")
	}
	return newTeacher, nil
}

func (e *TeacherRepoImpl) Delete(id int) error {
	var teacher = model.Teacher{}
	err := e.DB.Table("teacher").Where("ID = ?", id).Delete(&teacher).Error
	if err != nil {
		fmt.Println("[TeacherRepoImpl.Delete] Error When Execute Query")
		return fmt.Errorf("Delete data teacher is failed")
	}
	return nil
}