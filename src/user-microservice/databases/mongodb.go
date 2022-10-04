/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package databases

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
	"user-microservice/m/common"
	"user-microservice/m/models"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbSession      *mongo.Client
	Databasename     string
	MgCollectionName *mongo.Collection
}

// Init initializes mongo database
func (db *MongoDB) Init() {
	var err error
	db.Databasename = common.Config.MgDbName
	var uri string = "mongodb://" + common.Config.MgAddrs + "/?maxPoolSize=10&w=majority"
	db.MgDbSession, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := db.MgDbSession.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	db.MgCollectionName = db.MgDbSession.Database(db.Databasename).Collection(common.ColUsers)
	db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() {
	var user models.User
	err := db.MgCollectionName.FindOne(context.TODO(), bson.D{}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			var user models.User
			user = models.User{bson.NewObjectId(), "admin", "admin"}
			db.MgCollectionName.InsertOne(context.TODO(), &user)
		}
	}
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Disconnect(context.TODO())
	}
}
