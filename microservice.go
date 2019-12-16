package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thearyanahmed/cloud-native-dummy/api"
	"log"
	"net/http"
	"os"
)

func main () {

	fmt.Println("Starting server")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}


	http.HandleFunc("/",index)
	http.HandleFunc("/api/books",api.Books)


	fmt.Println("Listening to server.Running on port => " + port())

	res := http.ListenAndServe(port(), nil)

	fmt.Println(res)
}

func port()  string {
	port := os.Getenv("PORT")

	fmt.Println("GOT ENV PORT",port)

	if port == "" {
		port = "8080"
	}

	return ":" + port
}


func index(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Index struct")
}