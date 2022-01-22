package service

import (
	"SecretSanta/entity"
	"SecretSanta/repository"
)

type PersonServices interface {
	CreatePersonWish(personWish *entity.PersonWish) (*entity.PersonWish, error)
	GetAllWishes() ([]entity.PersonWish, error)
	AllocateSanta() error
}

type services struct {
}

var repo repository.PersonRepository

func NewPersonService(repos repository.PersonRepository) PersonServices {
	repo = repos
	return &services{}
}

func (*services) CreatePersonWish(personWish *entity.PersonWish) (*entity.PersonWish, error) {
	return repo.SavePerson(personWish)
}

func (*services) GetAllWishes() ([]entity.PersonWish, error) {
	return repo.GetAllWishes()
}

func (*services) AllocateSanta() error {
	return repo.AllocateSanta()
}
