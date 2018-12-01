package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Greeting struct {
	Statement  string `json:"statement"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", sayHello).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", getPort()), router))

}

func sayHello(response http.ResponseWriter, request *http.Request) {
	greeting := &Greeting{
		Statement: "Hello World!",
	}
	json.NewEncoder(response).Encode(greeting)
}

func getPort() string {
	if configuredPort := os.Getenv("PORT"); configuredPort == "" {
		return "3000"
	} else {
		return configuredPort
	}
}

