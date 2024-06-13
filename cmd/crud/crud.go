package crud

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Library struct {
	Shelf []Book
}

type Book struct {
	Title  string
	Author string
}

func (l *Library) Create() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter book name:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Println("Enter author name:")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := Book{title, author}

	l.Shelf = append(l.Shelf, book)
	fmt.Println("Book added")
}
func (l Library) Read() {
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
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter new book name:")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Println("Enter new author name:")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			book := Book{title, author}
			l.Shelf = append(l.Shelf, book)
			return
		}

	}
}
func (l *Library) Delete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which book do you want to delete?")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	for index, book := range l.Shelf {
		if book.Title == title {
			l.Shelf = append(l.Shelf[:index], l.Shelf[index+1:]...)
			break
		}

	}
}
