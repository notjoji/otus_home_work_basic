package main

import (
	"encoding/json"
	"fmt"
	"github.com/notjoji/otus_home_work_basic/hw09_serialize/book"
)

type Book book.Book

func NewBook() *Book {
	return &Book{
		Id:     1,
		Title:  "The Lord of the Rings",
		Author: "John R.R. Tolkien",
		Year:   1954,
		Size:   1820,
		Rate:   9.3,
	}
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b)
}

func MarshalToJSON(arr []*Book) ([]byte, error) {
	var bytes []byte
	for _, elem := range arr {
		b, err := elem.MarshalJSON()
		if err != nil {
			bytes = append(bytes, b...)
		} else {
			return nil, err
		}
	}
	return bytes, nil
}

func main() {
	books := []*Book{NewBook(), NewBook()}
	bytes, err := MarshalToJSON(books)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
