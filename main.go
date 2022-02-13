package main

import (
	"fmt"
	"time"
	"VisitBookingApp/helper"
)

const officeName = "My Office"
const maxpermittedEntriesPerDay uint = 10
var remainingEntries uint = 10
var visits = make([]VisitorData,0)

type VisitorData struct{
	firstName string
	lastName string
	email string
}

func main() {

	welcomeUsers()

	fmt.Printf("This is %v visit booking app and maximum %v times visit per day is allowed.\n", officeName, maxpermittedEntriesPerDay)

	for {
		firstName, lastName, email := gatherUserInputs()

		isvalidName, isValidEmail := helper.IsValidInputs(firstName, lastName, email)
		
		if isvalidName && isValidEmail{
			bookVisitor(firstName, lastName, email)
			
			fmt.Printf("First names of visitors for today are : %v\n", printVisitors())

			isEmptySlot := isBookingComplete()
			if isEmptySlot {
				fmt.Printf("Visiting slots are already booked. Please come tomarrow.\n")
				break
			}
		}	else {
				
			if !isvalidName {
				fmt.Printf("Invalid name. Please try again.\n")
			}
			
			if !isValidEmail {
				fmt.Printf("Invalid email. Please try again.\n")
			}
		}
	}
}

func welcomeUsers() {
	fmt.Printf("Welcome user at %v.\n", officeName)
}

func printVisitors() []string {
	firstNames := []string{}
	for _ , visit := range visits{
		firstNames = append(firstNames, visit.firstName)
	}

	return firstNames	
}

func gatherUserInputs() (string, string, string){	
	var firstName string
	var lastName string
	var email string

	fmt.Println("Print enter first name :")
	fmt.Scan(&firstName)

	fmt.Println("Print enter last name :")
	fmt.Scan(&lastName)		

	fmt.Println("Print enter email address :")
	fmt.Scan(&email)

	fmt.Printf("Date time : %v \n", time.Now().Format("2000.01.01 18:00:00"))	

	return firstName, lastName, email
}

func bookVisitor(firstName string, lastName string, email string) {
	var visitorData = VisitorData {
		firstName : firstName,
		lastName : lastName,
		email : email,
	}

	remainingEntries = remainingEntries - 1
	visits = append(visits, visitorData)
	fmt.Printf("Thank you %v for visiting %v.\n", firstName, officeName)
	fmt.Printf("Total visits for today: %v\n", len(visits))
	fmt.Printf("Remaining visits for today: %v\n", remainingEntries)

}

func isBookingComplete() bool{
	isEmptySlot := remainingEntries == 0
	return isEmptySlot
}