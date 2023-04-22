package mongodbsetup

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Establishes connection with mongoDB server
func createMongoDbConnect() (*mongo.Client, error) {

	//using viper to pull config info
	viper.SetConfigName("config.json")
	viper.AddConfigPath("D:\\github\\NetworkManagerMain")
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	//Used for Login
	log.Println("Reading config file for connecting to mongo db")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	//Credential pulled from viper string to connect
	log.Println("Read File")
	credential := options.Credential{
		AuthSource: viper.GetString("db.AuthSource"),
		Username:   viper.GetString("db.user"),
		Password:   viper.GetString("db.password"),
	}

	uri := viper.GetString("db.uri")

	log.Println("Establishing Connection to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil
}

// this will take in a filter and query the mongodb for one match and return a result of the struct passed in and an error if there is a failure
func FindOne(query interface{}, results interface{}, database string, collection string) (interface{}, error) {
	client, err := createMongoDbConnect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	err = client.Database(database).Collection(collection).FindOne(context.Background(), query).Decode(results)
	if err != nil {
		log.Println("Failed to Query the Database with Error: " + err.Error())
		return nil, err
	}

	return results, nil
}

// func FindAll(query interface{}, results interface{}, database string, collection string) (results interface{}, error) {

// 	client, err := createMongoDbConnect()
// 	if err != nil {
// 		return nil, err
// 	}

// 	//TODO
// 	// if client.Database(database).Collection(collection).Find(context.Background(), query).Decode(results) != nil {
// 	// 	log.Println("Failed to Query the Database with Error: " + err.Error())
// 	// 	return nil, err
// 	// }

// 	if client.Database(database).Collection(collection).Find(context.Background(),query).Decode(res)

// 	log.Println("Cleaning up allocated memory for monogdb connection")
// 	if client.Disconnect(context.Background()) != nil {
// 		log.Println("Failed to close mongodb connection")
// 		return nil, err
// 	}
// 	return nil, nil

// }
