package service

import "github.com/helmimuzkr/golang-play/repository"

type Service2 interface {
	Find() []string
	Delete(id int) string
}

type service2 struct {
	repo repository.Repository
}

func NewService2(repo repository.Repository) Service2 {
	return &service2{repo: repo}
}

func (s *service2) Find() []string {
	return s.repo.Find()
}
func (s *service2) Delete(id int) string {
	return s.repo.Delete(id)
}
