package drivers

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/astaxie/beego"
	"google.golang.org/api/option"
	"context"
	"log"
	"sync"
)

var (
	initOnce        sync.Once
	firebaseClient  *firestore.Client
	ctx             context.Context
	clientSearch    *search.Client
	searchPostIndex *search.Index
	searchStarIndex *search.Index
	driverFactory   *DriverFactory
)

type DriverFactory struct {
	Mutex				*sync.Mutex
	FirebaseClient		*firestore.Client
	SearchPostIndex		search.IndexInterface
	SearchStarIndex     search.IndexInterface
}

func GetDriverFactory() *DriverFactory {
	onceInitEverything()

	return driverFactory
}

func onceInitEverything() {
	initOnce.Do(func() {
		initCloudStore()
		initSearch()
		driverFactory = &DriverFactory{
			Mutex:          &sync.Mutex{},
			FirebaseClient: firebaseClient,
			SearchPostIndex:    searchPostIndex,
			SearchStarIndex:    searchStarIndex,
		}
	})
}

func initCloudStore() {
	ctx = context.Background()
	sa := option.WithCredentialsFile(beego.AppConfig.String("cloudstore::key_path"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Panic(err)
	}
	firebaseClient, err = app.Firestore(ctx)
	if err != nil {
		log.Panic(err)
	}
}

func initSearch() {
	clientSearch = search.NewClient(beego.AppConfig.String("algolia::app_id"),
		beego.AppConfig.String("algolia::key"))
	searchPostIndex = clientSearch.InitIndex("algolia::post_index_name")
	searchStarIndex = clientSearch.InitIndex("algolia::star_index_name")
}
