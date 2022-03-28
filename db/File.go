package db

import (
	"encoding/json"
	"fmt"
	"go-auth/model"
	"os"

	redisuser "go-auth/redis/users"
)

func Process() {
	readFile()
}

func readFile() {
	data, err := os.ReadFile("users.json")

	if err == nil {
		var users []model.UserModel

		err := json.Unmarshal(data, &users)
		if err != nil {
			panic(err)
		}
		redisuser.SetUsers(users)
	} else {
		fmt.Println(err)
	}
}
