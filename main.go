package main

import (
    "fmt"
    "booking-app/input"
)


func main() {
    const ticketNumber = 50
    var remainingTicketsAvailable = 50
    bookings := []input.UserData{}

    welcomeMessage(ticketNumber, remainingTicketsAvailable)

    for i := 0; i <= ticketNumber; i++ {
		userData := input.UserData{}
		booking, err := input.InputFields(&userData, remainingTicketsAvailable)

        if err != nil {
            fmt.Println(err)
            continue
        }

        bookings = append(bookings, *booking)
        remainingTicketsAvailable = remainingTicketsAvailable - booking.TicketNumber

        if remainingTicketsAvailable == 0 {
            fmt.Println("No tickets available")
            continue
        }

        fmt.Printf("Hello %v, you have booked %v of tickets we would send you a mail %v \n", booking.FirstName, booking.TicketNumber, booking.Email)

        fmt.Printf("There are now %v tickets remaining\n", remainingTicketsAvailable)

        fmt.Printf("Your bookings are %v\n", bookings)
    }

}

func welcomeMessage(ticketNumber int, remainingTicketsAvailable int) {
    conference := "Go conferences"
    fmt.Printf("Welcome to %v applications!\n", conference)
    fmt.Printf("We have total of %v tickets available and %v remaining tickets available", ticketNumber, remainingTicketsAvailable)
    fmt.Println("get your tickets here to attend...")
}
