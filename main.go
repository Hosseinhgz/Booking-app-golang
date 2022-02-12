package main

import (
	"fmt" // format package from Go
	"strings"
)

func main() {
	// Variables
	conferenceName := "Go conferance"
	const conferanceTickets = 50   //cant change the value
	var remainingTickets uint = 50 // uint used for positive int

	fmt.Printf("conferanceName var is %T, conferanceTickets var is %T, remainingTickets var is %T \n", conferanceTickets, conferenceName, remainingTickets)
	// Main
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available!\n", conferanceTickets, remainingTickets)
	fmt.Println("GET YOUR TICKET FROM HERE")

	// define array 3 methods
	// var bookings []string // [50]string is type of this array
	// var bookings = [50]string{}
	// var bookings = [50]string{"Hossein", "Jack"}

	//define slice
	var bookings []string

	// infinite loop
	for {

		// User Input Variables
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// IMP - print the the memory location of that variable
		//fmt.Println(&remainingTickets)

		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName+",")

		fmt.Printf("the whole slice: %v \n", bookings)
		fmt.Printf("slice first element: %v \n", bookings[0])
		fmt.Printf("the Slice type: %T \n", bookings)
		fmt.Printf("length of Slice: %v \n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v are still available\n", conferenceName, remainingTickets)

		firstNamesSlice := []string{}
		for _, booking := range bookings {
			var fullName = strings.Fields(booking)
			firstNamesSlice = append(firstNamesSlice, fullName[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNamesSlice)

		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out!!")
			break
		}
	}
}
