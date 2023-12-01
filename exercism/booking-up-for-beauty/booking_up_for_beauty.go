package main //booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// "7/25/2019 13:45:00"
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Cannot parse the date.")
	}
	return t
}

// HasPassed returns whether a date has passed.
// "July 25, 2019 13:45:00"
func HasPassed(date string) bool {
	layout := "January 02, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Cannot parse the date.")
	}

	if t.Before(time.Now()) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}

func main() {
	fmt.Println(Schedule("7/25/2019 13:45:00"))
	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
}
