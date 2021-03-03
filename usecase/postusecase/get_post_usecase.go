package postusecase

import (
	"blablastar/models"
	"blablastar/utils/logger"
)

//GetOne get post by id
func (su *PostUseCase) GetOne(id string) (*models.Post, error) {
	star := &models.Post{ID: id}
	err := star.Get()
	if err != nil {
		return nil, err
	}
	logger.Info("[Get post] Get post api, post id = %v", id)
	return star, err
}

//GetPage get list stars by page
func (su *PostUseCase) GetPage(pageNumber, pageSize int) ([]*models.Post, int, error) {
	//var
	post := &models.Post{}
	res := []*models.Post{}

	listID, total, err := post.GetPageID(pageNumber, pageSize)
	if err != nil {
		return res, 0, err
	}
	res, err = post.GetMulti(listID)
	if err != nil {
		return res, 0, err
	}
	logger.Info("[Get page posts] Get page post api, page_number = %v, page_size = %v", pageNumber, pageSize)
	return res, total, nil
}
