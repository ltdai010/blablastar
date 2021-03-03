package postusecase

import (
	"blablastar/models"
	"blablastar/models/data"
	"blablastar/utils/logger"
)

//AddOne add post
func (su *PostUseCase) AddOne(request data.Post) (*models.Post, error) {
	post := &models.Post{
		ID:   "",
		Post: request,
	}
	err := post.Add()
	if err != nil {
		return nil, err
	}
	logger.Info("[Create post] Create post api, post: %v", post)
	return post, err
}

