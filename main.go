package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

)

func main() {

	port := getPort()


	/*r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)*/
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	port := os.Getenv("PORT")
	fmt.Println("port is " + port)
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	return port
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Hello heroku")
}
