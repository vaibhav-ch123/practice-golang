package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Price    float64
	Quantity int
}

func add(books []Book) []Book {

	var id, quantity int
	var title, author string
	var price float64

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You want to add book now enter details: ")
	fmt.Print("Id ")
	fmt.Scanln(&id)
	fmt.Print("Title ")
	title, _ = reader.ReadString('\n')
	fmt.Print("Author ")
	author, _ = reader.ReadString('\n')
	fmt.Print("Price ")
	fmt.Scanln(&price)
	fmt.Print("Quantity ")
	fmt.Scanln(&quantity)

	books = append(books, Book{id, title, author, price, quantity})
	return books
}

func viewbooks(books []Book) {

	for _, book := range books {
		fmt.Println(book)
	}
}

func searchbook(books []Book) {

	var id int
	var title string

	fmt.Println("Enter the id and title for book which you want to search")
	fmt.Scanln(&id, &title)

	var sbook Book

	for _, book := range books {
		if book.ID == id && book.Title == title {
			sbook = book
			break
		}
	}

	fmt.Println(sbook)
}

func main() {

	var books []Book
	var choice int

outer:
	for {

		fmt.Println("Press 1 for Add book")
		fmt.Println("Press 2 for View books")
		fmt.Println("Press 3 for Search book")
		fmt.Println("Press 4 for Update book")
		fmt.Println("Press 5 for Delete book")
		fmt.Println("Press 6 for Borrow book")
		fmt.Println("Press 7 for Return book")
		fmt.Println("Press 8 for Show Statictics")
		fmt.Println("Press 9 for exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			books = add(books)
		case 2:
			viewbooks(books)
		case 3:
			searchbook(books)
		case 9:
			break outer
		}

	}
}
