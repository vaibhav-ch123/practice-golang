package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

type Book struct {
	ID       int
	Title    string
	Author   string
	Price    float64
	Quantity int
}

func add(books *[]Book) {

	var id, quantity int
	var title, author string
	var price float64

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

	*books = append(*books, Book{id, title, author, price, quantity})

}

func viewbooks(books []Book) {

	for _, book := range books {
		fmt.Println(book)
	}
}

func searchbook(books []Book) {

	var id int
	var title string

	fmt.Println("Enter detail for book which you want to search")
	fmt.Print("id: ")
	fmt.Scanln(&id)
	fmt.Print("title: ")
	title, _ = reader.ReadString('\n')

	var sbook Book
	var check bool

	for _, book := range books {
		if book.ID == id && book.Title == title {
			sbook = book
			check = true
			break
		}
	}

	if check {
		fmt.Println(sbook)
	} else {
		fmt.Println("Book not found")
	}
}

func updatebook(books []Book) {

	var id int
	fmt.Println("enter the id of book which you want to update: ")
	fmt.Scanln(&id)

	var title string
	var price float64
	var quantity int
	fmt.Print("enter updated title: ")
	title, _ = reader.ReadString('\n')
	fmt.Print("enter updated price: ")
	fmt.Scanln(&price)
	fmt.Print("enter updated quantity: ")
	fmt.Scanln(&quantity)

	for i := range books {
		if books[i].ID == id {
			books[i].Title = title
			books[i].Price = price
			books[i].Quantity = quantity
		}
	}

}

func deletebook(books *[]Book) {

	var id int
	fmt.Print("Enter book id which you want to delete: ")
	fmt.Scanln(&id)

	*books = append((*books)[0:id-1], (*books)[id:]...)
}

func borrowbook(books *[]Book, borrowcount *map[int]int) {

	var id int
	var check bool
	fmt.Print("enter the book id you want to borrow: ")
	fmt.Scanln(&id)

	for i := range *books {
		if (*books)[i].ID == id {
			if (*books)[i].Quantity > 0 {
				(*books)[i].Quantity -= 1
				(*borrowcount)[id]++
				check = true
			} else {
				break
			}
		}
	}

	if check {
		fmt.Println("book borrowed successfully")
	}
}

func returnbook(books *[]Book, borrowcount *map[int]int) {

	var id int
	var check bool
	fmt.Print("enter book id you want to return: ")
	fmt.Scanln(&id)

	for i := range *books {
		if (*books)[i].ID == id {
			if (*borrowcount)[id] > 0 {
				(*books)[i].Quantity += 1
				(*borrowcount)[id] -= 1
				check = true
			} else {
				break
			}
		}
	}

	if check {
		fmt.Println("book is return successfully")
	}
}

func showstatistic(books []Book) {

	fmt.Println("Statistics")

	var totalbooks, totalquantity int
	var totalprice, highpricebook, lowpricebook float64

	for _, book := range books {
		totalbooks += 1
		totalprice += book.Price
		totalquantity += book.Quantity
		if book.Price > highpricebook {
			highpricebook = book.Price
		}
		if book.Price < lowpricebook {
			lowpricebook = book.Price
		}
	}

	fmt.Printf("Total Books: %d\n", totalbooks)
	fmt.Printf("Total Quantity: %d\n", totalquantity)
	fmt.Printf("Most Expensive Book: %.2f\n", highpricebook)
	fmt.Printf("Cheapest Book: %.2f\n", lowpricebook)
	fmt.Printf("Average Price: %.2f\n", totalprice/float64(totalbooks))

}

func main() {

	var books []Book
	var choice int
	borrowcount := make(map[int]int)

outer:
	for {
		fmt.Println()
		fmt.Println("Press 1 for Add book")
		fmt.Println("Press 2 for View books")
		fmt.Println("Press 3 for Search book")
		fmt.Println("Press 4 for Update book")
		fmt.Println("Press 5 for Delete book")
		fmt.Println("Press 6 for Borrow book")
		fmt.Println("Press 7 for Return book")
		fmt.Println("Press 8 for Show Statistics")
		fmt.Println("Press 9 for exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			add(&books)
		case 2:
			viewbooks(books)
		case 3:
			searchbook(books)
		case 4:
			updatebook(books)
		case 5:
			deletebook(&books)
		case 6:
			borrowbook(&books, &borrowcount)
		case 7:
			returnbook(&books, &borrowcount)
		case 8:
			showstatistic(books)
		case 9:
			break outer
		}

	}
}
