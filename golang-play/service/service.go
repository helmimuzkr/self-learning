package service

import "github.com/helmimuzkr/golang-play/repository"

type Service interface {
	Find() []string
	Delete(id int) string
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Find() []string {
	find := &repository.RepsitoryImpl{}
	return find.Find()
}
func (s *service) Delete(id int) string {
	find := &repository.RepsitoryImpl{}
	return find.Delete(id)
}
