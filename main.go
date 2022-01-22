package main

import (
	"SecretSanta/api"
	"SecretSanta/controller"
	"SecretSanta/repository"
	"SecretSanta/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	personRepository repository.PersonRepository = repository.NewFlatFileRepository()

	personServices service.PersonServices = service.NewPersonService(personRepository)

	personController controller.PersonController = controller.NewPersonController(personServices)

	httpRouter api.Router = api.NewMuxRouter()
)

func main() {

	const port string = ":8382"

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		logrus.Info("Server Is Running !!")
	})

	httpRouter.GET("/personWish", personController.GetPersonWish)
	httpRouter.POST("/addPersonWish", personController.AddPersonWish)
	httpRouter.PUT("/allocateSanta", personController.AllocateSanta)

	httpRouter.SERVE(port)
}
