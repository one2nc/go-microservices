/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"user-microservice/m/databases"
	"user-microservice/m/models"
	"user-microservice/m/utils"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {

	collection := databases.Database.MgCollectionName
	var user models.User
	var users []models.User

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		defer cursor.Close(context.TODO())
		return users, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {
	collection := databases.Database.MgCollectionName
	var user models.User
	objectId := bson.ObjectIdHex(id)
	err := collection.
		FindOne(context.TODO(), bson.M{"_id": objectId}).
		Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}
	return nil
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	collection := databases.Database.MgCollectionName
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"name": name, "password": password}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil

}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return err
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": user.ID})
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}
