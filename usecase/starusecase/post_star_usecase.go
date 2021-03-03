package starusecase

import (
	"blablastar/models"
	"blablastar/models/data"
	"blablastar/utils/logger"
)

//AddOne add star
func (su *StarUseCase) AddOne(request data.Star) (*models.Star, error) {
	star := &models.Star{
		ID:   "",
		Star: request,
	}
	err := star.Add()
	if err != nil {
		return nil, err
	}
	logger.Info("[Create star] Create star api, star: %v", star)
	return star, err
}
