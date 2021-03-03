package starusecase

import (
	"blablastar/models"
	"blablastar/utils/logger"
)

//GetOne get star by id
func (su *StarUseCase) GetOne(id string) (*models.Star, error) {
	star := &models.Star{ID: id}
	err := star.Get()
	if err != nil {
		return nil, err
	}
	logger.Info("[Get star] Get star api, star id = %v", id)
	return star, err
}

//GetPage get list stars by page
func (su *StarUseCase) GetPage(pageNumber, pageSize int) ([]*models.Star, int, error) {
	//var
	star := &models.Star{}
	res := []*models.Star{}

	listID, total, err := star.GetPageID(pageNumber, pageSize)
	if err != nil {
		return res, 0, err
	}
	res, err = star.GetMulti(listID)
	if err != nil {
		return res, 0, err
	}
	return res, total, nil
}
