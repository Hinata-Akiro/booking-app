package input

import (
    "fmt"
    "strings"
)

type UserData struct {
    FirstName    string
    LastName     string
    Email        string
    TicketNumber int
}

func InputFields(booking *UserData, remainingTicketsAvailable int) (*UserData, error) {
    fmt.Println("Please input your firstName:")
    fmt.Scan(&booking.FirstName)

    if len(booking.FirstName) < 3 {
        return nil, fmt.Errorf("Please enter a valid first name")
    }

    fmt.Println("Please input your lastName:")
    fmt.Scan(&booking.LastName)

    if len(booking.LastName) < 3 {
        return nil, fmt.Errorf("Please enter a valid last name")
    }

    fmt.Println("Please input your email:")
    fmt.Scan(&booking.Email)

    if !strings.Contains(booking.Email, "@") {
        return nil, fmt.Errorf("Please enter a valid email")
    }

    fmt.Println("Please input your ticket number:")
    fmt.Scan(&booking.TicketNumber)

    if booking.TicketNumber > remainingTicketsAvailable {
        return nil, fmt.Errorf("Sorry, we have only %v tickets available", remainingTicketsAvailable)
    }

    return booking, nil
}
