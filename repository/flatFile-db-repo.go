package repository

import (
	"SecretSanta/entity"
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/sirupsen/logrus"
)

type repository struct {
}

func NewFlatFileRepository() PersonRepository {
	return &repository{}
}

func (*repository) SavePerson(personWish *entity.PersonWish) (*entity.PersonWish, error) {

	//Read from the File.
	readFile, err := ioutil.ReadFile("db.json")
	if err != nil {
		logrus.Errorf("error while reading the data", err)
	}

	//Define Slice of PersonWishes
	var entitypersons []*entity.PersonWish

	//Unmarshall it
	if err = json.Unmarshal([]byte(readFile), &entitypersons); err != nil {
		logrus.Errorf("erroer unmarshallin the data", err)
	}

	//Append this to entityperson
	entitypersons = append(entitypersons, personWish)

	//Now marshall it
	//Now result has your targeted JSON structures
	result, err := json.Marshal(entitypersons)
	if err != nil {
		logrus.Errorf("error marshalling")
	}

	//Write to the File
	_ = ioutil.WriteFile("db.json", result, 0777)

	return personWish, nil
}

func (*repository) GetAllWishes() ([]entity.PersonWish, error) {
	data := []entity.PersonWish{}

	//ReadFile from database File
	fileContent, err := ioutil.ReadFile("db.json")
	if err != nil {
		logrus.Error("error while reading data")
	}

	_ = json.Unmarshal([]byte(fileContent), &data)
	return data, nil
}

func (*repository) AllocateSanta() error {

	readFile, err := ioutil.ReadFile("db.json")
	if err != nil {
		logrus.Error("error while reading file")
	}

	var entitypersons []*entity.PersonWish

	if err := json.Unmarshal([]byte(readFile), &entitypersons); err != nil {
		logrus.Error("error in unmarshalling ")
	}

	length := len(entitypersons)
	id := rand.Intn(length)
	for i, element := range entitypersons {

		if element.Santa == "" {
			if i == id {
				id = id - 1
			}
			element.Santa = entitypersons[id].Name
		}

		//entitypersonss := append(entitypersons, entitypersons[i])
		entitypersonss := append(entitypersons, element)

		result, err := json.Marshal(entitypersonss)
		if err != nil {
			logrus.Error("error in marshling the data")
		}

		_ = ioutil.WriteFile("db.json", result, 0777)
	}
	return err

}
