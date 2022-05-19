package main

import "fmt"

func main(){
	var ConfName = "Go Conference"
	const ConfTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var tickets uint

	fmt.Printf("Welcome to %v Booking Application\n",ConfName)
	fmt.Println("Get your conference tickets here")
	fmt.Printf("We have a total of %v Tickets\n", ConfTickets)
	fmt.Printf("There are %v Tickets available\n", remainingTickets)
	
	fmt.Println("Please enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter email ID: ")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&tickets)
	remainingTickets = remainingTickets - tickets


	// ask the user for their name

	fmt.Printf("First Name: %v %v \nemail: %v\nTickets booked: %v \n",firstName,lastName,email,tickets)
	fmt.Printf("Remaining Tickets: %v\n",remainingTickets)




}