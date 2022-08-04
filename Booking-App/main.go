package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			printFirstNames()

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First name or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ or .")
			}
			if !isValidTicketNumber {
				fmt.Print("Number of tickets you entered is invalid.")
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your ticket here to attend.")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Printf("These are all our bookings: %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask the user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email ID: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remianing for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Sending ticket: \n%v to email address %v\n", ticket, email)
	fmt.Println("----------------------------------------------------------------------")
	wg.Done()
}
