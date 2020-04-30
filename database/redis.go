package database

import(
	"encoding/json"
	"fmt"
	"log"
	"errors"

	"github.com/gomodule/redigo/redis"

	"project/models"

)

type CacheLayer interface {
	GetCacheStruct(redis.Conn) error
	SetCacheStruct(redis.Conn) error
}

type CacheUser models.User

func GetCache(cache CacheLayer) error{
	pool := newPool()

	conn := pool.Get()
	defer conn.Close()

	err := ping(conn)
	if err != nil {
		log.Println(err)
		return err
	}

	err = cache.GetCacheStruct(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func SetCache(cache CacheLayer) error{
	pool := newPool()

	conn := pool.Get()
	defer conn.Close()

	err := ping(conn)
	if err != nil {
		log.Println(err)
		return err
	}

	err = cache.SetCacheStruct(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	return nil
}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

func (user CacheUser) GetCacheStruct(c redis.Conn) error{
	const objectPrefix string = "user:"
	
	s, err := redis.String(c.Do("GET", objectPrefix+user.Username))
	if err == redis.ErrNil {
		log.Println("User does not exist")
		return errors.New("Does not exist")
	} else if err != nil {
		return err
	}

	// usr := User{}
	err = json.Unmarshal([]byte(s), &user)

	fmt.Printf("%+v\n", user)
	return nil
}

func (user CacheUser) SetCacheStruct(c redis.Conn) error{
	const objectPrefix string = "user:"

	// serialize User object to JSON
	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// SET object
	_, err = c.Do("SET", objectPrefix+user.Username, json)
	if err != nil {
		return err
	}
	return nil	
}