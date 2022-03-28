package webserver

import (
	"fmt"
	"go-auth/web-server/handlers"
	"log"
	"net/http"
)

func Start(port string) {

	http.HandleFunc("/loginrbj", handlers.LoginRequestBodyJson)
	http.HandleFunc("/loginrfd", handlers.LoginRequestFormData)
	http.HandleFunc("/adduser", handlers.AddUser)

	fmt.Println(fmt.Sprint("Server will start at:", port))

	err := http.ListenAndServe(fmt.Sprint(":", port), nil)

	log.Fatal(err)

	fmt.Println(fmt.Sprint("Server started at:", port))
}
