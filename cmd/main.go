package main

import "papergirl"

func main() {
	books := papergirl.NewInMemoryBooks()
	papergirl.Initialize(books)
	papergirl.UpsertBook("")
}
