package moskus

import (
	"errors"
	"fmt"
	"time"
)

// Checks if the given time is at the given month and day
func atMD(date time.Time, month, day int) bool {
	return (date.Month() == time.Month(month)) && (date.Day() == day)
}

// Checks if the two given times are at the same months and days
func atDate(t, when time.Time) bool {
	return (t.Month() == when.Month()) && (t.Day() == when.Day())
}

// Return the count of a given weekday from day t, +- a few days
func numberOfWeekdaysInPeriod(date time.Time, days int, whichWeekday time.Weekday) int {
	specialWeekdayCounter := 0
	when := date
	if days < 0 {
		for i := days; i <= 0; i++ {
			when = date.AddDate(0, 0, i)
			if when.Weekday() == whichWeekday {
				specialWeekdayCounter++
			}
		}
	} else {
		for i := 0; i <= days; i++ {
			when = date.AddDate(0, 0, i)
			if when.Weekday() == whichWeekday {
				specialWeekdayCounter++
			}
		}
	}
	return specialWeekdayCounter
}

// Return the number of sundays from day t, +- a few days
func sundaysInPeriod(date time.Time, days int) int {
	return numberOfWeekdaysInPeriod(date, days, time.Sunday)
}

// Find a preceeding sunday
func searchBackwardsForSunday(date time.Time) (time.Time, error) {
	// Start with the day before the given date
	current := date.AddDate(0, 0, -1)

	// Stay within the same year
	for current.Year() == date.Year() {
		// Check if it's a Sunday
		if current.Weekday() == time.Sunday {
			// Found one
			return current, nil
		}

		// Go the previous day
		current = current.AddDate(0, 0, -1)
	}

	return date, errors.New("Could not find an earlier Sunday the same year!")
}

// Find the Nth type of weekday of a given year and month
func nthWeekdayOfMonth(date time.Time, n int, whichWeekday time.Weekday) (time.Time, error) {

	specialWeekdayCounter := 0

	// Start at the first day in the given month
	current := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same month
	for current.Month() == date.Month() {

		// Which weekday is it?
		if current.Weekday() == whichWeekday {
			specialWeekdayCounter++
		}

		// Is it the Nth occurance?
		if specialWeekdayCounter == n {
			return current, nil
		}

		// If it's the given weekday, advance almost one week forward
		if current.Weekday() == whichWeekday {
			current = current.AddDate(0, 0, 7)
			continue
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	return date, fmt.Errorf("Could not find the %dth %s in %s!", n, whichWeekday, date.Month())
}

// Find the Nth sunday of a given year and month
func nthSundayOfMonth(date time.Time, n int) (time.Time, error) {
	return nthWeekdayOfMonth(date, n, time.Sunday)
}
