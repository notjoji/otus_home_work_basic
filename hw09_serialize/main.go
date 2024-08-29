package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/notjoji/otus_home_work_basic/hw09_serialize/book"
)

type Book struct {
	ID     int32   `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author string  `json:"author,omitempty"`
	Year   int32   `json:"year,omitempty"`
	Size   int32   `json:"size,omitempty"`
	Rate   float32 `json:"rate,omitempty"`
}

func (b *Book) String() string {
	return fmt.Sprintf("id: %d title: %s author: %s year: %d size: %d rate: %f", b.ID, b.Title, b.Author, b.Year,
		b.Size, b.Rate)
}

func MapFromProto(b *book.Book) Book {
	return Book{
		ID:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	}
}

func MapToProto(b *Book) *book.Book {
	return &book.Book{
		Id:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	}
}

func (b *Book) MarshalJSON() ([]byte, error) {
	protoBook := MapToProto(b)
	return json.Marshal(&protoBook)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var b1 book.Book
	err := json.Unmarshal(data, &b1)
	if err != nil {
		return err
	}
	*b = MapFromProto(&b1)
	return nil
}

func MarshalToJSON(protoBooks *book.Books) ([]byte, error) {
	books := make([]Book, len(protoBooks.Books))
	for i, b := range protoBooks.Books {
		books[i] = MapFromProto(b)
	}
	return json.Marshal(books)
}

func UnmarshalFromJSON(data []byte) (book.Books, error) {
	books := make([]*Book, 0)
	err := json.Unmarshal(data, &books)

	res := make([]*book.Book, len(books))
	for i, b := range books {
		res[i] = MapToProto(b)
	}
	return book.Books{Books: res}, err
}

func main() {
	b1 := &book.Book{
		Id:     1,
		Title:  "A Game Of Thrones (A Song of Ice and Fire)",
		Author: "George RR Martin",
		Year:   1997,
		Size:   694,
		Rate:   9.1,
	}
	b2 := &book.Book{
		Id:     2,
		Title:  "The Lord of the Rings",
		Author: "John R.R. Tolkien",
		Year:   1954,
		Size:   1820,
		Rate:   9.3,
	}
	books := []*book.Book{b1, b2}
	protoBooks := &book.Books{Books: books}
	marshalled, err := MarshalToJSON(protoBooks)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Length:", len(marshalled), "Marshalled data:", string(marshalled))

	unmarshalled, err := UnmarshalFromJSON(marshalled)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, b := range unmarshalled.GetBooks() {
		fmt.Println("Book:", b.String())
	}
}
