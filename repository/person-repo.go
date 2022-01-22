package repository

import "SecretSanta/entity"

type PersonRepository interface {
	SavePerson(personWish *entity.PersonWish) (*entity.PersonWish, error)
	GetAllWishes() ([]entity.PersonWish, error)
	AllocateSanta() error
}
