package main

import (
	"fmt" // format package from Go
	"strings"
)

// package level Variables (global variables)
var conferenceName = "Go conferance"

const conferanceTickets = 50  //cant change the value
var remainingTickets int = 50 // int used for positive int
//define slice
var bookings = []string{}

func main() {

	// this function is greet to the users
	greetUsers()

	// define array 3 methods
	// var bookings []string // [50]string is type of this array
	// var bookings = [50]string{}
	// var bookings = [50]string{"Hossein", "Jack"}

	// infinite loop
	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNum := validationInputs(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNum {

			// check details of slice
			// fmt.Printf("the whole slice: %v \n", bookings)
			// fmt.Printf("slice first element: %v \n", bookings[0])
			// fmt.Printf("the Slice type: %T \n", bookings)
			// fmt.Printf("length of Slice: %v \n", len(bookings))

			bookTicket(userTickets, firstName, lastName, email)

			var firstNames = getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out!!")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Printf("Your email does not contains @ sign\n")

			}
			if !isValidName {
				fmt.Printf("Your name and family should be at least 2 character\n")
			}
			if !isValidTicketNum {
				fmt.Printf("Your ticket number is bigger than remaining tickets\n")

			}
		}
	}
}

func greetUsers() {
	fmt.Printf("conferanceName var is %T, conferanceTickets var is %T, remainingTickets var is %T \n", conferanceTickets, conferenceName, remainingTickets)
	// Main
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available!\n", conferanceTickets, remainingTickets)
	fmt.Println("GET YOUR TICKET FROM HERE")
}

func getFirstNames() []string {
	firstNamesSlice := []string{}
	for _, booking := range bookings {
		var fullName = strings.Fields(booking)
		firstNamesSlice = append(firstNamesSlice, fullName[0])
	}
	return firstNamesSlice
}

func validationInputs(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets

	// Go functions can return multiple values
	return isValidName, isValidEmail, isValidTicketNum

}

func getUserInput() (string, string, string, int) {
	// User Input Variables
	var firstName string
	var lastName string
	var email string
	var userTickets int

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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName+",")

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are still available\n", remainingTickets)
}
