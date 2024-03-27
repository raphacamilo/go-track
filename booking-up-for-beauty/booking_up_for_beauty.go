package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/02/2006 15:04:05"

	t, _ := time.Parse(layout, date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January _2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January _2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const layout = "1/_2/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	const layoutToFormat = "Monday, January 2, 2006, at 15:04"
	formattedDate := t.Format(layoutToFormat)

	return "You have an appointment on " + formattedDate + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
