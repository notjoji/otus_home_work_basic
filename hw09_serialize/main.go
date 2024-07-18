package main

import (
	"encoding/json"
	"fmt"
	"github.com/notjoji/otus_home_work_basic/hw09_serialize/book"
	"log"
)

type Book struct {
	Id     int32   `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author string  `json:"author,omitempty"`
	Year   int32   `json:"year,omitempty"`
	Size   int32   `json:"size,omitempty"`
	Rate   float32 `json:"rate,omitempty"`
}

func (b Book) String() string {
	return fmt.Sprintf("id: %d title: %s author: %s year: %d size: %d rate: %f", b.Id, b.Title, b.Author, b.Year, b.Size, b.Rate)
}

func Map(b *book.Book) Book {
	return Book{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	}
}

func (b Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b Book) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b)
}

func MarshalToJSON(arr []Book) ([]byte, error) {
	//data := make([]byte, len(arr))
	return json.Marshal(arr)
	//for i, elem := range arr {
	//	marshalled, err := elem.MarshalJSON()
	//	if err != nil {
	//		return nil, err
	//	}
	//	data[i] = marshalled
	//}
	//return data, nil
}

func UnmarshalFromJSON(data []byte) ([]*Book, error) {
	//res := make([]*MyBook, len(data))
	//for i, elem := range data {
	//	b := MyBook{}
	//	err := b.UnmarshalJSON(elem)
	//	if err != nil {
	//		return nil, err
	//	}
	//	res[i] = &b
	//}
	//return res, nil
	res := make([]*Book, 0)
	err := json.Unmarshal(data, &res)
	return res, err
}

func main() {
	b1 := Book{
		Id:     1,
		Title:  "A Game Of Thrones (A Song of Ice and Fire)",
		Author: "George RR Martin",
		Year:   1997,
		Size:   694,
		Rate:   9.1,
	}
	b2 := Book{
		Id:     2,
		Title:  "The Lord of the Rings",
		Author: "John R.R. Tolkien",
		Year:   1954,
		Size:   1820,
		Rate:   9.3,
	}
	books := []Book{b1, b2}
	marshalled, err := MarshalToJSON(books)
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
	for _, b := range unmarshalled {
		fmt.Println("Book:", b.String())
	}
}
