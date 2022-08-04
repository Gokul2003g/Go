package main

import "strings"

func validateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 1
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}
