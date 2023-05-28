package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Sagar-Chowdhury/MongoDBGo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://syncwithsagar:<sagar>@cluster0.nims6uq.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// basically used for initialization before main method.
// here it is serving the purpouse of connecting to the DB.
func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption) //todo basically placeholder context.
	if err != nil {
		panic(err)
	}
	fmt.Println("MongoDB connection success")
	collection = (client.Database(dbName).Collection(colName))
	fmt.Println("Collection instance is ready")

}

//mongodb helper methods

//these methods are not to be exported hence named in small case.

// insert 1 record.
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}

// delete one movie
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Items deleted =", deleteCount.DeletedCount)
}

// delete all movies
func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}) //no filter paramets as we
	//need to delete all.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted = ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount

}

// getting all movies
func getAllMovies() []bson.M {

	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var movies []bson.M

	for cursor.Next(ctx) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies

}
