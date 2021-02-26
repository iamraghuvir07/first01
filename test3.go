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

type MedicalStore struct {
	medicines []string
}

/***
*** interface to manage the invetory
***/
type inventory interface {
	addToCollection(o interface{})
	removeFromCollection(o interface{})
}

func (s *MedicalStore) addToCollection(o interface{}) {
	s.medicines = append(s.medicines, o.(string))
}
func (s *MedicalStore) removeFromCollection(o interface{}) {
	b := o.(string)
	for i, v := range s.medicines {
		fmt.Printf("\nTimeslot %d %d", i, v)
		if v == b {
			s.medicines = append(s.medicines[:i], s.medicines[i+1:]...)
			break
		}
	}
}
func (s *Store) addToCollection(o interface{}) {
	s.books = append(s.books, o.(Book))
}
func (s *Store) addBooksToCollection(o ...Book) { //variadic function
	s.books = append(s.books, o...)
}
func (s *Store) removeBooksFromCollection(o ...Book) {

	fmt.Printf("\n Removing from the list %v", s.books)
	for _, b := range o {
		for i, v := range s.books {
			if v == b {
				fmt.Printf("\nBooks to remove", i, b)
				s.books = append(s.books[:i], s.books[i+1:]...)
			}
		}
	}

}
func (s *Store) removeFromCollection(o interface{}) {
	b := o.(Book)
	for i, v := range s.books {
		fmt.Printf("\nTimeslot %d %d", i, v)
		if v == b {
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
	book3 := Book{title: "Book3", author: "Author3", price: 245.60}
	book4 := Book{title: "Book4", author: "Author4", price: 245.60}
	store1.addBooksToCollection(book1, book2)
	fmt.Printf("\nBook price changed? %v \n", book1)
	fmt.Printf("\nBook price changed? %v \n", store1)
	//store1.addToCollection(b)
	var medStore MedicalStore
	var inv inventory
	inv = &store1
	inv.addToCollection(book3)
	inv.removeFromCollection(book1)
	inv.removeFromCollection(book4)
	inv = &medStore
	inv.addToCollection("Paracetamal")
	fmt.Printf("\nMedicines %v \n", medStore.medicines)
	inv.removeFromCollection("Paracetamal")
	fmt.Printf("\nMedicines %v \n", medStore.medicines)
	store1.removeBooksFromCollection(book2, book3)
	fmt.Printf("\nBook price changed? %v \n", store1)

	t := time.Now()
	fmt.Printf("\n Current Date Time %v \n", t)

	t1, _ := time.Parse(time.RFC822, "18 Feb 21 15:30 IST")
	fmt.Printf("\n Custom Date %v \n", t1)

}
