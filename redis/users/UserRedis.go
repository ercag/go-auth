package users

import (
	"encoding/json"
	"fmt"
	"go-auth/crypto"
	"go-auth/guid"
	models "go-auth/model"
	"go-auth/redis"
)

func GetUser(key string) models.UserModel {
	client, err := redis.RedisClient()
	var retval models.UserModel

	if err == nil {
		val, err := client.Get(key).Result()
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(val), &retval)

		if err != nil {
			panic(err)
		}

		retval.Password = crypto.CryptMD5(retval.Password)
	}
	return retval
}

func SetUser(user models.UserModel) (string, error) {
	client, err := redis.RedisClient()

	if err != nil {
		return "Error try to connect Redis Server", err

	}

	json, err := json.Marshal(user)
	if err != nil {
		return "Error try to convert to json", err

	}

	err = client.Set(user.Username, json, 0).Err()
	if err != nil {
		return "Error try to set the key for redis cache", err
	}

	return "added to the redis cache with: " + user.Username + " key", err
}

func SetUsers(users []models.UserModel) {
	client, err := redis.RedisClient()

	if err == nil {
		for _, v := range users {
			v.Guid = guid.CreateGuid()

			json, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
			}

			err = client.Set(v.Username, json, 0).Err()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("added to the redis cache with: " + v.Username + " key")
		}
	}
}
