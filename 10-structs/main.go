package main

import (
	"fmt"
	"strconv"
	"time"
)

// define structure
type Book struct {
	rackNumber int
	bookId string
	issueDate string
}

// method (value receiver)
func (b Book) rackAndBookId() string{
	b.rackNumber++
	fmt.Println("Address at func(value receiver) of Book Id(",b.bookId,"): ",&b.rackNumber)
	return "Hi, the book with id "+ b.bookId+" is picked up from rack number "+strconv.Itoa(b.rackNumber)

}

// method (pointer receiver)
func (b *Book) increaseRackCount(){
	fmt.Println("Address at func(pointer receiver) of Book Id(",b.bookId,"): ",&b.rackNumber)
	b.rackNumber+=10
}

func main(){
	b1:=Book{
		rackNumber: 32,
		bookId:     "B1000",
		issueDate:  time.Now().Format("2006-01-02"),
	}
	fmt.Printf("Book Id: %s, Book Issue Date: %s\n",b1.bookId,b1.issueDate)
	b2:=Book{
		rackNumber: 11,
		bookId:     "BK212",
		issueDate:  time.Now().Format("2006-01-02"),
	}
	fmt.Printf("Book Id: %s, Book Issue Date: %s\n",b2.bookId,b2.issueDate)

	//method call
	fmt.Println("------Real Address (at main method) of Book Id(",b1.bookId,") : ",&b1.rackNumber,"---------")
	fmt.Println(b1.rackAndBookId())
	b1.increaseRackCount()
	fmt.Println(b1.rackAndBookId())
}