package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main () {

    fmt.Println("[+] Starting the second server server")

    http.HandleFunc("/",index)
    http.HandleFunc("/test-api",testApi)

    fmt.Println("[+] The second server is listening to server.Running on port => " + port())

    res := http.ListenAndServe(port(), nil)

    fmt.Println(res)
}

func port()  string {
    port := "8877"

    return ":" + port
}


func testApi(w http.ResponseWriter,r *http.Request) {
    url := "https://jsonplaceholder.typicode.com/todos/1"

    req, err := http.NewRequest(http.MethodGet, url, nil)

    w.Header().Add("Content-Type","application/json, charset=utf-8")

    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        err = json.Marshal(err)
        w.Write(err)
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type","application/json, charset=utf-8")

    _, _ = w.Write(res)


}

func index(w http.ResponseWriter,r *http.Request)  {
    res, _ := json.Marshal("Hello from second service.")

    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type","application/json, charset=utf-8")
    _, _ = w.Write(res)
}
