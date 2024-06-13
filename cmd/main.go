package main

import (
	"fmt"
	"github.com/muhammadjon1304/library/crud"
	"os"
)

func main() {
	newLib := crud.Library{
		Shelf: []crud.Book{
			crud.Book{"1984", "George Orwell"},
			crud.Book{"Black swan", "Nassim Taleb"},
		},
	}
	for true {
		fmt.Println("Menu:\n0.Exit\n1.Add book\n2.Show books\n3.Change books\n4.Delete book")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			os.Exit(0)
		case 1:
			newLib.Create()
			break
		case 2:
			newLib.Read()
			break
		case 3:
			newLib.Update()
			break
		case 4:
			newLib.Delete()
			break
		}
	}
}
