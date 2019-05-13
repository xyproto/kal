package calendar

// locale: en_US

// NOTE: Work in progress!

// Anything that's specific to the US, with the exception of some movable dates which are in movable.go
// This calendar has a corresponding locale code in the NewCalendar function in calendar.go

import (
	"time"
)

type USCalendar struct{}

// Create a new US calendar
func NewUSCalendar() USCalendar {
	return USCalendar{}
}

// Finds the US name for a day of the week.
// Note that time.Weekday starts at 0 with Sunday, not Monday.
func (nc USCalendar) DayName(day time.Weekday) string {
	return day.String()
}

// Finds the US name for a given month
func (nc USCalendar) MonthName(month time.Month) string {
	return month.String()
}

// Checks if a given date is a "red day" (public holiday) in the US calendar.
// Returns true/false, a description and true/false for if it's a flag day.
// The dates will never overlap.
func (nc USCalendar) RedDay(date time.Time) (bool, string, bool) {

	// Source: http://en.wikipedia.org/wiki/Public_holidays_in_the_United_States
	// Source: http://timpanogos.wordpress.com/flag-fly-dates/

	var (
		desc string
		flag bool
	)

	// Sundays
	if date.Weekday() == time.Sunday {
		desc = "Sunday"
	}

	// Election Day
	if atElectionDay(date) {
		desc = "Election Day"
		flag = false
	}

	// New Year's Day
	if atMD(date, 1, 1) {
		desc = "New Year's Day"
		flag = true
	}

	// Birthday of Dr. Martin Luther King, Jr.
	if atNthWeekdayOfMonth(date, 3, time.Monday, time.January) {
		desc = "Martin Luther King Day"
		flag = true
	}

	// Inauguration Day
	if atInaugurationDay(date) {
		desc = "Inauguration Day"
		flag = true
	}

	// Lincoln's birthday
	if atMD(date, 2, 12) {
		desc = "Lincoln's birthday"
		flag = true
	}

	// Washington's Birthday / Presidents' Day
	if atNthWeekdayOfMonth(date, 3, time.Monday, time.February) {
		desc = "Presidents' Day"
		flag = true
	}

	// Armed Forces Day
	if atNthWeekdayOfMonth(date, 3, time.Saturday, time.May) {
		desc = "Armed Forces Day"
		flag = true
	}

	// Memorial Day
	if atLastWeekday(date, time.Monday, time.May) {
		desc = "Memorial Day"
		flag = true
	}

	// 4th of July
	if atMD(date, 7, 4) {
		desc = "Independence Day"
		flag = true
	}

	// Labor Day
	if atNthWeekdayOfMonth(date, 1, time.Monday, time.September) {
		desc = "Labor Day"
		flag = true
	}

	// Columbus Day
	if atNthWeekdayOfMonth(date, 2, time.Monday, time.October) {
		desc = "Columbus Day"
		flag = true
	}

	// Veterans Day
	if atMD(date, 11, 11) {
		desc = "Veterans Day"
		flag = true
	}

	// Thanksgiving Day
	if atNthWeekdayOfMonth(date, 4, time.Thursday, time.November) {
		desc = "Thanksgiving Day"
		flag = true
	}

	// Christmas
	if atMD(date, 12, 25) {
		desc = "Christmas Day"
		flag = true
	}

	// Red days
	if desc != "" {
		// Then return
		return true, desc, flag
	}

	// Normal days
	return false, desc, false
}

// Some days are not red, but special in one way or another.
// Checks if a given date is notable. Returns true/false if the
// given date is notable, a comma separated description (in case there are more
// than one notable event that day) and true/false depending on if it's a flag
// flying day or not.
func (nc USCalendar) NotableDay(date time.Time) (bool, string, bool) {

	var (
	//descriptions []string
	//flag         bool
	)

	// Since days may overlap, flag flying days must come first.

	// --- Flag flying days ---

	// --- Other days ---

	// No notable events
	return false, "", false
}

// Checks if a given date is in a notable time range (summer holidays, for instance)
func (nc USCalendar) NotablePeriod(date time.Time) (bool, string) {
	// TODO:
	// Christmas time
	// summer/winter/spring/autumn time
	// etc
	return false, ""
}

// An ordinary day
func (nc USCalendar) NormalDay() string {
	return "Ordinary"
}
