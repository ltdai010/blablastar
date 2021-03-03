package starusecase

import (
	"blablastar/models"
	"blablastar/models/data"
	"sync"
)

var (
	onceDo sync.Once
	starUseCaseInterface StarUseCaseInterface
)

type StarUseCaseInterface interface {
	AddOne(request data.Star) (star *models.Star, err error)
	GetOne(id string) (star *models.Star, err error)
	UpdateOne(id string, request data.Star) (err error)
	DeleteOne(id string) (err error)
	GetPage(pageNumber, pageSize int) (star []*models.Star, total int, err error)
}

type StarUseCase struct {}

func GetStarUseCase() StarUseCaseInterface {
	onceDo.Do(onceInit)
	return starUseCaseInterface
}

func onceInit() {
	starUseCaseInterface = &StarUseCase{}
}

