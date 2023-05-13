package mongodbsetup

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collectionListSingleton map[string]*mongo.Collection

var (
	CollectionList           = make(map[string]*mongo.Collection) // Global MongoDB pointer for each Collection in viper
	singleCollectionInstance collectionListSingleton
	once                     sync.Once
)

// Singleton Design Pattern
// Used here in order to open  onnections to the database and maintain a single instance so that we do not overload the database.
type collectionListSingletonC map[string]*mongo.Collection

func getCollectionInstance() collectionListSingleton {
	once.Do(func() { //Atomic call will only happen once
		singleCollectionInstance = make(collectionListSingleton)
	})
	return singleCollectionInstance
}

func init() {

	log.Println("Setting up Viper config")

	viper.SetConfigName("config.json")
	viper.AddConfigPath("D:\\github\\NetworkManagerMain")
	// viper.AddConfigPath("C:\\Users\\mecon\\Desktop\\NetworkManagerMain\\")
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	// Reading in the config path
	log.Println("Reading in Config File")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Fatal Error Reading in Config" + err.Error())
		panic("Error")
	}

	//init mongoDB connection
	log.Println("Initializing Mongo DB Connection")
	MongoClient, err := createMongoDbConnect()
	if err != nil {
		log.Println("Error in creating a mongoDB Client: " + err.Error())
		panic(err)
	}

	log.Println("Initializing Mongo DB Collections")
	if !viper.IsSet("db.dbCollections") {
		panic("Db Collection not found")
	}

	collections := viper.GetStringSlice("db.dbCollections")
	fmt.Println(collections)
	for _, collection := range collections {
		getCollectionInstance()[collection] = MongoClient.Database(Database).Collection(collection)
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

	// Todo move to main init if needed

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

	err := getCollectionInstance()[collection].FindOne(context.Background(), query).Decode(results)
	if err != nil {
		log.Println("Failed to Query the Database with Error: " + err.Error())
		return err
	}

	return nil
}

// Find All creates a connection to the mongodb and then queries based upon provided query and updates the
// original result memory space
func FindAll(query interface{}, results interface{}, collection string) error {

	cur, err := getCollectionInstance()[collection].Find(context.Background(), query)

	err = cur.All(context.Background(), results)
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())

	return nil
}
