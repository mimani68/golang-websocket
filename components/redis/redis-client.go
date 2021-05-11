package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	l "blackoak.cloud/balout/v2/components/log"
	config "blackoak.cloud/balout/v2/config"
)

var ctx = context.Background()

// Singleton redis connection string
var Rdb *redis.Client

//
// Redis Client
//
func RedisClient() *redis.Client {
	l.Log("[DEBUG] " + config.REDIS_URL + ":" + config.REDIS_PORT)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_URL + ":" + config.REDIS_PORT,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})
	return Rdb
}

func SetKV(key string, inputStruct map[string]interface{}, expireTimeInSecond int) bool {
	var a []byte
	// fmt.Println(inputStruct["data"])
	// fmt.Printf("%T", inputStruct["data"])
	a, _ = json.Marshal(inputStruct)
	if fmt.Sprintf("%T", inputStruct["data"]) == "struct" {
		a, _ = json.Marshal(inputStruct["data"])
	}
	set, err := Rdb.SetNX(ctx, key, string(a), time.Duration(expireTimeInSecond)*time.Second).Result()
	// err := rdb.Set(ctx, key, a, -1).Err()
	if err != nil {
		// panic(err)
		return set
	}
	return set
}

func GetKVJson(key string) (bool, interface{}) {
	value, err := Rdb.Get(ctx, key).Result()
	if err != nil {
		//
		// 01
		//
		// panic(err)

		//
		// 02
		//
		// var a struct {
		// 	Status string
		// } = struct{ Status string }{
		// 	Status: "Failure",
		// }
		// return false, a

		//
		// 03
		//
		type aType struct {
			Status string
		}
		return false, &aType{}
	}
	var a map[string]interface{}
	_ = json.Unmarshal([]byte(value), &a)
	return true, a
}
