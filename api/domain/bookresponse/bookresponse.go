package bookresponse

import (
	"context"

	"github.com/tokopedia/perpustakaan/api/model"
	"github.com/tokopedia/perpustakaan/api/repository/author"
	"github.com/tokopedia/perpustakaan/api/repository/book"
	"github.com/tokopedia/perpustakaan/api/repository/bookauthor"
	"github.com/tokopedia/perpustakaan/api/repository/publisher"
)

var (
	BookResponse BookResponseInf
)

type BookResponseImpl struct {
	BookRes       book.BookInf
	AuthorRes     author.AuthorInf
	PublisherRes  publisher.PublisherInf
	BookAuthorRes bookauthor.BookAuthorInf
}

func NewBookRepsonse(Bookres book.BookInf, AuthorRes author.AuthorInf, PublisherRes publisher.PublisherInf, BookAuthorRes bookauthor.BookAuthorInf) {
	BookResponse = &BookResponseImpl{
		BookRes:       Bookres,
		AuthorRes:     AuthorRes,
		PublisherRes:  PublisherRes,
		BookAuthorRes: BookAuthorRes,
	}
}

func (br *BookResponseImpl) GetBookResponse(ctx context.Context, filter model.Filter) ([]model.BookResponse, error) {
	var result []model.BookResponse
	where := getFilter(filter)
	bookIds, err := br.BookRes.GetFileteredBooksIds(ctx, where)
	if err != nil {
		return result, err
	}

	books, err := br.BookRes.GetBooks(ctx, bookIds)
	if err != nil {
		return result, err
	}

	bookAuthor, err := br.BookAuthorRes.GetBookAuthor(ctx, bookIds)
	if err != nil {
		return result, err
	}

	//get list of authorid
	authorIds := getListOfAuthorID(bookAuthor)
	author, err := br.AuthorRes.GetAuthor(ctx, authorIds)
	if err != nil {
		return result, err
	}

	//get list of publisherid
	publisherIds := getListOfPublisherID(books)
	publisher, err := br.PublisherRes.GetPublishers(ctx, publisherIds)
	if err != nil {
		return result, err
	}

	for _, book := range books {
		var authors []model.Author
		for _, bookAuthor := range bookAuthor[book.BookID] {
			authors = append(authors, author[bookAuthor.AuthorID])
		}

		result = append(result, model.BookResponse{
			book,
			authors,
			publisher[book.PublisherID],
		})
	}

	return result, nil
}

func getFilter(filter model.Filter) string {
	result := "where 1=1 "

	result += "limit 10"

	return result
}

func getListOfAuthorID(bookAuthor map[int64][]model.BooksAuthor) []int64 {
	var result []int64

	for _, bookAuthors := range bookAuthor {
		for _, bookauthor := range bookAuthors {
			result = append(result, bookauthor.AuthorID)
		}
	}

	return result
}

func getListOfPublisherID(books map[int64]model.Book) []int64 {
	var result []int64
	for _, book := range books {
		result = append(result, book.PublisherID)
	}

	return result
}
