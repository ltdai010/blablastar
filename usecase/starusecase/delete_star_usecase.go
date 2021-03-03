package starusecase

import (
	"blablastar/models"
	"blablastar/utils/logger"
)

//DeleteOne delete a post
func (su *StarUseCase) DeleteOne(id string) error {
	star := &models.Star{ID: id}
	err := star.Get()
	if err != nil {
		return err
	}
	logger.Info("[Remove post] Remove star api, star id = %v", id)
	return star.Remove()
}
