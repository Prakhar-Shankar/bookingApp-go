package main

import (
	"fmt"
	"strings"

)

func main(){
	//A different syntax
	conferenceName := "Go Conference" // This equates to 'var conferenceName string = "Go Conference", := cannot be used for constants, it can only be used at variables also if user wants to specify a particular data type we dont use it.

	const totalTickets int = 50
	var remainingTickets uint = 50 

		// Creating a slice to show all the bookings 
	bookings := []string {} // This is slice and it is an abstraction of the array, it is dynamic hence no fixed size. For array we just need to specify the size, i.e., bookings := [50]string{}

	// Format the output from here 
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for{
	// Define username and print here 
	var firstName string
	var lastName string
	var email string
	var userTicket uint


	//Now lets get user input 
	fmt.Printf("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTicket)

	// Write the logic if user wants to book more tickets than remaining tickets, we want to either break the loop or keep asking the user about the details.

	if userTicket > remainingTickets{
		fmt.Printf("User is trying to book %v tickets, but only %v tickets are remaining\n", userTicket, remainingTickets)
		continue
	}
	
	// Writing the booking app logic that is counting the remaining ticket 
	remainingTickets = remainingTickets - userTicket

	fmt.Printf("The user %v %v, booked %v tickets. A confirmation mail will be sent at %v. Thank you !!\n", firstName, lastName, userTicket, email)
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	firstNames := []string{}
	
	// here we have to give index and a var but as we are not using index we are giving _.
	for _, booking := range bookings{
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}

	fmt.Printf("These are all the bookings: %v\n", firstNames)

	if remainingTickets == 0 {
		fmt.Println("All the tickets are booked. Please try next year.")
		break
	}

}
	

}