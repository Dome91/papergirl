package main

import "papergirl"

func main() {
	papergirl.Initialize(
		papergirl.NewInMemoryBooks(),
		nil,
		nil,
		nil,
		nil,
		NewBCryptPasswordHasher(),
	)

	papergirl.UpsertBook("")
}
