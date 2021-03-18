package main

import (
	"fmt"
	"log"
	"net/http"

	"api-dev-house/src/router"
)

func main() {
	fmt.Println("connect to api-dev-house")

	router := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
