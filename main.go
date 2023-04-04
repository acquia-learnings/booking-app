package main

// This means we will be creating an executable in bin directory
// This is not the only place where we can make changes.

import (
	// go looks for below package in built-in modules
	"fmt"
	"strconv"
	"sync"
	"time"

	// Below is a custom module `booking-app`
	// and our custom package `ui` is defined inside it
	"booking-app/ui"
)

// Below are package level variables and constants

// For the values that cannot change
// we can use const keyword
const conferenceTickets = 50
// syntactic sugar: for variables
// There are constraints: you cannot use it for constants and 
// you can also not define variable type
var conferenceName = "Go conference"
// For unit value can never be negative
var remainingTickets uint = 50
// Below is an example of Slice
// If we define the size for it, it will turn into an array
var bookings []string
// Though the list of maps below is initialized with size 0,
// This size increases dynamically as we keep adding values
var allBookingInfo = make([]map[string]string, 0)
var numOfTickets uint
var allUserInfo = make([]UserInfo, 0)
type UserInfo struct {
    firstName string
    lastName string
    email string
    optedForNewsletter bool
}
// WaitGroup ensures that we wait for
// the launched goroutine to finish
var wg = sync.WaitGroup{}


// Entry point for our application
// Entry point for execution
func main() {
    greetUsers()

    // Below is an implementation of an infinite for loop
    // Thus we need to use ctrl + C to exit the application
    for {
        setNumOfTickets()
        
        isValidNumOfTickets := validateNumOfTickets()

        if isValidNumOfTickets {
            
            firstName, lastName, email := getUserInfo()

            isValidName, isValidEmail := validateUserInfo(firstName, lastName, email)

            if isValidName && isValidEmail {
                
                bookTicket(firstName, lastName, email)
                
                ui.PrintFirstNames(bookings)

                if remainingTickets == 0 {
                    fmt.Println("We have run out of tickets!!!")
                    break
                }
            } else if !isValidName {
                fmt.Println("Please ensure that you have entered valid name.")
                continue
            } else {
                fmt.Println("Please ensure that you have entered valid e-mail.")
                continue
            }
        } else {
            fmt.Printf("Sorry we only have %v number of tickets available\n", remainingTickets)
            continue
        }
    }

    fmt.Println("Final list of confirmed bookings:")
    ui.PrintAllBookingInfo(allBookingInfo)

    fmt.Println("All user information:")
    printAllUserInfo(allUserInfo)

    sendTickets()
    // Blocks until the WaitGroup counter is zero
    wg.Wait()
}

func greetUsers() {
    fmt.Println("Welcome to conference!")
    // printf: takes template string that needs to be formatted
    // Also needs annotation verbs that tells fmt function
    // how to format the variable passed in
    fmt.Printf("The variable conferenceName if of type %T\n", conferenceName)
    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    // Note: & is used to reference the pointer
    // Thus here it prints out the memory location pointing out to variable conferenceName
    fmt.Println(&conferenceName)
    // Point to note here:
    // a trailing space is automatically added after string
    // thus in below output there are two spaces before (conferenceName)
    fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
    fmt.Println("Get your tickets here")
}

func setNumOfTickets() {
    fmt.Printf("Please enter number of tickets: ")
    fmt.Scan(&numOfTickets)
}

func getUserInfo() (string, string, string) {
    var firstName string
    var lastName string
    var email string
    
    fmt.Printf("Please enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Printf("Please enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Printf("Please enter your e-mail: ")
    fmt.Scan(&email)
    return firstName, lastName, email
}

func bookTicket(firstName string, lastName string, email string) {
    bookings = append(bookings, firstName + " " + lastName)

    remainingTickets = remainingTickets - numOfTickets

    setAllBookingInfo(firstName, lastName, email)
    setAllUserInfo(firstName, lastName, email)

    fmt.Println("Thank you", firstName, "for booking", numOfTickets, "tickets.")
    fmt.Printf("%v are now remaining.\n", remainingTickets)
    fmt.Printf("We have received bookings from %v\n", bookings)
    fmt.Printf("The total bookings: %v\n", len(bookings))
}

func setAllBookingInfo(firstName string, lastName string, email string) {
    // Maps in `go` can be of one data type only
    var bookingInfo = make(map[string]string)
    
    bookingInfo["firstName"] = firstName
    bookingInfo["lastName"] = lastName
    bookingInfo["email"] = email
    bookingInfo["numOfTickets"] = strconv.FormatUint(uint64(numOfTickets), 10)

    allBookingInfo = append(allBookingInfo, bookingInfo)
}

func setAllUserInfo(firstName string, lastName string, email string) {
    var userInfo = UserInfo {
        firstName: firstName,
        lastName: lastName,
        email: email,
        optedForNewsletter: true,
    }

    allUserInfo = append(allUserInfo, userInfo)
}

func printAllUserInfo(allUserInfo []UserInfo) {
    for index, userInfo := range allUserInfo {
        fmt.Println(index, ": Booking is from", userInfo.firstName, userInfo.lastName, "with email:", userInfo.email, "and has opted for Newsletter:", userInfo.optedForNewsletter)
    }
}

func sendTickets() {
    for _, bookingInfo := range allBookingInfo {
        // Saves formatted string into a variable
        ticketInfo := fmt.Sprintf("Sending %v tickets: \nTo user %v %v\nOn mail: %v\n", bookingInfo["numOfTickets"], bookingInfo["firstName"], bookingInfo["lastName"], bookingInfo["email"])
        // Sets the number of goroutines to wait for
        // increases the counter by the provided number
        wg.Add(1)
        // Having go keyword in front of any method
        // ensures that `go` starts a new goroutine
        go sendEmail(ticketInfo)
    }
}

func sendEmail(ticketInfo string) {
    // Sleep function stops current thread (goroutine) execution for defined duration
    time.Sleep(15 * time.Second)
    fmt.Println("###############")
    fmt.Printf("%v", ticketInfo)
    fmt.Println("###############")
    // Decrements the WaitGroup counter by 1
    // Thus it's called by goroutine
    // to indicate that it's finished
    wg.Done()
}
