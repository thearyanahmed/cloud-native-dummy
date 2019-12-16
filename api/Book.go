package api

import (
	json2 "encoding/json"
	"net/http"
)

/**
 * Book structure
 */

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

var BookList = []Book {
	Book {
		Title: "The Alchemist",
		Author: "Paulo Coelho",
		ISBN: "123222121221",
	},
	Book {
		Title: "The First Muslim",
		Author: "Lezley Hazleton",
		ISBN: "12312312",
	},
}

/**
 * Convert a Book struct to json
 */
func (b Book) JsonEncode ()[]byte {

	json, err:= json2.Marshal(b)

	if err != nil {
		panic(err)
	}

	return json
}

/**
 * Decode json to book structure
 */
func JsonDecode(json []byte) Book {
	book := Book{}

	err := json2.Unmarshal(json,&book)

	if err != nil {
		panic(err)
	}

	return book
}

func GetBookList() []Book {
	return BookList
}


func Books (w http.ResponseWriter, r *http.Request) {

	books , err := json2.Marshal(BookList)

	if err != nil {
		panic(err)
	}
	//w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type","application/json; charset=utf-8")
	_, _ = w.Write(books)
}

