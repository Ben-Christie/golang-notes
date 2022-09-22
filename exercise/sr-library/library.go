//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string // book title
type Name string  // memebrs name

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total number of books
	lended int // total books lended by the library
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string

		// if the book hasn't been checked in yet
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printAllMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()

	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}

	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more copies available to lend from the library")
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("Member did not checkout this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	// create books
	library.books["Harry Potter"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["A Song of Ice and Fire"] = BookEntry{
		total:  3,
		lended: 0,
	}

	library.books["The Last Wish"] = BookEntry{
		total:  5,
		lended: 0,
	}

	library.books["How to Survive the End of the World"] = BookEntry{
		total:  2,
		lended: 0,
	}

	//create members
	library.members["Ben"] = Member{
		"Ben",
		make(map[Title]LendAudit),
	}

	library.members["Alex"] = Member{
		"Alex",
		make(map[Title]LendAudit),
	}

	library.members["Rory"] = Member{
		"Rory",
		make(map[Title]LendAudit),
	}

	library.members["Cillian"] = Member{
		"Cillian",
		make(map[Title]LendAudit),
	}

	// print initial status of the library
	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printAllMemberAudits(&library)

	// check book out
	member := library.members["Ben"]
	checkedOut := checkoutBook(&library, "Harry Potter", &member)
	fmt.Println("\nCheck out a book!")

	if checkedOut {
		printLibraryBooks(&library)
		printAllMemberAudits(&library)
	}

	// return book
	returned := returnBook(&library, "Harry Potter", &member)
	fmt.Println("\nReturn a book!")

	if returned {
		printLibraryBooks(&library)
		printAllMemberAudits(&library)
	}
}
