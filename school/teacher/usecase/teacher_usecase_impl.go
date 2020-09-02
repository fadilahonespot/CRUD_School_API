package usecase 

import (
	"school/teacher"
	"school/model"
)

type TeacherUsecaseImpl struct {
	teacherRepo teacher.TeacherRepo
}

func CreateTeacherUsecase(teacherRepo teacher.TeacherRepo) teacher.TeacherUsecase {
	return &TeacherUsecaseImpl{teacherRepo}
}

func (e *TeacherUsecaseImpl) ViewAll() (*[]model.Teacher, error) {
	return e.teacherRepo.ViewAll()
}

func (e *TeacherUsecaseImpl) ViewById(id int) (*model.Teacher, error) {
	return e.teacherRepo.ViewById(id)
}

func (e *TeacherUsecaseImpl) ViewByName(name string) (*model.Teacher, error) {
	return e.teacherRepo.ViewByName(name)
}

func (e *TeacherUsecaseImpl) Insert(teacher *model.Teacher)(*model.Teacher, error) {
	return e.teacherRepo.Insert(teacher)
}

func (e *TeacherUsecaseImpl) Update(id int, teacher *model.Teacher)(*model.Teacher, error) {
	return e.teacherRepo.Update(id, teacher)
}

func (e *TeacherUsecaseImpl) Delete(id int) error {
	return e.teacherRepo.Delete(id)
}