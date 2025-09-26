package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePhone checks if phone number format is valid (basic validation)
func ValidatePhone(phone string) bool {
	if phone == "" {
		return true // Phone is optional
	}
	phoneRegex := regexp.MustCompile(`^\+?[\d\s\-\(\)]{10,15}$`)
	return phoneRegex.MatchString(phone)
}

// SanitizeString removes extra whitespace and trims string
func SanitizeString(input string) string {
	return strings.TrimSpace(input)
}

// ValidateStatus checks if status is one of the allowed values
func ValidateStatus(status string) bool {
	allowedStatuses := []string{"pending", "reviewed", "interviewed", "hired", "rejected"}
	for _, allowed := range allowedStatuses {
		if status == allowed {
			return true
		}
	}
	return false
}
