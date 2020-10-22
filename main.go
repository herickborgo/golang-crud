package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/router"

	"github.com/joho/godotenv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load(os.Getenv("GOPATH") + "/src/rest-api/.env")
	checkError(err)

	router := router.GetRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
