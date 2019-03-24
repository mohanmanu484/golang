package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":3000",nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Hello heroku")
}
