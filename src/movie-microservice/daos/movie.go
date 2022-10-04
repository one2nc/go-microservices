/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"movie-microservice/m/databases"
	"movie-microservice/m/models"
)

// Movie manages Movie CRUD
type Movie struct {
}

// GetAll gets the list of Movie
func (m *Movie) GetAll() ([]models.Movie, error) {
	collection := databases.Database.MgCollectionName
	var movie models.Movie
	var movies []models.Movie

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		defer cursor.Close(context.TODO())
		return movies, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&movie)
		if err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// GetByID finds a Movie by its id
func (m *Movie) GetByID(id string) (models.Movie, error) {
	collection := databases.Database.MgCollectionName
	var movie models.Movie
	objectId := bson.ObjectIdHex(id)
	err := collection.
		FindOne(context.TODO(), bson.M{"_id": objectId}).
		Decode(&movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

// Insert adds a new Movie into database'
func (m *Movie) Insert(movie models.Movie) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return err
	}
	return err
}

// Delete remove an existing Movie
func (m *Movie) Delete(movie models.Movie) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": movie.ID})
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing Movie
func (m *Movie) Update(movie models.Movie) error {
	collection := databases.Database.MgCollectionName
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": movie.ID}, bson.M{"$set": movie})
	return err
}
