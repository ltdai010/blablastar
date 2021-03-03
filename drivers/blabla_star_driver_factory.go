package drivers

import (
	"blablastar/utils/logger"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/beego/beego/v2/server/web"
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
		logger.Info("[Finish init] Finish initial service")
	})
}

func initCloudStore() {
	ctx = context.Background()
	s, err := web.AppConfig.String("cloudstore::key_path")
	if err != nil {
		log.Panic(err)
	}
	sa := option.WithCredentialsFile(s)
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
	appID, err := web.AppConfig.String("algolia::app_id")
	if err != nil {
		log.Panic(err)
	}
	key, err := web.AppConfig.String("algolia::key")
	if err != nil {
		log.Panic(err)
	}
	clientSearch = search.NewClient(appID, key)
	postIndexName, err := web.AppConfig.String("algolia::post_index_name")
	if err != nil {
		log.Panic(err)
	}
	starIndexName, err := web.AppConfig.String("algolia::star_index_name")
	if err != nil {
		log.Panic(err)
	}
	searchPostIndex = clientSearch.InitIndex(postIndexName)
	searchStarIndex = clientSearch.InitIndex(starIndexName)
}
