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
	router := router.GenerateRouter()

	fmt.Printf("Listen on localhost: %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))

}
