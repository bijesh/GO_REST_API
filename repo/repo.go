package repo
 
import (
	"fmt"
	"context"
	"log"
	"errors"
	"GO_REST_API/models"
	
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
 

func FindBlog(title string) models.Blog {
	var result models.Blog
	filter := bson.D{{"title", title}}
	collection := GetMongoCollection()
	// collection.FindOne(context.TODO(), filter).Decode(&result)
	
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
 
func FindAllBlog() models.Blogs {
	var results models.Blogs
	findOptions := options.Find()
	findOptions.SetLimit(10)
	collection := GetMongoCollection()
	cur, err := collection.Find(context.TODO(), nil, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var blog models.Blog
		err := cur.Decode(&blog)
		if err != nil {
			log.Fatal(err)
		}
		
		// results = append(results, &blog)
	}
	
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	
	// Close the cursor once finished
	cur.Close(context.TODO())
	return results
}

func CreateBlog(t models.Blog) error {
   
	collection := GetMongoCollection()
	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
		return errors.New("insert failed")
	}
	
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    return nil
}

func GetMongoCollection() *mongo.Collection {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("Blogs")

	return collection
}