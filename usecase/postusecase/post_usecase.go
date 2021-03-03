package postusecase

import (
	"blablastar/models"
	"blablastar/models/data"
	"sync"
)

var (
	onceDo sync.Once
	postUseCaseInterface PostUseCaseInterface
)

type PostUseCaseInterface interface {
	AddOne(request data.Post) (star *models.Post, err error)
	GetOne(id string) (post *models.Post, err error)
	UpdateOne(id string, request data.Post) (err error)
	DeleteOne(id string) (err error)
	GetPage(pageNumber, pageSize int) (star []*models.Post, total int, err error)
}


type PostUseCase struct {}

func GetPostUseCase() PostUseCaseInterface {
	onceDo.Do(onceInit)
	return postUseCaseInterface
}

func onceInit() {
	postUseCaseInterface = &PostUseCase{}
}