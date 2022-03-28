package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/model"
	redis "go-auth/redis"
	redisuser "go-auth/redis/users"
	"net/http"
)

/*
Request Body must be json
Example Request Body Json:
{
    "username":"test",
    "password": "test"
}
*/
func LoginRequestBodyJson(w http.ResponseWriter, r *http.Request) {
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

		password := user.Password //before we get from the cache we are take to variable right for compare

		user = redisuser.GetUser(user.Username) //get from the cache

		if password == user.Password {
			response = model.ResponseModel{
				ResCode:    1,
				ResMessage: "Success",
				ResData:    user,
			}
		} else {
			response = model.ResponseModel{
				ResCode:    -1,
				ResMessage: "Error",
				ResData:    "Wrong Password",
			}
		}

		json, _ := json.Marshal(response)

		fmt.Fprintln(w, string(json))
	}
}

func LoginRequestFormData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.UserModel
		var response model.ResponseModel
		w.Header().Set("Content-Type", "application/json")
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

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

		if username == "" || password == "" {
			response = model.ResponseModel{
				ResCode:    -1,
				ResMessage: "Error",
				ResData:    "Username or Password Can Not Be Empty ",
			}
			json, _ := json.Marshal(response)

			http.Error(w, string(json), http.StatusBadRequest)
			return
		}

		user = redisuser.GetUser(username) //get from the cache

		if password == user.Password {
			response = model.ResponseModel{
				ResCode:    1,
				ResMessage: "Success",
				ResData:    user,
			}
		} else {
			response = model.ResponseModel{
				ResCode:    -1,
				ResMessage: "Error",
				ResData:    "Wrong Password",
			}
		}

		json, _ := json.Marshal(response)

		fmt.Fprintln(w, string(json))
	}
}
