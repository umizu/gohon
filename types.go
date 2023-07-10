package main

import "math/rand"

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func NewBook(title string) *Book {
	return &Book{
		ID:    rand.Intn(10000),
		Title: title,
	}
}
