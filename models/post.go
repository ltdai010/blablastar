package models

import (
	"blablastar/consts"
	"blablastar/drivers"
	"blablastar/models/data"
	"cloud.google.com/go/firestore"
	"context"
)

type Post struct {
	ID string	`json:"id"`
	data.Post
}

//Add add a new Post into the collection and the search index
func (s *Post) Add() error {
	ref, _, err := drivers.GetDriverFactory().FirebaseClient.Collection(consts.POST).Add(context.Background(), s.Post)
	if err != nil {
		return err
	}
	s.ID = ref.ID
	go drivers.GetDriverFactory().SearchPostIndex.SaveObject(data.SearchPost{
		ObjectID: s.ID,
		Post:     s.Post,
	})

	return nil
}

//Update update a Post in the collection and index
func (s *Post) Update() error {
	return drivers.GetDriverFactory().FirebaseClient.RunTransaction(context.Background(),
		func(ctx context.Context, transaction *firestore.Transaction) error {
			err := transaction.Set(drivers.GetDriverFactory().FirebaseClient.Collection(consts.POST).Doc(s.ID), s.Post)
			if err != nil {
				return err
			}
			_, err = drivers.GetDriverFactory().SearchPostIndex.SaveObject(data.SearchPost{
				ObjectID: s.ID,
				Post:     s.Post,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}

//Get get a Post with it's key
func (s *Post) Get() error {
	doc, err := drivers.GetDriverFactory().FirebaseClient.Doc(s.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&s.Post)
	if err != nil {
		return err
	}
	return nil
}

//Remove remove a Post with it's key
func (s *Post) Remove() error {
	return drivers.GetDriverFactory().FirebaseClient.RunTransaction(context.Background(),
		func(ctx context.Context, transaction *firestore.Transaction) error {
			err := transaction.Delete(drivers.GetDriverFactory().FirebaseClient.Collection(consts.POST).Doc(s.ID))
			if err != nil {
				return err
			}
			_, err = drivers.GetDriverFactory().SearchPostIndex.SaveObject(data.SearchPost{
				ObjectID: s.ID,
				Post:     s.Post,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}
