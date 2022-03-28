package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/guid"
	"go-auth/model"
	redis "go-auth/redis"
	redisuser "go-auth/redis/users"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.UserModel
		var response model.ResponseModel
		w.Header().Set("Content-Type", "application/json")

		_, err := redis.RedisClient()

		if err != nil {
			response = model.ResponseModel{
				ResCode:    -99,
				ResMessage: "Redis Error",
				ResData:    err.Error(),
			}
			json, _ := json.Marshal(response)

			http.Error(w, string(json), http.StatusBadRequest)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			response = model.ResponseModel{
				ResCode:    -3,
				ResMessage: "Error",
				ResData:    err.Error(),
			}
			json, _ := json.Marshal(response)

			http.Error(w, string(json), http.StatusBadRequest)
			return
		}

		if user.Username == "" || user.Password == "" {
			response = model.ResponseModel{
				ResCode:    -2,
				ResMessage: "Error",
				ResData:    "Username or Password Can Not Be Empty",
			}
			json, _ := json.Marshal(response)

			http.Error(w, string(json), http.StatusBadRequest)
			return
		}

		user.Guid = guid.CreateGuid()
		message, err := redisuser.SetUser(user)

		if err == nil {
			response = model.ResponseModel{
				ResCode:    1,
				ResMessage: "Success",
				ResData:    message,
			}
		} else {
			response = model.ResponseModel{
				ResCode:    -1,
				ResMessage: err.Error(),
				ResData:    message,
			}
		}
		json, _ := json.Marshal(response)

		fmt.Fprintln(w, string(json))
	}
}
