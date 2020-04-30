package todo

import (
	"context"
	// "fmt"
	"errors"
	// "log"

	"go.mongodb.org/mongo-driver/bson"

	"project/models"
	db "project/database"
)

func Insert(todo models.ToDo)  (models.ResponseResult, error){
	var res models.ResponseResult
	// var result models.ToDo

	collection, err := db.GetDBCollection("todo")
	if err != nil {
		return res, err
	}
	
	_, err = collection.InsertOne(context.TODO(), todo)
	if err != nil {
		return res, errors.New("Error While Inserting, Try Again")
	}
	res.Result = "200"
	return res, nil
}

func Fetch(username string) ([]models.ToDo, error){
	var result []models.ToDo

	collection, err := db.GetDBCollection("todo")
	if err != nil {
		return result, err
	}

	// type M map[string]interface{}

	cur, err := collection.Find(context.TODO(), bson.M{"username": username})
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()){
		var tempTodo models.ToDo

		err := cur.Decode(&tempTodo)
		if err != nil{
			return result, err
		}
		result = append(result, tempTodo)
	}
	// err = collection.Find(M{"username":username}).All(&result)
	if err != nil {
	    return result, errors.New("Cannot Retrive")
	}
	
	return result, nil
}