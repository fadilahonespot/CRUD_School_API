package teacher

import (
	"school/model"
)

type TeacherRepo interface {
	ViewAll() (*[]model.Teacher, error)
	ViewById(id int) (*model.Teacher, error)
	ViewByName(name string) (*model.Teacher, error)
	Insert(teacher *model.Teacher)(*model.Teacher, error)
	Update(id int, teacher *model.Teacher)(*model.Teacher, error)
	Delete(id int) error
}