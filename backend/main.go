package main

import (
	"api/controller"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("API")

	mux := http.NewServeMux()
	fmt.Println("api")
	mux.HandleFunc("/hello", controller.CreateUser)

	if err := http.ListenAndServe("localhost:8001", mux); err != nil {
		fmt.Println(err.Error())
	}
}
