package models

import (
	"blablastar/consts"
	"blablastar/drivers"
	"blablastar/models/data"
	"cloud.google.com/go/firestore"
	"context"
)

type Star struct {
	ID	string	`json:"id"`
	data.Star
}

//Add add a new Star into the collection and the search index
func (s *Star) Add() error {
	ref, _, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.STAR).Add(context.Background(), s.Star)
	if err != nil {
		return err
	}
	s.ID = ref.ID
	go drivers.GetDriverFactory().SearchStarIndex.SaveObject(data.SearchStar{
		ObjectID: s.ID,
		Star:     s.Star,
	})

	return nil
}

//Update update a Star in the collection and index
func (s *Star) Update() error {
	return drivers.GetDriverFactory().FirebaseClient.RunTransaction(context.Background(),
		func(ctx context.Context, transaction *firestore.Transaction) error {
			err := transaction.Set(drivers.GetDriverFactory().FirebaseClient.Collection(consts.STAR).Doc(s.ID), s.Star)
			if err != nil {
				return err
			}
			_, err = drivers.GetDriverFactory().SearchStarIndex.SaveObject(data.SearchStar{
				ObjectID: s.ID,
				Star:     s.Star,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}

//Get get a Star with it's key
func (s *Star) Get() error {
	doc, err := drivers.GetDriverFactory().FirebaseClient.Doc(s.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&s.Star)
	if err != nil {
		return err
	}
	return nil
}

//Remove remove a Star with it's key
func (s *Star) Remove() error {
	return drivers.GetDriverFactory().FirebaseClient.RunTransaction(context.Background(),
		func(ctx context.Context, transaction *firestore.Transaction) error {
			err := transaction.Delete(drivers.GetDriverFactory().FirebaseClient.Collection(consts.STAR).Doc(s.ID))
			if err != nil {
				return err
			}
			_, err = drivers.GetDriverFactory().SearchStarIndex.SaveObject(data.SearchStar{
				ObjectID: s.ID,
				Star:     s.Star,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}