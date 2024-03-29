package controller_test

import (
	"SecretSanta/controller"
	"SecretSanta/entity"
	"SecretSanta/mocks"
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddPersonWishFile(t *testing.T) {
	//Control Initilizing the Mock Controller
	control := gomock.NewController(t)

	//Removing the instance of controller
	defer control.Finish()

	//mockService Creating mockrepository varible
	mockService := mocks.NewMockPersonServices(control)

	//Mocking the services Constructor with mockRepository for testing
	testService := controller.NewPersonController(mockService)

	//response to the test func
	response := httptest.NewRecorder()

	//PersonWish byte array to request body
	personWishesByte := []byte(`{Name: "Dayna", Wishes: []string{"state","Rest","deposit mobile"}`)

	//request for new request to test func
	request := httptest.NewRequest("POST", "/addPersonWish", bytes.NewBuffer(personWishesByte))

	//pass reuqest and responce to cotroller fun
	testService.AddPersonWish(response, request)

	//Check Status code of the reponse
	statusCode := response.Code

	//Decode the http responce
	var person entity.PersonWish
	json.NewDecoder(io.Reader(response.Body)).Decode(&person)

	//assertions for failing
	//assert.Equal(t, 200, statusCode)
	assert.Equal(t, 500, statusCode)
	assert.Equal(t, "", person.Name)
	assert.Equal(t, []string(nil), person.Wishes)

}

func TestGetPersonWish(t *testing.T) {

	//Control Initilizing the Mock Controller
	control := gomock.NewController(t)

	//Removing the instance of controller
	defer control.Finish()

	//mockService Creating mockrepository varible
	mockService := mocks.NewMockPersonServices(control)

	//Data to be mocked
	personWishes := entity.PersonWish{Name: "John", Wishes: []string{"BatMobile", "Destiny"}}

	//Creating the expectation and call the required function
	mockService.EXPECT().GetAllWishes().Return([]entity.PersonWish{personWishes}, nil)

	//Mocking the services Constructor with mockRepository for testing
	testService := controller.NewPersonController(mockService)

	//response to the test func
	response := httptest.NewRecorder()

	//request for new request to test func
	request := httptest.NewRequest("", "/personWish", nil)

	//pass reuqest and responce to cotroller fun
	testService.GetPersonWish(response, request)

	//Check Status code of the reponse
	statusCode := response.Code

	//Decode the http responce
	var persons []entity.PersonWish
	json.NewDecoder(io.Reader(response.Body)).Decode(&persons)

	//assertions
	assert.Equal(t, personWishes.Name, persons[0].Name)
	assert.Equal(t, personWishes.Wishes, persons[0].Wishes)
	assert.Equal(t, 200, statusCode)
}

func TestAllocateSanta(t *testing.T) {

	//Control Initilizing the Mock Controller
	control := gomock.NewController(t)

	//Removing the instance of controller
	defer control.Finish()

	//mockService Creating mockrepository varible
	mockService := mocks.NewMockPersonServices(control)

	//Creating the expectation and call the required function
	mockService.EXPECT().AllocateSanta().Return(nil)

	//Mocking the services Constructor with mockRepository for testing
	testService := controller.NewPersonController(mockService)

	//response to the test func
	response := httptest.NewRecorder()

	//request for new request to test func
	request := httptest.NewRequest("PUT", "/allocateSanta", nil)

	//pass reuqest and responce to cotroller fun
	testService.AllocateSanta(response, request)

	//Check Status code of the reponse
	statusCode := response.Code

	//Decode the http responce
	var persons []entity.PersonWish
	json.NewDecoder(io.Reader(response.Body)).Decode(&persons)

	//assertions

	assert.Equal(t, 200, statusCode)
}
