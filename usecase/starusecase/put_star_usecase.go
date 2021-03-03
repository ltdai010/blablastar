package starusecase

import (
	"blablastar/models"
	"blablastar/models/data"
)

//UpdateOne check if star exist then update star
func (su *StarUseCase) UpdateOne(id string, request data.Star) error {
	star := &models.Star{ID: id}
	err := star.Get()
	if err != nil {
		return err
	}
	star.Star = request
	err = star.Update()
	return err
}
