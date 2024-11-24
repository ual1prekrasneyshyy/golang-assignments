package main

import (
	"regexp"
	"strings"
)

func VaidateId(Id string) bool {
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(Id)
}

func ValidateString(Value string) bool {
	toUpper := strings.ToUpper(Value)

	if strings.Contains(toUpper, "SELECT") {
		return false
	}

	if strings.Contains(toUpper, "INSERT") {
		return false
	}

	if strings.Contains(toUpper, "UPDATE") {
		return false
	}

	if strings.Contains(toUpper, "DELETE") {
		return false
	}

	if strings.Contains(toUpper, "UNION") {
		return false
	}

	if strings.Contains(toUpper, "DROP") {
		return false
	}

	if strings.Contains(toUpper, "FROM") {
		return false
	}

	return true
}
