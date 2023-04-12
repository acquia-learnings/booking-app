package ui

import (
    "fmt";
    "strings"
)

func PrintFirstNames(bookings []string) {
    var firstNames []string

    // Using blank identifier to ignore index variable
    // as it is not used anywhere
    for _, booking := range bookings {
        var name = strings.Fields(booking)
        firstNames = append(firstNames, name[0])
    }

    fmt.Printf("First Names of all the bookings %v\n", firstNames)
}

func PrintAllBookingInfo(allBookingInfo []map[string]string) {
    for index, bookingInfo := range allBookingInfo {
        fmt.Println(index, ": Booking is from", bookingInfo["firstName"], bookingInfo["lastName"], "with email:", bookingInfo["email"], "and has booked", bookingInfo["numOfTickets"], "tickets.")
    }
}