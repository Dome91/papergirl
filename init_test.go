package papergirl

func initializeTest() {
	Initialize(
		NewInMemoryBooks(),
		NewInMemoryBookStorage(),
		NewMockBookTypeIdentifier(CBZ, nil),
		NewMockComicBookContentStreamer([]BookPage{}),
	)
}
