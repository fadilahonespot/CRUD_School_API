package subject

import "school/model"

type SubjectRepo interface {
	ViewAll() (*[]model.Subject, error)
	ViewByid(id int) (*model.Subject, error)
	Insert(subject *model.Subject) (*model.Subject, error)
	Update(id int, subject *model.Subject) (*model.Subject, error)
	Delete(id int) error 
}