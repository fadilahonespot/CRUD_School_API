package student

import (
	"school/model"
)

type StudentRepo interface {
	ViewAll() (*[]model.Student, error)
	Insert(student *model.Student)(*model.Student, error)
	ViewById(id int) (*model.Student, error)
	ViewByName(name string) (*model.Student, error)
	Update(id int, student *model.Student)(*model.Student, error)
	Delete(id int) error
}
