package papergirl

import (
	"io"
)

type MockBookTypeIdentifier struct {
	bookType BookType
	err      error
}

func NewMockBookTypeIdentifier(bookType BookType, err error) BookTypeIdentifier {
	return &MockBookTypeIdentifier{
		bookType: bookType,
		err:      err,
	}
}

func (identifier *MockBookTypeIdentifier) Identify(reader io.Reader) (BookType, error) {
	return identifier.bookType, identifier.err
}

type MockComicBookContentStreamer struct {
	pages []BookPage
	err   error
}

func NewMockComicBookContentStreamer(pages []BookPage) ComicBookContentStreamer {
	return &MockComicBookContentStreamer{
		pages: pages,
	}
}

func (streamer *MockComicBookContentStreamer) Files(reader io.Reader) ([]BookPage, error) {
	return streamer.pages, streamer.err
}

func (*MockComicBookContentStreamer) Stream(reader io.Reader, page BookPage, consumer func(reader io.Reader) error) error {
	panic("unimplemented")
}

type MockPasswordHasher struct {
}

func NewMockPasswordHasher() PasswordHasher {
	return &MockPasswordHasher{}
}

func (*MockPasswordHasher) Hash(password Password) (Password, error) {
	var result string
	for _, ch := range password {
		result = string(ch) + result
	}

	return Password(result), nil
}
