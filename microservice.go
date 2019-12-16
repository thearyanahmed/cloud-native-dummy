package main

import (
	"./api/Controllers"
	"fmt"
	"net/http"
	"os"
)

func main () {

	http.HandleFunc("/",index)
	http.HandleFunc("/api/books",Controllers.Books)


	fmt.Println("Listening to server.Running on port : " + port())

	_ = http.ListenAndServe(port(), nil)

}

func port()  string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	return ":" + port
}


func index(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Index struct")
}