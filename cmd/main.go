package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	crud "github.com/muhammadjon1304/library/crud"
	"os"
)

var db *sql.DB

func main() {
	db := crud.ConnectPostgresDB()
	newLib := crud.Library{
		Shelf: db,
	}
	for true {
		fmt.Println("Menu:\n0.Exit\n1.Add book\n2.Show books\n3.Change books\n4.Delete book")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			os.Exit(0)
		case 1:
			newLib.Create(db)
			break
		case 2:
			newLib.Read(db)
			break
		case 3:
			newLib.Update(db)
			break
		case 4:
			newLib.Delete(db)
			break
		}
	}
}
