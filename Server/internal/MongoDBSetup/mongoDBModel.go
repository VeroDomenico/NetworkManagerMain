package mongodbsetup

import "go.mongodb.org/mongo-driver/mongo"

/*
Data Structures are going to be located here
*/
const (
	Database = "NetworkManager"
)

// dbModel to reference
type dbModel struct {
	Database        string
	CollectionName  string
	MongoCollection *mongo.Collection
}

// Database Model Struct for Data
