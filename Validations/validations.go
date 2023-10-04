package validations

import (
	"strings"
	"unicode"
)

func IsValidEmail(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return nil
	}
	return CustomError{message: "=>Email should contain (@ and .) please enter valid email"}
}

func IsValidSalary(salary float64) error {

	if salary > 0 {
		return nil
	}
	return CustomError{message: "=>Salary cannot be negative please enter valid salary!"}

}
func IsValidRole(role string) error {
	allowedRoles := []string{"admin", "manager", "developer", "tester"}
	lowercaseRole := strings.ToLower(role)
	for _, allowedRole := range allowedRoles {
		if lowercaseRole == allowedRole {
			return nil
		}
	}

	return CustomError{message: "=>Invalid Role Please Enter Valid Role"}
}
func IsValidEntry(name string) error {
	if name == "" || validNameEntry(name) {
		return CustomError{message: "=>Field cannot be empty or cannot have Number in it please provide the valid value "}
	}
	return nil
}

func validNameEntry(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return true
		}
	}
	return false

}
func IsNumberValid(num string) error {
	if len(num) < 10 || validNumberEntry(num) {
		return CustomError{message: "=>Number cannot be less than 10 digit and cannot contain character Please try again"}
	}
	return nil
}

func validNumberEntry(name string) bool {
	for _, char := range name {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false

}
