// ref - https://youtu.be/yyUHQIec83I?si=gHcKorrhDgIE8FZD

// Each executeable must have a main package
package main

import (
	"fmt"
	"strings"
)

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

	// for { } = for true { }
	for remainingTickets > 0 && len(bookings) < 50{
		//Initializing var name
		var userName string
		var userTickets uint //Only positive

		fmt.Println("Enter Name : ")
		fmt.Scan(&userName)

		fmt.Println("Enter Tickets : ")
		fmt.Scan(&userTickets)

		isValidName := len(userName) >= 2
		// isValidEmail := strings.Contains(email, '@')
		isValidTicket := userTickets > 0 && userTickets < uint(remainingTickets)

		if !isValidName || !isValidTicket {
			fmt.Println("Invalid Input!")
			continue
		}

		remainingTickets -= int(userTickets)

		bookings = append(bookings, userName)

		fmt.Printf("Thank you %s for buying %d Tickets!\n", userName, userTickets)
		fmt.Printf("%d Tickets remaining for %v.\n", remainingTickets, conferenceName)
		fmt.Printf("Bookings : \n")

		firstNames := []string{}
		for _, booking := range bookings {
			name := strings.Fields(booking)	//Similar to split() in py
			firstNames = append(firstNames, name[0])
		}

		for i := 0; i < len(bookings); i++ {
			fmt.Printf("%v\n", bookings[i])
		}

		// status := remainingTickets == 0 
		if remainingTickets == 0 {
			fmt.Println("We're full!")
			break
		}
	}
}
