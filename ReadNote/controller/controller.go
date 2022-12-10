package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Chief-Rishab/mymodule/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var collection *mongo.Collection

// connection with the database

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")

	clientOption := options.Client().ApplyURI(connectionString)
	dbclient, err := mongo.Connect(context.Background(), clientOption) // connection request

	if err != nil {
		log.Fatal(err)
	}
	collection = (*mongo.Collection)(dbclient.Database(dbName).Collection(collectionName)) // Collection reference is listening
}

func init() {
	loadTheEnv()
	createDBInstance()
}

func createArticle(article model.Article) {
	res, err := collection.InsertOne(context.Background(), article)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Article created successfully with ID: ", res.InsertedID)
}

// filter(by ID) and update(from body), MongoDB creates a record with key as _id
// Update Article marks it as read

func updateArticle(articleID string) {
	// to convert articleID to make MongoDB understand we convert string to objectID
	id, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		log.Fatal("Error in converting articleID")
	}

	filter := bson.M{"_id": id} // query parameters in parenthesis
	action := bson.M{"$set": bson.M{"read": true}}

	res, err := collection.UpdateOne(context.Background(), filter, action)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of articles updated: ", res.ModifiedCount)
}

// delete one with given articleID
func deleteArticle(articleID string) {
	id, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		log.Fatal("Error in converting articleID")
	}
	filter := bson.M{"_id": id} // query parameters in parenthesis

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of articles deleted: ", res.DeletedCount)
}

// delete all articles
func deleteAllArticles() {

	filter := bson.D{{}}

	res, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of articles deleted: ", res.DeletedCount)
}

// cursor contains the results, and we can loop through it to get the result
func getAllArticles() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var articles []primitive.M

	for cursor.Next(context.Background()) {
		var article bson.M
		err := cursor.Decode(&article)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}

	return articles
}

func getArticle(articleID string) primitive.M {

	id, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		log.Fatal("Error in converting articleID")
	}

	filter := bson.M{"_id": id} // query parameters in parenthesis

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var article bson.M
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&article)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer cursor.Close(context.Background())

	return article
}

// controllers
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := getAllArticles()
	json.NewEncoder(w).Encode(articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	link, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(link)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)

	str := fmt.Sprintf("%v", jsonMap["url"])
	url := "http://api.linkpreview.net?key=37b4140b084c83d985e6c93124ac530c&q=" + str

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	s := string(data)
	var m model.Message
	errm := json.Unmarshal([]byte(s), &m)
	if errm != nil {
		log.Fatal(err)
	}

	var article model.Article
	article.Title = m.Title
	article.Description = m.Description
	article.ImageURL = m.Image
	article.URL = m.URL

	createArticle(article)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	updateArticle(id)
	article := getArticle(id)

	json.NewEncoder(w).Encode(article)
}

func DeleteAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := getAllArticles()

	deleteAllArticles()

	json.NewEncoder(w).Encode(articles)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	article := getArticle(id)

	deleteArticle(id)
	json.NewEncoder(w).Encode(article)
}
