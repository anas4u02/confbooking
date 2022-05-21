package main

import (
	"fmt"
	"strings"
)

func main(){
	var ConfName = "Go Conference"
	const ConfTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var tickets uint
	var booking string

	var bookings [] string
	// Another way of declaring the array using shorthand operator or technique
	// bookings:= []string{}

	fmt.Printf("Welcome to %v Booking Application\n",ConfName)
	fmt.Println("Get your conference tickets here")
	fmt.Printf("We have a total of %v Tickets\n", ConfTickets)
	fmt.Printf("There are %v Tickets available\n", remainingTickets)

	for {
		
		fmt.Println("Please enter your First Name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Please enter your Last Name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Please enter email ID: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter no of tickets: ")
		fmt.Scan(&tickets)
		if tickets > remainingTickets{
			fmt.Println("Ticket count exceeded the remaining tickets")
			fmt.Printf("Only %v tickets are remaining\n",remainingTickets)
			continue
			
		}
		remainingTickets = remainingTickets - tickets
	
		fmt.Println("Tickets booked for")
		fmt.Printf("First Name: %v %v \nemail: %v\nTickets booked: %v \n",firstName,lastName,email,tickets)
		fmt.Printf("Remaining Tickets: %v\n",remainingTickets)
	
		// Storing the user info in an array
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings,firstName + " " + lastName)
		fmt.Printf("The whole array is %v\n",bookings)
		
		FirstNames := []string{}
		
		for _, booking = range bookings{
			var names = strings.Fields(booking)
			// var name = names[0]
			FirstNames= append(FirstNames, names[0])
		}

		fmt.Printf("The First Names are %v\n",FirstNames)
		// fmt.Printf("The first value: %v\n",bookings[0])

		// var noTickets bool = remainingTickets == 0
		noTickets := remainingTickets == 0
		
		if noTickets{
			fmt.Println("Tickets sold out")
			break
		}
	
	
	
	}
	
	
}