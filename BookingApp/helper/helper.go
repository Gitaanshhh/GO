// each pkg usually in its own folder
package helper

// import "strings"

// Export -> makes it avail for all packages in the app => done by capitalizing the first letter
func ValidateUserInputs(userName string, userTickets uint, remainingTickets uint) (bool, bool) {
	isValidName := len(userName) >= 2
	// isValidEmail := strings.Contains(email, '@')
	isValidTicket := userTickets > 0 && userTickets < remainingTickets
	return isValidName, isValidTicket

}