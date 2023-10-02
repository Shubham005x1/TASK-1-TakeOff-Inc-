package validations

import (
	"strings"
	"unicode"
)

func IsValidEmail(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return nil
	}
	return CustomError{message: "Please Enter Valid Email"}
}

func IsValidSalary(salary float64) error {

	if salary >= 0 {
		return nil
	}
	return CustomError{message: "Salary Cannot be Negative Please Enter Valid Salary!"}

}
func IsValidRole(role string) error {
	allowedRoles := []string{"admin", "manager", "developer", "tester"}
	lowercaseRole := strings.ToLower(role)
	for _, allowedRole := range allowedRoles {
		if lowercaseRole == allowedRole {
			return nil
		}
	}

	return CustomError{message: "Invalid Role Please Enter Valid Role"}
}
func IsValidEntry(name string) error {
	if name == "" || validEntry(name) {
		return CustomError{message: "Field cannot be empty or cannot have Number in it please provide the valid value "}
	}
	return nil
}

func validEntry(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return true
		}
	}
	return false

}
