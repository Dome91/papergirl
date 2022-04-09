package papergirl

func Initialize(_books Books,
	_bookStorage BookStorage,
	_bookTypeIdentifier BookTypeIdentifier,
	_comicBookContentStreamer ComicBookContentStreamer,
	_users Users,
	_passwordHasher PasswordHasher,
) {
	books = _books
	bookStorage = _bookStorage
	bookTypeIdentifier = _bookTypeIdentifier
	comicBookContentStreamer = _comicBookContentStreamer
	users = _users
	passwordHasher = _passwordHasher
}
