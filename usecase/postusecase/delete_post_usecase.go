package postusecase

import (
	"blablastar/models"
	"blablastar/utils/logger"
)

//DeleteOne delete post with it's id
func (pu *PostUseCase) DeleteOne(id string) error {
	post := &models.Post{ID: id}
	err := post.Get()
	if err != nil {
		return err
	}
	logger.Info("[Remove post] Remove post api, post id = %v", id)
	return post.Remove()
}

