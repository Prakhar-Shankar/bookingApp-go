package main

import (
	"strings"
)


//Function for user input validation
func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

//To export the validateUserInput to another package just capitalize the first letter of the function.