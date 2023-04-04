package main

import (
    "strings"
)

func validateNumOfTickets() bool {
    isValidNumOfTickets := numOfTickets > 0 && numOfTickets <= remainingTickets
    
    return isValidNumOfTickets
}

func validateUserInfo(firstName string, lastName string, email string) (bool, bool) {
    isValidName := len(firstName) >= 2 && len(lastName) >= 2
    isValidEmail := strings.Contains(email, "@")

    return isValidName, isValidEmail
}
