/*
ref - https://youtu.be/yyUHQIec83I?si=gHcKorrhDgIE8FZD

Must run with : go run main.go helper.go (or) go run . //all files in curr dir
and not : go run main.go
as now there are > 1 files in the package with dependencies
*/

// Each executeable must have a main package
package main

import (
	"bookingApp/helper"
	"fmt"
	"strconv"
	"strings"
	"time"
	"sync"
)

// Package level vars => Withing a package : need syntax w var keyword
// Can be passed to a var but then becomes a shadow of the global var => a local copy
var remainingTickets uint = 50

/*
type -> similar to typedef, means we are creating a new type
*/
type UserData struct {
	name                   string
	tickets                uint
	isOptedInForNewsLetter bool
}

var bookingsNew = make([]UserData, 0)

var wg = sync.WaitGroup {}

// Global vars => Capitalized first letter => similar to export

// Entry point -> 1 per app
func main() {
	conferenceName := "GO Conference"
	const conferenceTickets = 50

	// %T -> the var type
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	var bookings = make([]map[string]string, 0) //0 -> inital size, will inc dynamically
	// var bookings = []string{"Gitaansh"}
	// Or init as an empty arr -> var bookings [50] string

	// for { } = for true { }
	for remainingTickets > 0 && len(bookings) < 50 {
		userName, userTickets := getUserInput()
		isValidName, isValidTicket := helper.ValidateUserInputs(userName, userTickets, remainingTickets)

		city := "London"

		switch city {
		case "New York", "Hong Kong":
			//some code
		case "London":
			fmt.Println("London!")
		default:
			fmt.Println("No such conference.")
		}

		if !isValidName || !isValidTicket {
			fmt.Println("Invalid Input!")
			continue
		}

		bookings = bookTicket(userTickets, bookings, userName, conferenceName)
		
		wg.Add(1)	//2	: it sets the num of go routines to wait for by inc a counter
		go sendTicket(userTickets, userName)		//magically makes it run on a new t lmao
		// go dosmt()

		fmt.Printf("Bookings : \n")
		firstNames := getFirstNames(bookings)

		for i := 0; i < len(firstNames); i++ {
			fmt.Printf("%v\n", firstNames[i])
		}
		fmt.Println()
		// OR fmt.Printf("%v\n", firstNames)

		// status := remainingTickets == 0
		if remainingTickets == 0 {
			fmt.Println("We're full!")
			break
		}
	}
	wg.Wait()
}

// Global passed as a param -> local copy in the fn
func greetUser(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Println("Welcome to", confName, "!") // Automatrically added the whitespace before and after the var
	fmt.Printf("Out of %v we have %d tickets availble!", conferenceTickets, remainingTickets)
	fmt.Println("Get tickets here.")
}

func getFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // _, bookingNew := range bookingsNew
		name := strings.Fields(booking["Name"])  // Similar to split() in py
		firstNames = append(firstNames, name[0]) // bookingNew.name
	}
	// Safe - underlying array remains alive because the returned slice still references
	// If a local variable needs to outlive the function, Go automatically allocates it on the heap.
	return firstNames
}

func getUserInput() (string, uint) {
	//Initializing var name
	var userName string
	var userTickets uint //Only positive

	fmt.Println("Enter Name : ")
	fmt.Scan(&userName)

	fmt.Println("Enter Tickets : ")
	fmt.Scan(&userTickets)
	return userName, userTickets
}

// Because Go passes arguments by value.
func bookTicket(userTickets uint, bookings []map[string]string, userName string, conferenceName string) []map[string]string {

	remainingTickets -= userTickets

	var userData = make(map[string]string)
	userData["Name"] = userName
	userData["Tickets"] = strconv.FormatUint(uint64(userTickets), 10) //Base 10

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings Map : %v Tickets!\n", bookings)

	var userDataNew = UserData{
		name:                   userName,
		tickets:                userTickets,
		isOptedInForNewsLetter: true,
	}
	bookingsNew = append(bookingsNew, userDataNew)
	fmt.Printf("List of bookings Struct : %v Tickets!\n", bookingsNew)

	fmt.Printf("Thank you %s for buying %d Tickets!\n", userName, userTickets)
	fmt.Printf("%d Tickets remaining for %v.\n\n", remainingTickets, conferenceName)
	return bookings
}

func sendTicket(userTickets uint, userName string) {
	time.Sleep(10 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets for %v.", userTickets, userName)
	fmt.Println("==============================")
	fmt.Printf("Sending Ticket:\n %v \nvia email.\n", Ticket)
	fmt.Println("==============================")
	wg.Done() //remove t from waiting list
}