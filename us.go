package moskus

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
	if date.Weekday() == 0 {
		desc = "Sunday"
	}

	// New Year's Day
	if atMD(date, 1, 1) {
		desc = "New Year's Day"
		flag = true
	}

	// Birthday of Dr. Martin Luther King, Jr.
	if atDrMartinLutherKingJrBirthday(date) {
		desc = "Birthday of Dr. Martin Luther King, Jr."
		flag = true
	}

	// TODO: The rest

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

	// Source: http://www.timeanddate.no/kalender/merkedag-innhold
	// Source: http://no.wikipedia.org/wiki/Norges_offisielle_flaggdager

	var (
		//descriptions []string
		//flag         bool
	)

	// Since days may overlap, "flaggdager" must come first for the flag
	// flying days to be correct.

	// --- Flag days ---

	// --- Non-flag days ---

	// No notable events
	return false, "", false
}

// Checks if a given date is in a notable time range (summer holidays, for instance)
func (nc USCalendar) NotablePeriod(date time.Time) (bool, string) {
	// TODO:
	// uke 28, 29 og 30, fellesferie
	// jul
	// fastelavn
	// faste
	// sommer/vinter/vår/høst (legg denne først i listen)
	return false, ""
}

// An ordinary day
func (nc USCalendar) NormalDay() string {
	return "Ordinary"
}
