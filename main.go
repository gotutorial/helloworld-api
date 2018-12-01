package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"io/ioutil"
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
	router.HandleFunc("/swagger.json", GetSwagger).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(getPort(), handlers.CORS(originsOk)(router)))

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	greeting := &Greeting{
		Statement: "Hello World!",
	}
	json.NewEncoder(w).Encode(greeting)
}

func GetSwagger(w http.ResponseWriter, r *http.Request) {

	b,_ := ioutil.ReadFile("swagger.json");

	rawIn := json.RawMessage(string(b))
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(rawIn, &objmap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(objmap)

	json.NewEncoder(w).Encode(objmap)
}


func getPort() string {
	var port string
	if configuredPort := os.Getenv("PORT"); configuredPort == "" {
		port = "3000"
	} else {
		port = configuredPort
	}

	return fmt.Sprintf(":%v", port)
}

