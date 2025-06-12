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

	//Function call
	greetUser(conferenceName, totalTickets, remainingTickets)

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

	// For validation we are defining three vars which contains three different conditions, we are using these conditions in the for loop.

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)


	if isValidName && isValidEmail && isValidTicketNumber{
		// Writing the booking app logic that is counting the remaining ticket 
	remainingTickets = remainingTickets - userTicket

	fmt.Printf("The user %v %v, booked %v tickets. A confirmation mail will be sent at %v. Thank you !!\n", firstName, lastName, userTicket, email)
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	//Function call
	firstName := getFirstNames(bookings)
	fmt.Printf("These are the bookings %v\n", firstName)
	
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
func greetUser(confName string, confTickets int, remainingTickets uint){
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", confTickets, remainingTickets)
	fmt.Printf("Get the tickets to attend!!\n")

}

//Creating a function which prints the first names of the users.
func getFirstNames(bookings []string) []string{
	firstNames := []string{} 
	
	// here we have to give index and a var but as we are not using index we are giving _.
	for _, booking := range bookings{
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}

	return firstNames
}

//Function for user input validation
func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}