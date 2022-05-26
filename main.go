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

	greetUsers(ConfName,ConfTickets,remainingTickets)


	for remainingTickets>0{
		
		fmt.Println("Please enter your First Name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Please enter your Last Name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Please enter email ID: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter no of tickets: ")
		fmt.Scan(&tickets)

		isValidName := len(firstName) >= 2 && len(lastName)>=2
		isValidEmail := strings.Contains(email,"@")
		isValidTicket := tickets > 0 && tickets <= remainingTickets

		if isValidTicket && isValidName && isValidEmail{
			remainingTickets = remainingTickets - tickets

			displayUserInfo(firstName,lastName,email,int(tickets),remainingTickets)
		
			// Storing the user info in an array
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings,firstName + " " + lastName)
			fmt.Printf("The whole array is %v\n",bookings)
			
			FirstNames := []string{}
			
			for _, booking = range bookings{
				fmt.Printf("Values inside booking variable: %v\n", booking)
				var names = strings.Fields(booking)
				fmt.Printf("Names array: %v\n",names)
				// var name = names[0]
				FirstNames= append(FirstNames, names[0])
			}


	
			fmt.Printf("The First Names are %v\n",FirstNames)
			// fmt.Printf("The first value: %v\n",bookings[0])
	
			// var noTickets bool = remainingTickets == 0
			// noTickets := remainingTickets == 0
			// if noTickets{
			// 	fmt.Println("Tickets sold out")
			// 	break
			// }
	
			
		}else
		{
			if !isValidEmail {
				fmt.Println("Please enter the correct email ID")
			}
			if !isValidName{
				fmt.Println("Please enter a valid name")
			} 
			if !isValidTicket{

				fmt.Println("Ticket count exceeded the remaining tickets")
				fmt.Printf("Only %v tickets are remaining\n",remainingTickets)
			}
			
			continue
		}
	
	}
	
}

func greetUsers(confname string,ConfTickets uint, remainingTickets uint)  {
	fmt.Printf("Welcome to %v \n",confname)
	fmt.Printf("%v tickets are remaining out of %v\n",remainingTickets,ConfTickets)
} 	

func displayUserInfo(firstName string, lastName string, email string,tickets int, remTickets uint){

	fmt.Println("Tickets booked for")
	fmt.Printf("First Name: %v %v \nemail: %v\nTickets booked: %v \nTickets Remaining: %v\n",firstName,lastName,email,tickets,remTickets)
}

/*
func getUserInfo() []string{
	return 
}*/