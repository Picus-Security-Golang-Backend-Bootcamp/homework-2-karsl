package library

import (
	"errors"
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-karsl/helper"
	"strings"
	"time"
)

var idCount = 0
var books []Book

type Book struct {
	Id            int
	Name          string
	StockCode     string
	ISBN          int
	NumberOfPages int
	Price         float64
	Quantity      int
	Author        Author
	isDeleted     bool
}

type Author struct {
	Name      string
	BirthDate time.Time
}

type Deletable interface {
	Delete() error
}

func panicOnError(book Book, err error) Book {
	if err != nil {
		panic(err.Error())
	}

	return book
}

func init() {
	books = append(books,
		panicOnError(Construct("In Search of Lost Time", "SKU-012", Author{
			Name:      "Marcel Proust",
			BirthDate: time.Now(),
		})),
		panicOnError(Construct("Ulysses", "SKU-044", Author{
			Name:      "James Joyce",
			BirthDate: time.Date(1882, 2, 2, 0, 0, 0, 0, time.UTC),
		})),
		panicOnError(Construct("Don Quixote", "SKU-011", Author{
			Name:      "Miguel de Cervantes",
			BirthDate: time.Date(1547, 9, 29, 0, 0, 0, 0, time.UTC),
		})),
		panicOnError(Construct("One Hundred Years of Solitude", "SKU-103", Author{
			Name:      "Gabriel García Márquez",
			BirthDate: time.Date(1927, 3, 6, 0, 0, 0, 0, time.UTC),
		})),
		panicOnError(Construct("The Great Gatsby", "SKU-224", Author{
			Name:      "F. Scott Fitzgerald",
			BirthDate: time.Date(1896, 9, 24, 0, 0, 0, 0, time.UTC),
		})))
}

// Construct initializes a new Book whose ISBN, NumberOfPages, Price, Quantity fields generated randomly, isDeleted
// property set to false, Id increases incrementally.
func Construct(name string, stockCode string, author Author) (Book, error) {
	idCount += 1

	isbn, err := helper.GetRandomInt64(100000000000)
	if err != nil {
		return Book{}, err
	}
	numberOfPages, err := helper.GetRandomInt64(2000)
	if err != nil {
		return Book{}, err
	}
	price, err := helper.GetRandomFloat64(1000, 2)
	if err != nil {
		return Book{}, err
	}
	quantity, err := helper.GetRandomInt64(2000)
	if err != nil {
		return Book{}, err
	}

	return Book{
		Id:            idCount,
		Name:          name,
		StockCode:     stockCode,
		ISBN:          int(isbn),
		NumberOfPages: int(numberOfPages),
		Price:         price,
		Quantity:      int(quantity),
		Author:        author,
		isDeleted:     false,
	}, nil
}

func (author Author) String() string {
	return fmt.Sprintf("{Name: %s, BirthDate: %s}", author.Name, author.BirthDate.Format("02/01/2006"))
}

func (book Book) String() string {
	return fmt.Sprintf("{Id: %d, Name: %s, StockCode: %s, ISBN: %d, NumberOfPages: %d, Price: %f, Quantity: %d, Author: %s}",
		book.Id, book.Name, book.StockCode, book.ISBN, book.NumberOfPages, book.Price, book.Quantity, book.Author)
}

// matches checks Book if it matches with given term by considering Name, Author.Name, StockCode fields of Book.
func (book Book) matches(term string) bool {
	fieldsToCheck := []string{strings.ToLower(book.Name), strings.ToLower(book.Author.Name), strings.ToLower(book.StockCode)}
	for _, field := range fieldsToCheck {
		if strings.Contains(field, strings.ToLower(term)) {
			return true
		}
	}

	return false
}

// Search returns all books that match with term in the bookshelf.
func Search(term string) []Book {
	foundBooks := make([]Book, 0, len(books))

	for _, book := range books {
		if book.matches(term) && !book.isDeleted {
			foundBooks = append(foundBooks, book)
		}
	}

	return foundBooks
}

// List returns all books in the bookshelf.
func List() []Book {
	booksToList := make([]Book, 0, len(books))

	for _, book := range books {
		if !book.isDeleted {
			booksToList = append(booksToList, book)
		}
	}

	return booksToList
}

// buy decreases stock count. A Book can't be bought if there is not enough stock or deleted already.
func (book *Book) buy(quantityToBuy int) error {
	if book.Quantity < quantityToBuy {
		return errors.New("there is not enough items in the stock")
	} else if book.isDeleted {
		return errors.New("book was already deleted")
	}

	book.Quantity -= quantityToBuy
	return nil
}

// Buy find the book with given id and buys it.
func Buy(bookId, quantity int) error {
	foundBook, err := findBookById(bookId)
	if err != nil {
		return err
	}

	err = foundBook.buy(quantity)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the given book and removes it from books
func (book *Book) Delete() error {
	if book.isDeleted {
		return errors.New("book was already deleted")
	}

	book.isDeleted = true
	removeBookFromBookshelf(book.Id)

	return nil
}

// removeBookFromBookshelf removes the book with given id from the books slice and has no effect if no Book with given
// id is not found.
func removeBookFromBookshelf(id int) {
	for i, book := range books {
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
}

// findBookById returns the first book with given id in books.
func findBookById(id int) (Book, error) {
	for _, book := range books {
		if book.Id == id && !book.isDeleted {
			return book, nil
		}
	}

	return Book{}, errors.New("no such book found")
}

// DeleteBookById deletes the first book with given id in books.
func DeleteBookById(bookId int) error {
	book, err := findBookById(bookId)
	if err != nil {
		return err
	}

	err = book.Delete()
	if err != nil {
		return err
	}

	return nil
}
