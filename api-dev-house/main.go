package main

import (
	"fmt"
	"log"
	"net/http"

	"api-dev-house/src/router"

	"api-dev-house/src/config"
)

func main() {

	config.Load()

	fmt.Println("connect to api-dev-house")
	fmt.Printf("Listen on localhost: %d", config.Port)

	router := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))

}
