package papergirl

func Initialize(_books Books, _bookStorage BookStorage, _bookTypeIdentifier BookTypeIdentifier, _comicBookContentStreamer ComicBookContentStreamer) {
	books = _books
	bookStorage = _bookStorage
	bookTypeIdentifier = _bookTypeIdentifier
	comicBookContentStreamer = _comicBookContentStreamer
}
