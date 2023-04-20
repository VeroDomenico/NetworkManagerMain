package mongodbsetup

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoDbConnect() (*mongo.Client, error) {

	//using viper to pull config info
	viper.SetConfigName("configFile.json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	//Used for Login
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	//Credential pulled from viper string
	credential := options.Credential{
		AuthSource: viper.GetString("db.AuthSource"),
		Username:   viper.GetString("db.user"),
		Password:   viper.GetString("db.password"),
	}
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
