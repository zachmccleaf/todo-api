package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zachmccleaf/todo-api/helper"
	"github.com/zachmccleaf/todo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDB()

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var todos []models.Todo

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var todo models.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var todo models.Todo
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&todo)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Range, Content-Disposition, Content-Description,Origin, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var todo models.Todo

	_ = json.NewDecoder(r.Body).Decode(&todo)

	result, err := collection.InsertOne(context.TODO(), todo)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Range, Content-Disposition, Content-Description,Origin, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var todo models.Todo

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&todo)

	update := bson.D{
		{"$set", bson.D{
			{"isbn", todo.Isbn},
			{"message", todo.Message},
			{"color", todo.Color},
			{"isComplete", todo.IsComplete},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&todo)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	todo.ID = id

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodos(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	r.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With, Accept")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
