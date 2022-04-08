package papergirl

var books Books

type Book struct {
	id      ID
	Title   string
	Path    string
	CoverID string
}

func (book Book) ID() ID {
	return book.id
}

type Books interface {
	Repository[Book]
}

type BookStorage interface {
	Storage
}

func UpsertBook(path string) {
	book := Book{Title: "myTitle"}
	books.Save(book)
	all, _ := books.FindAll()
	println(all[0].Title)
}

type InMemoryBooks struct {
	*InMemoryRepository[Book]
}

func NewInMemoryBooks() *InMemoryBooks {
	repository := NewInMemoryRepository[Book]()
	return &InMemoryBooks{repository}
}
