package usecase

import (
	"school/student"
	"school/model"
)

type StudentUsecaseImpl struct {
	studentRepo student.StudentRepo
}

func CreateStudentUsecase(studentRepo student.StudentRepo) student.StudentUsecase {
	return &StudentUsecaseImpl{studentRepo}
}

func (e *StudentUsecaseImpl) ViewAll() (*[]model.Student, error) {
	return e.studentRepo.ViewAll()
}

func (e *StudentUsecaseImpl) Insert(student *model.Student)(*model.Student, error) {
	return e.studentRepo.Insert(student)
}

func (e *StudentUsecaseImpl) ViewById(id int) (*model.Student, error) {
	return e.studentRepo.ViewById(id)
}

func (e *StudentUsecaseImpl) ViewByName(name string) (*model.Student, error) {
	return e.studentRepo.ViewByName(name)
}

func (e *StudentUsecaseImpl) Update(id int, student *model.Student)(*model.Student, error) {
	return e.studentRepo.Update(id, student)
}

func (e *StudentUsecaseImpl) Delete(id int) error {
	return e.studentRepo.Delete(id)
}