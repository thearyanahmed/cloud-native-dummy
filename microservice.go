package main

import (
	"fmt"
	"net/http"
	"os"
	"api"
)

func main () {

	http.HandleFunc("/",index)
	http.HandleFunc("/api/books",api.Books)


	fmt.Println("Listening to server.Running on port : " + port())

	_ = http.ListenAndServe(port(), nil)

}

func port()  string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}


func index(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Index struct")
}