package papergirl

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComic(t *testing.T) {
	initializeTest()
	page1 := BookPage("B.jpg")
	page2 := BookPage("A.jpg")
	comicBookContentStreamer = NewMockComicBookContentStreamer([]BookPage{page1, page2})

	path := Path("path/book.cbz")
	bookStorage.Store(path, bytes.NewReader([]byte{1, 2, 3, 4}))

	err := UpsertBook(path)
	assert.Nil(t, err)

	allBooks, _ := books.FindAll()
	assert.Len(t, allBooks, 1)

	createdBook := allBooks[0]
	assert.Equal(t, BookTitle("book"), createdBook.Title)
	assert.Equal(t, path, createdBook.Path)
	assert.Equal(t, CBZ, createdBook.Type)
	assert.Len(t, createdBook.Pages, 2)
	assert.Equal(t, page2, createdBook.Pages[0])
	assert.Equal(t, page1, createdBook.Pages[1])
}
