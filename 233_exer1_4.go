package main

import (
	"fmt"
)

// Train struct represents information about a train
type Train struct {
	ID          int
	Name        string
	Destination string
	Departure   string
}

// Ticket struct represents information about a ticket
type Ticket struct {
	TrainID int
	Seat    int
}

func main() {
	// Creating a map to store train schedules
	trainSchedules := make(map[int]Train)
	// Adding sample train schedules
	trainSchedules[1] = Train{ID: 1, Name: "Rani Channamma", Destination: "Bengaluru", Departure: "08:00"}
	trainSchedules[2] = Train{ID: 2, Name: "Shatabdi", Destination: "Mysuru", Departure: "09:30"}
	trainSchedules[3] = Train{ID: 3, Name: "Rajdhani", Destination: "Mangaluru", Departure: "11:00"}

	// Creating slices to manage tickets for different trains
	tickets := make(map[int][]Ticket)

	for {
		fmt.Println("\nRailway Ticket Booking System")
		fmt.Println("1. View Trains")
		fmt.Println("2. Book Ticket")
		fmt.Println("3. View Tickets")
		fmt.Println("4. Delete Ticket")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nAvailable Trains:")
			for _, train := range trainSchedules {
				fmt.Printf("Train ID: %d, Name: %s, Destination: %s, Departure: %s\n", train.ID, train.Name, train.Destination, train.Departure)
			}
		case 2:
			var trainID, seatNumber int
			fmt.Print("\nEnter Train ID: ")
			fmt.Scanln(&trainID)
			fmt.Print("Enter Seat Number(1-50): ")
			fmt.Scanln(&seatNumber)

			if _, ok := trainSchedules[trainID]; ok {
				if seatNumber > 0 && seatNumber <= 50 {
					tickets[trainID] = append(tickets[trainID], Ticket{TrainID: trainID, Seat: seatNumber})
					fmt.Println("Ticket booked successfully!")
				} else {
					fmt.Println("Invalid seat number. Please choose another seat.")
				}
			} else {
				fmt.Println("Train not found in the schedules.")
			}
		case 3:
			fmt.Println("\nBooked Tickets:")
			for trainID, trainTickets := range tickets {
				for _, ticket := range trainTickets {
					fmt.Println("Train ID: ", trainID)
					fmt.Printf("Seat Number: %d\n", ticket.Seat)
				}
			}
		case 4:
			var trainID, seatNumber int
			fmt.Print("\nEnter Train ID to delete ticket: ")
			fmt.Scanln(&trainID)
			fmt.Print("Enter Seat Number to delete ticket: ")
			fmt.Scanln(&seatNumber)

			if _, ok := trainSchedules[trainID]; ok {
				ticketFound := false
				for i, ticket := range tickets[trainID] {
					if ticket.Seat == seatNumber {
						// Delete the ticket
						tickets[trainID] = append(tickets[trainID][:i], tickets[trainID][i+1:]...)
						fmt.Println("Ticket deleted successfully!")
						ticketFound = true
						break
					}
				}
				if !ticketFound {
					fmt.Println("Ticket not found.")
				}
			} else {
				fmt.Println("Train not found in the schedules.")
			}

		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
