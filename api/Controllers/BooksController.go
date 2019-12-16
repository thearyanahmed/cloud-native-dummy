package Controllers

import (
	"../Book"
	"encoding/json"
	"net/http"
)

func Books (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	books , err := json.Marshal(Book.GetBookList())

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type","application/json; charset=utf-8")
	_, _ = w.Write(books)
}