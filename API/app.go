package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tokopedia/perpustakaan/API/model"
	"github.com/tokopedia/perpustakaan/API/repository/author"

	"github.com/tokopedia/perpustakaan/API/repository/book"

	"github.com/tokopedia/perpustakaan/API/repository/bookauthor"

	"github.com/tokopedia/perpustakaan/API/repository/publisher"

	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/perpustakaan/API/domain/bookresponse"

	"github.com/tokopedia/perpustakaan/API/database"
	gcfg "gopkg.in/gcfg.v1"
)

var configuration struct {
	Database database.Options
}

func main() {
	err := gcfg.ReadFileInto(&configuration, "files/etc/perpustakaan/perpustakaan.ini")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	database.InitConnection(configuration.Database)
	//init resource
	bookauthor.New(database.Conn["perpustakaan"])
	book.NewBookRes(database.Conn["perpustakaan"])
	author.NewAuthorRes(database.Conn["perpustakaan"])
	publisher.NewPublisherRes(database.Conn["perpustakaan"])

	//init domain
	bookresponse.NewBookRepsonse(book.BookRes, author.AuthorRes, publisher.PublisherRes, bookauthor.BookAuthorRes)

	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := bookresponse.BookResponse.GetBookResponse(context.Background(), model.Filter{})
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	res, err := json.Marshal(&result)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
