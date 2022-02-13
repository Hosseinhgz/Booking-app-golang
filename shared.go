package main

import (
	"strings"
)

func validationInputs(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets

	// Go functions can return multiple values
	return isValidName, isValidEmail, isValidTicketNum

}
