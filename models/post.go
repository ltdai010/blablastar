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

func (p *Post) getCollection() *firestore.CollectionRef {
	return drivers.GetDriverFactory().FirebaseClient.Collection(consts.POST)
}

//Add add a new Post into the collection and the search index
func (p *Post) Add() error {
	ref, _, err := p.getCollection().Add(context.Background(), p.Post)
	if err != nil {
		return err
	}
	p.ID = ref.ID
	go drivers.GetDriverFactory().SearchPostIndex.SaveObject(data.SearchPost{
		ObjectID: p.ID,
		Post:     p.Post,
	})

	return nil
}

//Update update a Post in the collection and index
func (s *Post) Update() error {
	return drivers.GetDriverFactory().FirebaseClient.RunTransaction(context.Background(),
		func(ctx context.Context, transaction *firestore.Transaction) error {
			err := transaction.Set(s.getCollection().Doc(s.ID), s.Post)
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
	doc, err := s.getCollection().Doc(s.ID).Get(context.Background())
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
			err := transaction.Delete(s.getCollection().Doc(s.ID))
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


func (s *Post) GetMulti(listID []string) ([]*Post, error) {
	res := []*Post{}
	for _, id := range listID {
		post := &Post{ID: id}
		err := post.Get()
		if err != nil {
			continue
		}
		res = append(res, post)
	}

	return res, nil
}

//GetPage get paginate
func (s *Post) GetPageID(pageNumber, pageSize int) ([]string, int, error) {
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