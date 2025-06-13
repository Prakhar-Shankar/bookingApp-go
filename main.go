package main

import (
	"fmt"
	"time"
)

//These are package level variable which are available for all the functions

//A different syntax
var conferenceName = "Go Conference" // This equates to 'var conferenceName string = "Go Conference", := cannot be used for constants, it can only be used at variables also if user wants to specify a particular data type we dont use it.

const totalTickets int = 50
var remainingTickets uint = 50 

// Creating a slice to show all the bookings 
var bookings = make([]userData, 0) // This is slice and it is an abstraction of the array, it is dynamic hence no fixed size. For array we just need to specify the size, i.e., bookings := [50]string{}

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}


func main(){
	
	//Function call
	greetUser()

	for{
	// Define username and print here 
	firstName, lastName, email, userTicket := getUserInput()

	// For validation we are defining three vars which contains three different conditions, we are using these conditions in the for loop.

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket)


	if isValidName && isValidEmail && isValidTicketNumber{
	// Writing the booking app logic that is counting the remaining ticket 
	bookTicket(userTicket, firstName, lastName, email)

	// using send ticket function over here
	//go create a concurrent thread which wont wait for the generation of the ticket and it will keep on taking the user Input
	go sendTicket(userTicket, firstName, lastName, email)
	//Function call
	firstNames := getFirstNames()
	fmt.Printf("These are the bookings %v\n", firstNames)
	
	}else{
		// Here we want to notify the user where the input data is invalid hence we have written output for every condition, instead of a generic statement.

		if !isValidName{
			fmt.Println("First name or the last name entered by the user is invalid.")
		}
		if !isValidEmail{
			fmt.Println("The email entered by the user does not contain @")
		}
		if !isValidTicketNumber{
			fmt.Println("The number of ticket entered by the user is invalid")
		}

	}
}
}

//Lets create a function which does the job of greeting, we want to keep the main fucntion clean.
func greetUser(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", totalTickets, remainingTickets)
	fmt.Printf("Get the tickets to attend!!\n")

}

//Creating a function which prints the first names of the users.
func getFirstNames() []string{
	firstNames := []string{} 
	
	// here we have to give index and a var but as we are not using index we are giving _.
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

//A function to get the user inputs
func getUserInput() (string, string, string, uint){
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

	return firstName, lastName, email, userTicket
}

//This is the logic of booking a ticket in the system
func bookTicket(userTicket uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTicket


	// create a map which contains all the user details

	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTicket,
	}

	
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("The user %v %v, booked %v tickets. A confirmation mail will be sent at %v. Thank you !!\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket (userTicket uint, firstName string, lastName string, email string){
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending Ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("########################")

}