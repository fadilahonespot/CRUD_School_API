package usecase

import (
	"school/model"
	"school/subject"
)

type SubjectUsecaseImpl struct {
	subjectRepo subject.SubjectRepo
}

func CreateSubjectUsecase(subjectRepo subject.SubjectRepo) subject.SubjectUsecase {
	return &SubjectUsecaseImpl{subjectRepo}
}

func (e *SubjectUsecaseImpl) ViewAll()(*[]model.Subject, error) {
	return e.subjectRepo.ViewAll()
}

func (e *SubjectUsecaseImpl) ViewByid(id int) (*model.Subject, error){
	return e.subjectRepo.ViewByid(id)
}

func (e *SubjectUsecaseImpl) Insert(subject *model.Subject) (*model.Subject, error) {
	return e.subjectRepo.Insert(subject)
}

func (e *SubjectUsecaseImpl) Update(id int, subject *model.Subject) (*model.Subject, error) {
	return e.subjectRepo.Update(id, subject)
}

func (e *SubjectUsecaseImpl) Delete(id int) error {
	return e.subjectRepo.Delete(id)
}