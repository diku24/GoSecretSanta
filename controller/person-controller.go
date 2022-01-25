package controller

import (
	"SecretSanta/entity"
	errors "SecretSanta/errors"
	"SecretSanta/service"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type PersonController interface {
	AddPersonWish(response http.ResponseWriter, request *http.Request)
	GetPersonWish(response http.ResponseWriter, request *http.Request)
	AllocateSanta(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

var personServices service.PersonServices

func NewPersonController(service service.PersonServices) PersonController {
	personServices = service
	return &controller{}
}

func (*controller) AddPersonWish(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var person entity.PersonWish

	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error Unmarshalling the Request"})
		return
	}

	result, err := personServices.CreatePersonWish(&person)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error Saving the request"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*controller) GetPersonWish(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("context-type", "application/json")

	nameParam := request.URL.Query().Get("name")
	santaParam := request.URL.Query().Get("santa")

	person, err := personServices.GetAllWishes()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error Getting the wishes"})
		return
	}

	var newArr []entity.PersonWish

	if len(request.URL.Query()) > 0 {
		for _, element := range person {

			if nameParam != "" && santaParam != "" {
				if element.Name == nameParam || element.Santa == santaParam {
					newArr = append(newArr, element)
				}
			} else if nameParam != "" {
				if element.Name == nameParam {
					newArr = append(newArr, element)
				}
			} else if santaParam != "" {
				if element.Name == santaParam {
					newArr = append(newArr, element)
				}
			}
		}
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(newArr)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(person)
	}
}

func (*controller) AllocateSanta(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("context-type", "application/json")
	err := personServices.AllocateSanta()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error Allocatting the Santa to existing wishes"})
		return
	}
	response.WriteHeader(http.StatusOK)
}
