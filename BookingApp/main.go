// Each executeable must have a main package
package main

import "fmt"

// Entry point -> 1 per app
func main() {
	conferenceName := "GO Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	// %T -> the var type
	fmt.Println("Welcome to", conferenceName, "!") //Automatrically added the whitespace before and after the var
	fmt.Printf("Out of %v we have %d tickets availble!", conferenceTickets, remainingTickets)
	fmt.Println("Get tickets here.")

	var bookings = []string{"Gitaansh"}
	// Or init as an empty arr -> var bookings [50] string

	//Initializing var name
	var userName string
	var userTickets uint //Only positive

	fmt.Println("Enter Name : ")
	fmt.Scan(&userName)

	fmt.Println("Enter Tickets : ")
	fmt.Scan(&userTickets)
	remainingTickets -= int(userTickets)

	fmt.Printf("Thank you %s for buying %d Tickets!\n", userName, userTickets)
	fmt.Printf("%d Tickets remaining for %v.\n", remainingTickets, conferenceName)
	fmt.Printf("Bookings : \n")
	for i := 0; i < len(bookings); i++ {
		fmt.Printf("%v\n", bookings[i])
	}
}
