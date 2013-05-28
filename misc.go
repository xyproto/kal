package norwegiantime

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

// Return the number of sundays from day t, +- a few days
func sundaysInPeriod(date time.Time, days int) int {
	sundayCounter := 0
	when := date
	if days < 0 {
		for i := days; i <= 0; i++ {
			when = date.AddDate(0, 0, i)
			if when.Weekday() == time.Sunday {
				sundayCounter++
			}
		}
	} else {
		for i := 0; i <= days; i++ {
			when = date.AddDate(0, 0, i)
			if when.Weekday() == time.Sunday {
				sundayCounter++
			}
		}
	}
	return sundayCounter
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

// Find the Nth sunday of a given year and month
func nthSundayOfMonth(date time.Time, n int) (time.Time, error) {

	sundaycounter := 0

	// Start at the first day in the given month
	current := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same month
	for current.Month() == date.Month() {

		// Is it a Sunday?
		if current.Weekday() == time.Sunday {
			sundaycounter++
		}

		// Is it the Nth sunday?
		if sundaycounter == n {
			return current, nil
		}

		// If it's a sunday, advance almost one week forward
		if current.Weekday() == time.Sunday {
			current = current.AddDate(0, 0, 7)
			continue
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	return date, errors.New(fmt.Sprintf("Could not find the %dth Sunday in %s!", n, date.Month()))
}

// Get the week number, from 1 to 53
func WeekNum(date time.Time) int {
	_, weeknum := date.ISOWeek()
	return weeknum
}

// Returns the third boolean argument given a time.Time value and
// a function that takes a time.Time and returns a bool, a string and a bool
func thirdBool(date time.Time, fn func(time.Time) (bool, string, bool)) bool {
	_, _, b := fn(date)
	return b
}
