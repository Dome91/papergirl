package papergirl

import (
	"bytes"
	"errors"
	"io"
	_path "path"
	"sort"
	"strings"
)

var books Books
var bookStorage BookStorage

const (
	CBZ  BookType = "CBZ"
	EPUB BookType = "EPUB"
)

type BookTitle string
type BookType string
type BookPage string
type Book struct {
	Id      ID
	Title   BookTitle
	Path    Path
	Type    BookType
	Pages   []BookPage
	CoverId ID
}

func NewBook(path Path, bookType BookType, pages []BookPage) Book {
	base := _path.Base(string(path))
	title := BookTitle(strings.TrimSuffix(base, _path.Ext(base)))
	sort.Slice(pages, func(i, j int) bool {
		return pages[i] < pages[j]
	})

	return Book{
		Title: title,
		Path:  path,
		Type:  bookType,
		Pages: pages,
	}
}

func (book Book) ID() ID {
	return book.Id
}

type Books interface {
	Repository[Book]
	FindByPath(path Path) (Book, error)
}

type BookStorage interface {
	Storage
}

func UpsertBook(path Path) error {
	book, err := books.FindByPath(path)
	if errors.Is(err, ErrNotFound) {
		book, err = createBook(path)
		if err != nil {
			return err
		}
		return books.Save(book)
	}

	if err != nil {
		return err
	}

	book, err = updateBook(book)
	if err != nil {
		return err
	}

	return books.Save(book)
}

func createBook(path Path) (Book, error) {
	var err error
	var bookType BookType

	err = bookStorage.Retrieve(path, func(reader io.Reader) error {
		bookType, err = bookTypeIdentifier.Identify(reader)
		return err
	})
	if err != nil {
		return Book{}, err
	}

	switch bookType {
	case CBZ:
		return createComic(path, bookType)
	default:
		log.Info("Unsupported Book Type " + string(bookType))
		return Book{}, nil
	}
}

func createComic(path Path, bookType BookType) (Book, error) {
	var err error
	var pages []BookPage
	err = bookStorage.Retrieve(path, func(reader io.Reader) error {
		pages, err = comicBookContentStreamer.Files(reader)
		return err
	})
	if err != nil {
		return Book{}, err
	}

	return NewBook(path, bookType, pages), nil
}

func updateBook(book Book) (Book, error) {
	createdBook, err := createBook(book.Path)
	if err != nil {
		return Book{}, nil
	}

	createdBook.Id = book.Id
	return createdBook, nil
}

type InMemoryBooks struct {
	*InMemoryRepository[Book]
}

func NewInMemoryBooks() Books {
	repository := NewInMemoryRepository[Book]()
	return &InMemoryBooks{repository}
}

func (books *InMemoryBooks) FindByPath(path Path) (Book, error) {
	var foundBook Book
	for _, book := range books.store {
		if book.Path == path {
			foundBook = book
		}
	}

	if foundBook.ID() == "" {
		return foundBook, ErrNotFound
	}

	return foundBook, nil
}

type InMemoryBookStorage struct {
	*InMemoryStorage
}

func NewInMemoryBookStorage() BookStorage {
	return &InMemoryBookStorage{
		NewInMemoryStorage(),
	}
}

func (storage *InMemoryBookStorage) Retrieve(path Path, consumer func(io.Reader) error) error {
	data := storage.storage[path]
	reader := bytes.NewReader(data)
	return consumer(reader)
}
