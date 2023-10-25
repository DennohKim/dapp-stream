package main

import (
	"fmt"
)

//App entry point
func main(){
	var conferenceName = "Go Conference";
	const conferenceTickets = 100;
	var remainingTickets = 50

	fmt.Printf("conferenceName: %T, Remaining Tickets: %T ", conferenceName, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName);
	fmt.Printf("We currently have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets);
	fmt.Println("Get your tickets here to attend!");


	var userName string
	var userTickets int

	//Ask user for name
	userName = "Tom";
	userTickets = 3;
	fmt.Printf("User %v has booked %v tickets.\n", userName, userTickets);

}

