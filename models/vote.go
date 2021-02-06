package models

import (
	"blablastar/consts"
	"blablastar/drivers"
	"blablastar/models/data"
	"context"
)

type Vote struct {
	ID	string	`json:"id"`
	data.Vote
}


//Add add a new Vote into the collection and the search index
func (s *Vote) Add() error {
	ref, _, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.VOTE).Add(context.Background(), s.Vote)
	if err != nil {
		return err
	}
	s.ID = ref.ID

	return nil
}

//Update update a Vote in the collection and index
func (s *Vote) Update() error {
	_, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.VOTE).Doc(s.ID).Set(context.Background(), s.Vote)
	return err
}

//Get get a Vote with it's key
func (s *Vote) Get() error {
	doc, err := drivers.GetDriverFactory().FirebaseClient.Doc(s.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&s.Vote)
	if err != nil {
		return err
	}
	return nil
}

//Remove remove a Vote with it's key
func (s *Vote) Remove() error {
	_, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.VOTE).Doc(s.ID).Delete(context.Background())
	return err
}