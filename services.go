package papergirl

import "io"

var bookTypeIdentifier BookTypeIdentifier
var comicBookContentStreamer ComicBookContentStreamer

type BookTypeIdentifier interface {
	Identify(reader io.Reader) (BookType, error)
}

type ComicBookContentStreamer interface {
	Files(reader io.Reader) ([]BookPage, error)
	Stream(reader io.Reader, page BookPage, consumer func(reader io.Reader) error) error
}
