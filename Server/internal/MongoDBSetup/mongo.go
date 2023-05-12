package mongodbsetup

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	CollectionList = make(map[string]*mongo.Collection) // Global MongoDB pointer for each Collection in viper
	singleInstance *collectionListSingleton
	lock           = &sync.Mutex{}
)

// Singleton Design Pattern
type collectionListSingleton struct {
	// This construct is prvate, there will only be a singleton class can create instance of itself
	CollectionList map[string]*mongo.Collection
}

func getInstance() *collectionListSingleton {
	if singleInstance == nil {
		lock.Lock()
		// In case of error unlock
		defer lock.Unlock()
		if singleInstance == nil {
			log.Println("Creating singleton for collections")
			singleInstance = &collectionListSingleton{}
		} else {
			log.Println("Single Instance already Created")
		}
	} else {
		log.Println("Single Instance already Created")
	}
	return singleInstance
}

func init() {
	log.Println("Initializing Mongo DB Connection")
	MongoClient, err := createMongoDbConnect()
	if err != nil {
		log.Println("Error in creating a mongoDB Client: " + err.Error())
		panic(err)

	}
	MongoClient.Ping(context.Background(), readpref.Primary())

	log.Println("Initializing Mongo DB Collections")
	if !viper.IsSet("db.dbCollections") {
		panic("Db Collection not found")
	}

	collections := viper.GetStringSlice("db.dbCollections")
	fmt.Println(collections)
	for _, collection := range collections {
		CollectionList[collection] = MongoClient.Database(Database).Collection(collection)
	}
}

// Establishes connection with mongoDB server
func createMongoDbConnect() (*mongo.Client, error) {

	// TODO if credentials are needed
	//Credential pulled from viper string to connect
	// log.Println("Loading Credentials")
	// credential := options.Credential{
	// 	AuthSource: viper.GetString("db.AuthSource"),
	// 	Username:   viper.GetString("db.user"),
	// 	Password:   viper.GetString("db.password"),
	// }

	uri := viper.GetString("db.uri")

	log.Println("Establishing Connection to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri)) //.SetAuth(credential)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println("Connecting to client")
	err = client.Connect(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return client, nil
}

// Find one creates a connection with the mongo db and returns a queried result but only one
// and an error if there is a failure
func FindOne(query interface{}, results interface{}, collection string) error {

	err := CollectionList[collection].FindOne(context.Background(), query).Decode(results)
	if err != nil {
		log.Println("Failed to Query the Database with Error: " + err.Error())
		return err
	}

	return nil
}

// Find All creates a connection to the mongodb and then queries based upon provided query and updates the
// original result memory space
func FindAll(query interface{}, results interface{}, collection string) error {

	cur, err := CollectionList[collection].Find(context.Background(), query)

	err = cur.All(context.Background(), results)
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())

	return nil
}
