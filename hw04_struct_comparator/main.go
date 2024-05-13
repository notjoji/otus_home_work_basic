package main

import "fmt"

type Book struct {
	id     int32
	title  string
	author string
	year   int32
	size   int32
	rate   float32
}

func (b *Book) ID() int32 {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int32 {
	return b.year
}

func (b *Book) Size() int32 {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetID(id int32) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int32) {
	b.year = year
}

func (b *Book) SetSize(size int32) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

type CompareBy int8

const (
	Year CompareBy = iota
	Size
	Rate
)

type BookComparator struct {
	compareBy CompareBy
}

func NewBookComparator(compareBy CompareBy) *BookComparator {
	return &BookComparator{compareBy}
}

func (bc *BookComparator) Compare(book1 *Book, book2 *Book) bool {
	switch bc.compareBy {
	case Year:
		return book1.Year() > book2.Year()
	case Size:
		return book1.Size() > book2.Size()
	case Rate:
		return book1.Rate() > book2.Rate()
	}
	return false
}

func main() {
	book1 := Book{
		1,
		"A Game Of Thrones (A Song of Ice and Fire)",
		"George RR Martin",
		1997,
		694,
		9.1,
	}
	book2 := Book{
		2,
		"The Lord of the Rings",
		"John RR Tolkien",
		1954,
		1077,
		9.3,
	}
	fmt.Printf("1st book is %s by %s\n", book1.Title(), book1.Author())
	fmt.Printf("2nd book is %s by %s\n", book2.Title(), book2.Author())
	fmt.Println("Is 1st book newer than 2nd? ", NewBookComparator(Year).Compare(&book1, &book2))
	fmt.Println("Is 1st book larger than 2nd? ", NewBookComparator(Size).Compare(&book1, &book2))
	fmt.Println("Has 1st book better rating than 2nd? ", NewBookComparator(Rate).Compare(&book1, &book2))
}
