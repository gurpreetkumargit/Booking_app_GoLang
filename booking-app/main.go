package main

import (
	"fmt"
	"sync"
	"time"
)

// data (variables) for our conference

const conferenceTickets int = 50

var conferenceName string = "GO conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName, lastName, email string
	numOfTickets               uint
}

var wg = sync.WaitGroup{}

func main() {

	// calling greet func
	greetUsers()

	// infinite loop  (breaks only if our tickets completely sold)
	// for {

	firstName, lastName, email, userTickets := getUserInput()

	// checking validation (require for ticket booking) using func
	isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// if all validation of user input is correct then confirm booking
	if isValidName && isValidEmail && isValidTicket {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)

		go sendTicket(firstName, lastName, userTickets, email)
		// calling firstName func and storing return value in new variable
		first_Names := getFirstName()

		fmt.Printf("firstName of booked users: %v\n", first_Names)

		// if tickets are sold out then break out ticket booking loop
		if remainingTickets == 0 {
			fmt.Println("\nall tickets are sold out. please try next year.")
			// break
		}
	} else {

		if !isValidName {
			fmt.Println("first or last name you filled is too short.")
		}
		if !isValidEmail {
			fmt.Println("your email is not correct.")
		}
		if !isValidTicket {
			fmt.Println("The ticket counts you entered is incorrect.")
		}
		fmt.Println("Please fill correct information .")
	}

	wg.Wait()
}

// }

// greeting to our users
func greetUsers() {

	fmt.Println("welcome to our", conferenceName, "booking application")
	fmt.Println("get your ticket to attend")
	fmt.Printf("we have %d tickets and %d tickets remaining\n", conferenceTickets, remainingTickets)

}

// get user input
// declare the empty variable (will be filled by user input)
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// taking input from user
	fmt.Println("Enter your firstName: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no. of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// book ticket

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// count remaining tickets
	remainingTickets = remainingTickets - userTickets

	// creating map

	userData := userData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		numOfTickets: userTickets}

	// filling booking array with users data
	bookings = append(bookings, userData)
	fmt.Printf("list of all bookings : %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation on your %v email\n", firstName, lastName, userTickets, email)

	fmt.Printf("remaining %v tickets for %v\n", remainingTickets, conferenceName)

}

// getting firstName of users

func getFirstName() []string {

	first_Names := []string{}

	// using for loop as for-each loop to take out only first name of user
	for _, booking := range bookings {
		first_Names = append(first_Names, booking.firstName)
	}

	return first_Names

}

func sendTicket(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(8 * time.Second)
	var ticketEmail = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println(" ")
	fmt.Println("###################")
	fmt.Printf("Sending Ticket:\n%v \nto %v\n", ticketEmail, email)
	fmt.Println("###################")
	fmt.Println(" ")

	wg.Done()
}
