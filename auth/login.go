package auth

import (
	"context"
	"fmt"
	"errors"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"project/models"
	db "project/database"
)

var mainuser = models.User{
	Username: "asad",
	Password: "1234",
	FirstName: "asad",
	LastName: "ahmed",
}

func Register(user models.User)  (error){
	// var res models.ResponseResult
	var result models.User

	collection, err := db.GetDBCollection("users")
	if err != nil {
		return err
	}
	
	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
			if err != nil {
				return errors.New("Error While Hashing Password, Try Again")
			}
			user.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				return errors.New("Error While Creating User, Try Again")
			}
			// res.Result = "Registration Successful"
			return nil
		}
		return err
	}

	// res.Result = "Username already Exists"
	return errors.New("Username already Exists")
}

func Login(username, password string) (models.User, error){
	var result models.User

	collection, err := db.GetDBCollection("users")

	err = collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		return result, errors.New("Invalid username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return result, errors.New("Invalid password")
	}

	claims := &models.Claims{Username: result.Username}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return result, errors.New("Error while generating token")
	}

	result.Token = tokenString
	result.Password = ""

	return result, nil

}

func Profile(username string) (models.User){
	var result models.User

	result.Username = username
	err := db.GetCache(db.CacheUser(result))

	if(err!=nil){
		collection, err := db.GetDBCollection("users")
		fmt.Println(err)
		err = collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
		result.Password = ""

		err = db.SetCache(db.CacheUser(result))
		if err != nil{
			log.Println(err)
		}
	}

	return result
}

func LoginStatic(username, password string) (models.User, error){
	var result models.User
	fmt.Println(username, password)
	if username != mainuser.Username || password != mainuser.Password{
		return result, errors.New("Invalid Credentials")
	}
	result = mainuser
	claims := &models.Claims{Username: result.Username}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return result, errors.New("Error while generating token")
	}

	result.Token = tokenString
	result.Password = ""

	return result, nil

}

func ProfileStatic(username string) (models.User){
	var result models.User
	result = mainuser

	return result
}