package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Swayamshu/mangadex/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://swayamshu:swayamshu@cluster0.cyokxow.mongodb.net/?retryWrites=true&w=majority"
const dbName = "mangako"
const collectionName = "id-mapper"

var collection *mongo.Collection

// connect to MongoDB
func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection established!")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created.")
}

func InsertMapperData(mapperData model.MapperData) {
	inserted, err := collection.InsertOne(context.Background(), mapperData)

	if err != nil {
		log.Fatal("Error inserting Mapper Data.", err)
	}
	fmt.Println("Inserted Mapper Data with ID:", inserted.InsertedID)
}
