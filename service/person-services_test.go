package service_test

import (
	"SecretSanta/entity"
	"SecretSanta/mocks"
	"SecretSanta/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePersonWish(t *testing.T) {

	//Control Initilizing the Mock Controller
	control := gomock.NewController(t)

	//Removing the instance of controller
	defer control.Finish()

	//mockPersonRepo Creating mockrepository varible
	mockPersonRepo := mocks.NewMockPersonRepository(control)

	//Data to be mocked
	personWishes := entity.PersonWish{Name: "John", Wishes: []string{"BatMobile", "Destiny"}}

	//Creating the expectation and call the required function
	mockPersonRepo.EXPECT().SavePerson(gomock.Any()).Return(&personWishes, nil)

	//Mocking the services Constructor with mockRepository for testing
	testService := service.NewPersonService(mockPersonRepo)

	//Mocking the the service Function eith mock data for testing.
	result, err := testService.CreatePersonWish(&personWishes)

	//Checking with expectations and actual values
	assert.Equal(t, "John", result.Name)
	assert.Equal(t, []string{"BatMobile", "Destiny"}, result.Wishes)
	assert.Nil(t, err)
}

func TestGetAllWishes(t *testing.T) {

	//Control Initilizing the Mock Controller
	control := gomock.NewController(t)

	//Removing the instance of controller
	defer control.Finish()

	//mockPersonRepo Creating mockrepository varible
	mockPersonRepo := mocks.NewMockPersonRepository(control)

	//Data to be mocked
	personWishes := entity.PersonWish{Name: "John", Wishes: []string{"BatMobile", "Destiny"}}

	//Creating the expectation and call the required function
	mockPersonRepo.EXPECT().GetAllWishes().Return([]entity.PersonWish{personWishes}, nil)

	//Mocking the services Constructor with mockRepository for testing
	testService := service.NewPersonService(mockPersonRepo)

	//Mocking the the service Function eith mock data for testing.
	result, err := testService.GetAllWishes()

	//Checking with expectations and actual values
	assert.Equal(t, "John", result[0].Name)
	assert.Equal(t, []string{"BatMobile", "Destiny"}, result[0].Wishes)
	assert.Nil(t, err)

}

func TestAllocateSanta(t *testing.T) {
	//Control Initilizing the Mock Controller
	control := gomock.NewController(t)

	//Removing the instance of controller
	defer control.Finish()

	//mockPersonRepo Creating mockrepository varible
	mockPersonRepo := mocks.NewMockPersonRepository(control)

	//Data to be mocked
	//personWishes := entity.PersonWish{Name: "John", Wishes: []string{"BatMobile", "Destiny"}}

	//Creating the expectation and call the required function
	mockPersonRepo.EXPECT().AllocateSanta().Return(nil)

	//Mocking the services Constructor with mockRepository for testing
	testService := service.NewPersonService(mockPersonRepo)

	//Mocking the the service Function eith mock data for testing.
	//result, err := testService.GetAllWishes()
	err := testService.AllocateSanta()

	//Checking with expectations and actual values
	assert.Nil(t, err)
}
