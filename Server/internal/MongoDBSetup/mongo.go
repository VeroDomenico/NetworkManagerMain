package mongodbsetup

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoDbConnect() (*mongo.Client, error) {

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
	fmt.Printf("Hello %s", credential.Username)
	uri := viper.GetString("db.uri")
	// Connect to MongoDB server
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
