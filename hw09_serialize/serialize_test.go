package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	b1 := Book{
		ID:     1,
		Title:  "A Game Of Thrones (A Song of Ice and Fire)",
		Author: "George RR Martin",
		Year:   1997,
		Size:   694,
		Rate:   9.1,
	}
	b2 := Book{
		ID:     2,
		Title:  "The Lord of the Rings",
		Author: "John R.R. Tolkien",
		Year:   1954,
		Size:   1820,
		Rate:   9.3,
	}
	books := []Book{b1, b2}
	marshalled, err := MarshalToJSON(books)
	assert.NoError(t, err)
	assert.NotEmpty(t, marshalled)
	assert.Equal(t,
		`[{"id":1,"title":"A Game Of Thrones (A Song of Ice and Fire)","author":"George RR Martin","year":1997,`+
			`"size":694,"rate":9.1},{"id":2,"title":"The Lord of the Rings","author":"John R.R. Tolkien","year":1954,`+
			`"size":1820,"rate":9.3}]`,
		string(marshalled))
	unmarshalled, err := UnmarshalFromJSON(marshalled)
	assert.NoError(t, err)
	assert.NotEmpty(t, marshalled)
	for i, book := range unmarshalled {
		assert.Equal(t, book.ID, books[i].ID)
		assert.Equal(t, book.Title, books[i].Title)
		assert.Equal(t, book.Author, books[i].Author)
		assert.Equal(t, book.Year, books[i].Year)
		assert.Equal(t, book.Size, books[i].Size)
		assert.Equal(t, book.Rate, books[i].Rate)
	}
}
