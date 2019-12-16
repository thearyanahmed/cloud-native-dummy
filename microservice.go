package main

import (
	"fmt"
	"net/http"
	"os"
)

func main () {

	http.HandleFunc("/",index)
	fmt.Println("Listening to server")
	res := http.ListenAndServe(port(), nil)
	fmt.Println(res)
}

func port()  string {
	port := os.Getenv("PORT")

	if port != "" {
		port = "8000"
	}

	return ":" + port
}


func index(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Index struct")
}