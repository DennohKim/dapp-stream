package main

import (
	"fmt"
	"strings"
)

// App entry point
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 100
	var remainingTickets uint = 50
	var bookings = []string{}

	fmt.Printf("conferenceName: %T, Remaining Tickets: %T ", conferenceName, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We currently have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//Ask user for name
		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address")
		fmt.Scan(&email)

		fmt.Println("How many tickets would you like to book?")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v, you have booked %v tickets for the %v conference. A confirmation email will be sent to %v\n", firstName, lastName, userTickets, conferenceName, email)
			fmt.Printf("There are now %v tickets  for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			// _ is a blank identifier
			for _, booking := range bookings {

				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0

			if remainingTickets == 0 {
				// End program
				fmt.Printf("Sorry, there are no more tickets available for %v\n", conferenceName)
				break
			}

		} else {
			fmt.Printf("Sorry, we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		}

	}

}
