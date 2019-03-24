package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/app",handler)
	http.ListenAndServe(":8000",nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Hello heroku")
}
