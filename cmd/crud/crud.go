package crud

import (
	"fmt"
)

type Library struct {
	Shelf []Book
}

type Book struct {
	Title  string
	Author string
}

func (l *Library) Create() {
	fmt.Println("Enter book name:")
	var title string
	fmt.Scanln(&title)
	fmt.Println("Enter author name:")
	var author string
	fmt.Scanln(&author)

	var book Book
	book.Title = title
	book.Author = author
	_ = append(l.Shelf, book)
	fmt.Println("Book added")
}
func (l *Library) Read() {
	for index, book := range l.Shelf {
		fmt.Printf("%d.%s-%s\n", index+1, book.Title, book.Author)
	}
}
func (l *Library) Update() {
	fmt.Println("Which book do you want to change?")
	var title string
	fmt.Scanln(&title)

	for _, book := range l.Shelf {
		if book.Title == title {
			fmt.Println("Enter new book name:")
			var title string
			fmt.Scanln(&title)
			fmt.Println("Enter new author name:")
			var author string
			fmt.Scanln(&author)
			book.Title = title
			book.Author = author
		}

	}
}
func (l *Library) Delete() {
	fmt.Println("Which book do you want to delete?")
	var title string
	fmt.Scanln(&title)

	for index, book := range l.Shelf {
		if book.Title == title {
			_ = append(l.Shelf[:index], l.Shelf[index+1:]...)
			break
		}

	}
}
