package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tokopedia/perpustakaan/api/common/apiwriter"
	"github.com/tokopedia/perpustakaan/api/model"
	"github.com/tokopedia/perpustakaan/api/repository/author"

	"github.com/tokopedia/perpustakaan/api/repository/book"

	"github.com/tokopedia/perpustakaan/api/repository/bookauthor"

	"github.com/tokopedia/perpustakaan/api/repository/publisher"

	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/perpustakaan/api/domain/bookresponse"

	"github.com/tokopedia/perpustakaan/api/database"
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
	apiwriter.NewAPIWriter()
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
		apiwriter.ApiWriter.WriteFailResp(w, err)
		return
	}

	apiwriter.ApiWriter.WriteSuccesaResp(w, result)
	return
}
