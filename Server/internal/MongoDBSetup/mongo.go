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
	// viper.AddConfigPath("C:\\Users\\mecon\\Desktop\\NetworkManagerMain\\")
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
	err = client.Connect(context.TODO())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil
}

// Find one creates a connection with the mongo db and returns a queried result but only one
// and an error if there is a failure
func FindOne(query interface{}, results interface{}, database string, collection string) error {
	client, err := createMongoDbConnect()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	err = client.Database(database).Collection(collection).FindOne(context.Background(), query).Decode(results)
	if err != nil {
		log.Println("Failed to Query the Database with Error: " + err.Error())
		return err
	}

	return nil
}

// Find All creates a connection to the mongodb and then queries based upon provided query and updates the
// original result memoryspace
func FindAll(query interface{}, results interface{}, db string, col string) error {
	client, err := createMongoDbConnect()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	cur, err := client.Database(db).Collection(col).Find(context.Background(), query)
	defer cur.Close(context.Background())

	err = cur.All(context.Background(), results)
	if err != nil {
		return err
	}

	return nil
}
