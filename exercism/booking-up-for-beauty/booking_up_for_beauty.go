package booking

import (
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
// "October 3, 2019 20:32:00"
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

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
// "Thursday, July 25, 2019 13:45:00"
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Cannot parse the date.")
	}

	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
// "7/25/2019 13:45:00"
func Description(date string) string {
	layoutDateIn := "1/2/2006 15:04:05"
	layoutDateOut := "Monday, January 2, 2006, at 15:04"

	t, err := time.Parse(layoutDateIn, date)
	if err != nil {
		panic("Cannot parse the date.")
	}

	return "You have an appointment on " + t.Format(layoutDateOut) + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year, _, _ := time.Now().Date()
	month := time.September
	day := 15
	anniversaryDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return anniversaryDate
}
