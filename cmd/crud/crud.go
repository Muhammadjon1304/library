package crud

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var title string

var author string

type Library struct {
	Shelf *sql.DB
}

type Book struct {
	Title  string
	Author string
}

func (l *Library) Create(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter book name:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Println("Enter author name:")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := Book{title, author}

	InsertIntoPostgres(db, book.Title, book.Author)
	fmt.Println("Book added")
}

//	func (l Library) Read() {
//		for index, book := range l.Shelf {
//			fmt.Printf("%d.%s-%s\n", index+1, book.Title, book.Author)
//		}
//	}
func (l *Library) Update(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which book do you want to change?")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Println("Enter new book name:")
	newTitle, _ := reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)
	fmt.Println("Enter new author name:")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)
	Update(db, newTitle, author, title)
}

func (l *Library) Delete(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which book do you want to delete?")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	Delete(db, title)
}
func ConnectPostgresDB() *sql.DB {
	connstring := "user=muhammadjonparpiyev dbname=library password='root' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func InsertIntoPostgres(db *sql.DB, title, author string) {
	_, err := db.Exec("INSERT INTO  library(title,author) VALUES($1,$2)", title, author)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value inserted")
	}
}
func (shelf Library) Read(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM library")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Title:   Author:")
		for rows.Next() {
			rows.Scan(&title, &author)
			fmt.Printf("%s - %s \n", title, author)
		}

	}
}
func Update(db *sql.DB, newTitle, newAuthor, title string) {
	_, err := db.Exec("UPDATE library SET title=$1,author=$2 WHERE title=$3", newTitle, newAuthor, title)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book updated")
	}
}
func Delete(db *sql.DB, title string) {
	_, err := db.Exec("DELETE FROM library WHERE title=$1", title)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book deleted")
	}
}
