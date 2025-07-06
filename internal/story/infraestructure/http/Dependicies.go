package http

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/story/application"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/ports"
	adapters "github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/adapter"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/http/controllers"
)

var StoryRepository ports.StoryRepository

func init() {
	var err error
	StoryRepository, err = adapters.NewStoryRepositoryMongoDB()

	if err != nil{
		panic("error al inicializar el adapter")
	}
}


func SetUpCreate()*controllers.CreateController{
	createUseCase := application.NewCreateUseCase(StoryRepository)
	return controllers.NewCreateUseCase(createUseCase)
}

func SetUpGet()*controllers.GetController{
	getUseCae := application.NewGetUseCase(StoryRepository)
	return controllers.NewGetController(getUseCae)
}

func SetUpGetById()*controllers.GetByIdController{
	getByIdUseCase := application.NewGetByIdUseCase(StoryRepository)
	return controllers.NewGetByIdController(getByIdUseCase)
}