package models

import (
	"blablastar/consts"
	"blablastar/drivers"
	"blablastar/models/data"
	"context"
)

type User struct {
	ID	string	`json:"id"`
	data.User
}

//Add add a new User into the collection and the search index
func (s *User) Add() error {
	ref, _, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.USER).Add(context.Background(), s.User)
	if err != nil {
		return err
	}
	s.ID = ref.ID

	return nil
}

//Update update a User in the collection and index
func (s *User) Update() error {
	_, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.USER).Doc(s.ID).Set(context.Background(), s.User)
	return err
}

//Get get a User with it's key
func (s *User) Get() error {
	doc, err := drivers.GetDriverFactory().FirebaseClient.Doc(s.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&s.User)
	if err != nil {
		return err
	}
	return nil
}

//Remove remove a User with it's key
func (s *User) Remove() error {
	_, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.USER).Doc(s.ID).Delete(context.Background())
	return err
}