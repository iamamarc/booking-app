package helper

import "strings"

func ValidateUserInput(lastName string, firstName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(lastName) >= 2 && len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
