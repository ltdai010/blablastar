package postusecase

import (
	"blablastar/models"
	"blablastar/models/data"
)

//UpdateOne check if star exist then update star
func (su *PostUseCase) UpdateOne(id string, request data.Post) error {
	post := &models.Post{ID: id}
	err := post.Get()
	if err != nil {
		return err
	}
	post.Post = request
	err = post.Update()
	return err
}