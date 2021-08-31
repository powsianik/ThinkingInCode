package dbAccess

import (
	"context"
	"fmt"
	"github.com/powsianik/thinking-in-code/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var connectionString string = "mongodb+srv://Admin:QRUP)&$!qrup0741@thinkingincodeweb.ixhdb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func connect(uri string)(*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30 * time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func(){

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil{
			panic(err)
		}
	}()
}

func Ping(client *mongo.Client, ctx context.Context) error{

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occored, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func Write(post models.PostData){
	client, ctx, cancel, err := connect(connectionString)
	defer close(client, ctx, cancel)
	if err != nil {
		panic(err)
	}

	postsCollection := client.Database("thinkingInCodeBlog").Collection("posts")
	_, err = postsCollection.InsertOne(ctx, post)
	if err != nil {
		panic(err)
	}
}

func Read(key string, value interface{}) models.PostData{
	client, ctx, cancel, err := connect(connectionString)
	defer close(client, ctx, cancel)
	if err != nil {
		panic(err)
	}

	postsCollection := client.Database("thinkingInCodeBlog").Collection("posts")
	var postToRead models.PostData
	if err = postsCollection.FindOne(ctx, bson.M{key: value}).Decode(&postToRead); err != nil {
		log.Fatal(err)
	}

	return postToRead
}

func ReadAll() []models.PostData{
	client, ctx, cancel, err := connect(connectionString)
	defer close(client, ctx, cancel)
	if err != nil {
		panic(err)
	}

	postsCollection := client.Database("thinkingInCodeBlog").Collection("posts")

	cursor, err := postsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var posts []models.PostData
	if err = cursor.All(ctx, &posts); err != nil {
		log.Fatal(err)
	}

	return posts
}

func Update(data models.PostData){
	client, ctx, cancel, err := connect(connectionString)
	defer close(client, ctx, cancel)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": data.Id}
	postsCollection := client.Database("thinkingInCodeBlog").Collection("posts")

	var postToRead models.PostData
	if err = postsCollection.FindOne(ctx, bson.M{"_id": data.Id}).Decode(&postToRead); err != nil {
		log.Fatal(err)
	}

	result, err := postsCollection.ReplaceOne(ctx, filter, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func ClearCollection(){
	client, ctx, cancel, err := connect(connectionString)
	defer close(client, ctx, cancel)
	if err != nil {
		panic(err)
	}

	postsCollection := client.Database("thinkingInCodeBlog").Collection("posts")
	result, err := postsCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)
}