package helper

import (
	"strings"
)

// check user validation
func UserInputValidation(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidTicket
}
