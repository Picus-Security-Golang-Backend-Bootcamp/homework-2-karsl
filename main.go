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
	if args := os.Args; len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("Available commands for %s: \n search => search books\n list => list all books\n buy => buy a book\n delete => delete a book\n", projectName)
	} else {
		switch command, commandArgs := args[1], args[2:]; {
		case command == "list":
			list()
		case command == "search":
			search(commandArgs)
		case command == "buy":
			buy(commandArgs)
		case command == "delete":
			deleteBook(commandArgs)
		default:
			fmt.Println("Invalid command.")
		}
	}
}

func list() {
	booksInTheBookShelf := library.List()
	if len(booksInTheBookShelf) > 0 {
		fmt.Printf("Books in the bookshelf: %v.\n", booksInTheBookShelf)
	} else {
		fmt.Println("No books in the bookshelf!")
	}
}

func search(args []string) {
	if len(args) < 1 {
		fmt.Println("Please enter name of the books you would like to search.")
		return
	}

	searchTerm := strings.Join(args, " ")
	foundBooks := library.Search(searchTerm)
	if len(foundBooks) > 0 {
		fmt.Printf("Found books in the bookshelf: %v.\n", foundBooks)
	} else {
		fmt.Println("No books found!")
	}
}

func buy(args []string) {
	if len(args) < 2 {
		fmt.Println("Please enter book id and quantity to be bought")
		return
	}

	bookId, err := strconv.Atoi(args[0])
	if err != nil || bookId <= 0 {
		fmt.Println("Invalid book id")
		return
	}

	quantity, err := strconv.Atoi(args[1])
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
}

func deleteBook(args []string) {
	if len(args) < 1 {
		fmt.Println("Please enter a book id.")
		return
	}

	bookId, err := strconv.Atoi(args[0])
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
}
