package models

import (
	"blablastar/consts"
	"blablastar/drivers"
	"blablastar/models/data"
	"context"
)

type Comment struct {
	ID	string	`json:"id"`
	data.Comment
}

//Add add a new Comment into the collection and the search index
func (s *Comment) Add() error {
	ref, _, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.COMMENT).Add(context.Background(), s.Comment)
	if err != nil {
		return err
	}
	s.ID = ref.ID

	return nil
}

//Update update a Comment in the collection and index
func (s *Comment) Update() error {
	_, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.COMMENT).Doc(s.ID).Set(context.Background(), s.Comment)
	return err
}

//Get get a Comment with it's key
func (s *Comment) Get() error {
	doc, err := drivers.GetDriverFactory().FirebaseClient.Doc(s.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&s.Comment)
	if err != nil {
		return err
	}
	return nil
}

//Remove remove a Comment with it's key
func (s *Comment) Remove() error {
	_, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.COMMENT).Doc(s.ID).Delete(context.Background())
	return err
}
