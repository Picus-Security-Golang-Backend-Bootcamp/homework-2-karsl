package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-karsl/library"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	switch args := os.Args; {
	case len(args) == 1:
		projectName := path.Base(args[0])
		fmt.Printf("Available commands for %s: \n search => search books\n list => list all books\n buy => buy a book\n delete => delete a book\n", projectName)

	case len(args) == 2 && args[1] == "list":
		booksInTheBookShelf := library.List()
		if len(booksInTheBookShelf) > 0 {
			fmt.Printf("Books in the bookshelf: %v.\n", booksInTheBookShelf)
		} else {
			fmt.Println("No books in the bookshelf!")
		}
	case len(args) == 2 && args[1] == "search":
		fmt.Println("Please enter name of the books you would like to search.")
	case len(args) > 2 && args[1] == "search":
		searchTerm := strings.Join(args[2:], " ")
		foundBooks := library.Search(searchTerm)
		if len(foundBooks) > 0 {
			fmt.Printf("Found books in the bookshelf: %v.\n", foundBooks)
		} else {
			fmt.Println("No books found!")
		}

	case (len(args) == 2 || len(args) == 3) && args[1] == "buy":
		fmt.Println("Please enter book id and quantity to be bought")
	case len(args) == 4 && args[1] == "buy":
		bookId, err := strconv.Atoi(args[2])
		if err != nil || bookId <= 0 {
			fmt.Println("Invalid book id")
			return
		}

		quantity, err := strconv.Atoi(args[3])
		if err != nil || quantity <= 0 {
			fmt.Println("Invalid quantity")
			return
		}

		err = library.Buy(uint(bookId), uint(quantity))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("You have successfully bought %d books with id %d.\n", quantity, bookId)

	case len(args) == 2 && args[1] == "delete":
		fmt.Println("Please enter a book id.")
	case len(args) == 3 && args[1] == "delete":
		bookId, err := strconv.Atoi(args[2])
		if err != nil || bookId <= 0 {
			fmt.Println("Invalid book id")
			return
		}

		err = library.DeleteBookById(uint(bookId))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Successfully deleted book with id %d\n", bookId)
	default:
		fmt.Println("Invalid arguments.")
	}
}
