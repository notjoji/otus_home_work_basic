package main

import (
	"fmt"

	"github.com/notjoji/otus_home_work_basic/hw09_serialize/book"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MyBook book.Book

func Map(b *MyBook) book.Book {
	return book.Book{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	}
}

func (b *MyBook) ProtoReflect() protoreflect.Message {
	b2 := Map(b)
	return b2.ProtoReflect()
}

func (b *MyBook) Reset() {
	*b = MyBook{}
}

func (b *MyBook) String() string {
	b2 := Map(b)
	return b2.String()
}

func (*MyBook) ProtoMessage() {
}

func (b *MyBook) MarshalJSON() ([]byte, error) {
	return proto.Marshal(b)
}

func (b *MyBook) UnmarshalJSON(data []byte) error {
	return proto.Unmarshal(data, b)
}

func MarshalToJSON(arr []*MyBook) ([][]byte, error) {
	data := make([][]byte, len(arr))
	for i, elem := range arr {
		marshalled, err := elem.MarshalJSON()
		if err != nil {
			return nil, err
		}
		data[i] = marshalled
	}
	return data, nil
}

func UnmarshalFromJSON(data [][]byte) ([]*MyBook, error) {
	res := make([]*MyBook, len(data))
	for i, elem := range data {
		b := MyBook{}
		err := b.UnmarshalJSON(elem)
		if err != nil {
			return nil, err
		}
		res[i] = &b
	}
	return res, nil
}

func main() {
	b1 := &MyBook{
		Id:     1,
		Title:  "A Game Of Thrones (A Song of Ice and Fire)",
		Author: "George RR Martin",
		Year:   1997,
		Size:   694,
		Rate:   9.1,
	}
	b2 := &MyBook{
		Id:     2,
		Title:  "The Lord of the Rings",
		Author: "John R.R. Tolkien",
		Year:   1954,
		Size:   1820,
		Rate:   9.3,
	}
	books := []*MyBook{b1, b2}
	marshalled, err := MarshalToJSON(books)
	if err != nil {
		fmt.Println(err)
		return
	}
	var allData []byte
	for _, b := range marshalled {
		allData = append(allData, b...)
	}
	fmt.Println("Length:", len(allData), "Marshalled data:", string(allData))

	unmarshalled, err := UnmarshalFromJSON(marshalled)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, b := range unmarshalled {
		fmt.Println("Book:", b.String())
	}
}
