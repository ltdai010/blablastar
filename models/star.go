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

func (s *Star) getCollection() *firestore.CollectionRef {
	return drivers.GetDriverFactory().FirebaseClient.Collection(consts.STAR)
}

//Add add a new Star into the collection and the search index
func (s *Star) Add() error {
	ref, _, err := s.getCollection().Add(context.Background(), s.Star)
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
			err := transaction.Set(s.getCollection().Doc(s.ID), s.Star)
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
	doc, err := s.getCollection().Doc(s.ID).Get(context.Background())
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
			err := transaction.Delete(s.getCollection().Doc(s.ID))
			if err != nil {
				return err
			}
			_, err = drivers.GetDriverFactory().SearchStarIndex.DeleteObject(s.ID)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func (s *Star) GetMulti(listID []string) ([]*Star, error) {
	res := []*Star{}
	for _, id := range listID {
		star := &Star{ID: id}
		err := star.Get()
		if err != nil {
			continue
		}
		res = append(res, star)
	}

	return res, nil
}

//GetPage get paginate
func (s *Star) GetPageID(pageNumber, pageSize int) ([]string, int, error) {
	pageNumber--
	//var
	start := pageNumber * pageSize
	end := (pageNumber + 1) * pageSize
	total := 0
	listID := []string{}
	listRef, err := s.getCollection().DocumentRefs(context.Background()).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total = len(listRef)
	if total < start {
		return nil, 0, data.NotMore
	}
	if total < end {
		end = total
	}

	for _, ref := range listRef[start:end] {
		listID = append(listID, ref.ID)
	}

	return listID, total, nil
}