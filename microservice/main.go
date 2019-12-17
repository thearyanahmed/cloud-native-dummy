package main

import (
	"encoding/json"
	"fmt"
	"github.com/thearyanahmed/cloud-native-dummy/api"
	"net/http"
	"os"
)

func main () {

	fmt.Println("[+] Starting server")

	http.HandleFunc("/",index)
	http.HandleFunc("/api/books",api.Books)


	fmt.Println("[+] Listening to server.Running on port => " + port())

	res := http.ListenAndServe(port(), nil)

	fmt.Println(res)
}

func port()  string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "9988"
	}

	return ":" + port
}


func index(w http.ResponseWriter,r *http.Request)  {
	res, _ := json.Marshal("Hello world from 'go container'.")

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type","application/json, charset=utf-8")
	_, _ = w.Write(res)
}
