package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"school/model"
	"school/student"
)

type StudentRepoImpl struct {
	DB *gorm.DB
}

func CreateStudentRepo(DB *gorm.DB) student.StudentRepo {
	return &StudentRepoImpl{DB}
}

func (e *StudentRepoImpl) ViewAll() (*[]model.Student, error) {
	var student []model.Student
	err := e.DB.Find(&student).Error
	if err != nil {
		fmt.Println("[StudentRepoImpl.ViewAll] Error When Execute Query")
		return nil, fmt.Errorf("Opps Server Somting Wrong")
	}
	return &student, nil
}

func (e *StudentRepoImpl) ViewById(id int) (*model.Student, error) {
	var student = new(model.Student)
	err := e.DB.Table("student").Where("ID = ?", id).First(&student).Error
	if err != nil {
		fmt.Println("[StudentRepoImpl.ViewById] Error when execute query")
		return nil, fmt.Errorf("id does not exist")
	}
	return student, nil
}

func (e *StudentRepoImpl) ViewByName(name string) (*model.Student, error) {
	var student = new(model.Student)
	err := e.DB.Table("student").Where("FirsName = ?", name).First(&student).Error
	if err != nil {
		fmt.Println("[StudentRepoImpl.ViewById] Error when execute query")
		return nil, fmt.Errorf("name does not exist")
	}
	return student, nil
}

func (e *StudentRepoImpl) Insert(student *model.Student)(*model.Student, error) {
	err := e.DB.Save(student).Error
	if err != nil {
		fmt.Println("[StudentRepoImpl.Insert] Error When insert Query")
		return nil, fmt.Errorf("Insert data is failed")
	}
	return student, nil
}

func (e *StudentRepoImpl) Update(id int, student *model.Student)(*model.Student, error) {
	var newStudent = new(model.Student)
	err := e.DB.Table("student").Where("ID = ?", id).First(&newStudent).Update(&student).Error
	if err != nil {
		fmt.Println("[StudentRepoImpl.Update] Error When update Query")
		return nil, fmt.Errorf("Update data is failed")
	}
	return newStudent, nil
}

func (e *StudentRepoImpl) Delete(id int) error {
	var student = model.Student{}
	err := e.DB.Table("student").Where("ID = ?", id).Delete(&student).Error
	if err != nil {
		fmt.Println("[StudentRepoImpl.Delete] Error When delete Query")
		return fmt.Errorf("Delete data is failed")
	}
	return nil
}