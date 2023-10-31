package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets = 100

var remainingTickets uint = 50
var bookings = make([]map[string]string,0)

// App entry point
func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// End program
				fmt.Printf("Sorry, there are no more tickets available for %v\n", conferenceName)
				break
			}

		} else {
			//fmt.Printf("Sorry, we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to our booking application for %v \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

}

func getFirstNames() []string {
	firstNames := []string{}

	// _ is a blank identifier
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = make(map[string]string);

	userData["firstName"] = firstName;
	userData["lastName"] = lastName;
	userData["email"] = email;
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10);


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings);

	fmt.Printf("Thank you %v %v, you have booked %v tickets for the %v conference. A confirmation email will be sent to %v\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("There are now %v tickets  for %v\n", remainingTickets, conferenceName)
}
