package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Golang Programming Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // list of dictionaries or maps

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("######################################")
	greetUsers()
	fmt.Println("######################################")

	firstName, lastName, email, userTickets := getUserInputs()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(lastName,
		firstName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) // concurrency using goroutines

		// call function print first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered is not valid")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// fmt.Println(&remainingTickets)

	// ask user for their name
	fmt.Println("Please enter your first name:")
	// read user input
	fmt.Scan(&firstName) // passing the memory address of the variable
	fmt.Println("Please enter your last name:")
	// read user input
	fmt.Scan(&lastName) // passing the memory address of the variable
	fmt.Println("Please enter your email address:")
	// read user input
	fmt.Scan(&email) // passing the memory address of the variable
	fmt.Println("Enter number of tickets:")
	// read user input
	fmt.Scan(&userTickets) // passing the memory address of the variable

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for the user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Println(bookings)
	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n",
		firstName, lastName, userTickets, email)

	fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("######################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("######################################")
	wg.Done()
}
