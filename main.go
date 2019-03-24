package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		//port="8000"
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/",handler)
	http.ListenAndServe(":"+port,nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Hello heroku")
}
