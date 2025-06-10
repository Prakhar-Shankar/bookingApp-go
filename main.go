package main

import "fmt"

func main(){
	//A different syntax
	conferenceName := "Go Conference" // This equates to 'var conferenceName string = "Go Conference", := cannot be used for constants, it can only be used at variables also if user wants to specify a particular data type we dont use it.
	
	const totalTickets int = 50
	var remainingTickets uint = 50 

	// Format the output from here 
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Define username and print here 
	var userName string
	var userTicket int

	userName = "Tom"
	userTicket = 2

	fmt.Printf("User %v booked %v tickets\n", userName, userTicket)


}