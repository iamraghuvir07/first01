/***
***Test Package to demo struct & interaface
***/

package main

import (
	"fmt"
	"time"
)

/***
***  Represents a book
***/
type Book struct {
	title  string
	author string
	price  float32
}

/***
***Struct represent a book store
***which will have a collection of books
***/
type Store struct {
	name  string
	place string
	books []Book
}

/***
*** interface to manage the invetory
***/
type inventory interface {
	addToCollection(o Book)
	removeFromCollection(o Book)
}

func (s *Store) addToCollection(o Book) {
	s.books = append(s.books, o)
}
func (s *Store) addBooksToCollection(o ...Book) { //variadic function
	s.books = append(s.books, o...)
}
func (s *Store) removeFromCollection(o Book) {
	for i, v := range s.books {
		fmt.Printf("\nTimeslot %d %d", i, v)
		if v == o {
			s.books = append(s.books[:i], s.books[i+1:]...)
			break
		}
	}
}

func (b *Book) printBook() {
	fmt.Printf("\nBook *** %v \n", b)
	b.price = 100
}

func main() {
	var book1 Book
	var store1 Store
	book1.author = "Author1"
	book1.title = "Book Tit1e1"
	book1.price = 156.5
	book1.printBook()

	store1 = Store{name: "Durga Book Store", place: "Banaglore", books: []Book{}}
	book2 := Book{title: "Book2", author: "Author2", price: 45.60}
	store1.addBooksToCollection(book1, book2)
	fmt.Printf("\nBook price changed? %v \n", book1)
	fmt.Printf("\nBook price changed? %v \n", store1)
	//store1.addToCollection(b)
	var inv inventory
	inv = &store1
	inv.removeFromCollection(book1)
	fmt.Printf("\nBook price changed? %v \n", store1)

	t := time.Now()
	fmt.Printf("\n Current Date Time %v \n", t)

	t1, _ := time.Parse(time.RFC822, "18 Feb 21 15:30 IST")
	fmt.Printf("\n Custom Date %v \n", t1)

}
