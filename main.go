package main

import (
	"SecretSanta/api"
	"SecretSanta/controller"
	"SecretSanta/repository"
	"SecretSanta/service"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	//personRepository repository.PersonRepository = repository.NewFlatFileRepository()
	personRepository repository.PersonRepository = repository.NewCouchbaseRepository()

	personServices service.PersonServices = service.NewPersonService(personRepository)

	personController controller.PersonController = controller.NewPersonController(personServices)

	httpRouter api.Router = api.NewMuxRouter()
)

func main() {

	//Loading the Configuration Files
	err := godotenv.Load("configurations/config.env")

	if err != nil {
		logrus.Errorln("Error while loading the Configuration Variables: ", err)
	}

	//The Application Port Number
	//const port string = ":8382"
	port := os.Getenv("APPLICATION_PORT")
	logrus.Infoln(port)

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		logrus.Info("Server Is Running !!")
	})

	httpRouter.GET("/personWish", personController.GetPersonWish)
	httpRouter.POST("/addPersonWish", personController.AddPersonWish)
	httpRouter.PUT("/allocateSanta", personController.AllocateSanta)

	httpRouter.SERVE(port)
}
