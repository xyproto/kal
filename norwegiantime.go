package norwegiantime

// TODO: Craft a nicer API
// TODO: Add tests in norwegiantime_test.go

import (
	"time"
)

// Gauss's algorithm
// From: http://no.it.programmering.delphi.narkive.com/oDY0xYOW/algoritme-for-norske-bevegelige-helligdager
func EasterDay(year int) (month, day int) {
	if (year < 1583) || (year > 4199) {
		return 0, 0
	}
	g := (year % 19) + 1            // Golden year number
	c := (year / 100) + 1           // Century number
	x := (3 * c >> 2) - 12          // Lost leap years correction
	z := ((8*c + 5) / 25) - 5       // Moon's orbit correction
	d := ((5 * year) >> 2) - x - 10 // Find a Sunday in March
	e := (11*g + 20 + z - x) % 30   // Epact
	if (e == 24) || ((e == 25) && (g > 11)) {
		e++
	}
	n := 44 - e // Full Moon
	if n < 21 {
		n += 30
	}
	day = n + 7 - ((d + n) % 7) // Advance to Sunday
	if day > 31 {
		month = 4
		day -= 31
	} else {
		month = 3
	}
	return month, day
}

// Checks if the given time is easter day
func IsEasterDay(t time.Time) bool {
	eastermonth, easterday := EasterDay(t.Year())
	return (t.Month() == time.Month(eastermonth)) && (t.Day() == easterday)
}

// Checks if the given time is at the given month and day
func At(t time.Time, month, day int) bool {
	return (t.Month() == time.Month(month)) && (t.Day() == day)
}

// Checks if the two given times are at the same months and days
func AtDay(t, when time.Time) bool {
	return (t.Month() == when.Month()) && (t.Day() == when.Day())
}

// Checks if a day is at easterday +- a few days
func AtEasterPlus(year int, t time.Time, days int) bool {
	const (
		daysec = 3600 * 24
	)
	eastermonth, easterday := EasterDay(year)
	easter := time.Date(year, time.Month(eastermonth), easterday, 0, 0, 0, 0, time.UTC)
	when := easter.AddDate(0, 0, days)
	return AtDay(t, when)
}

// Norwegian red days
// From: http://www.diskusjon.no/index.php?showtopic=1084239
func RedDay(year, month, day int) bool {

	// The time to check
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// 1. januar
	if At(t, 1, 1) {
		return true
	}

	// Palmesøndag (easter - 7d)
	if AtEasterPlus(year, t, -7) {
		return true
	}

	// Skjertorsdag (easter - 3d)
	if AtEasterPlus(year, t, -3) {
		return true
	}

	// Langfredag (easter - 2d)
	if AtEasterPlus(year, t, -2) {
		return true
	}

	// 1. påskedag
	if AtEasterPlus(year, t, 0) {
		return true
	}

	// 2. påskedag (easter + 1d)
	if AtEasterPlus(year, t, 1) {
		return true
	}

	// 1. mai
	if At(t, 5, 1) {
		return true
	}

	// 17. mai
	if At(t, 5, 17) {
		return true
	}

	// Kristi himmelfartsdag (40. påskedag: easter + 39d)
	if AtEasterPlus(year, t, 39) {
		return true
	}

	// 1. pinsedag (50. påskedag: easter + 49d)
	if AtEasterPlus(year, t, 49) {
		return true
	}

	// 2. pinsedag (51. påskedag: easter + 50d)
	if AtEasterPlus(year, t, 50) {
		return true
	}

	// 1. juledag (25. desember)
	if At(t, 12, 25) {
		return true
	}

	// 2. juledag (26. desember)
	if At(t, 12, 26) {
		return true
	}

	// Søndag
	if t.Weekday() == 0 {
		return true
	}

	// Hverdag
	return false
}

// TODO: Take a time.Weekday instead of an int
func N2d(dayoftheweek int, language string) string {
	var days []string
	if language == "no" {
		days = []string{"mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag", "søndag"}
	} else {
		return time.Weekday(dayoftheweek).String()
	}
	if (dayoftheweek >= 0) && (dayoftheweek <= 6) {
		// TODO: Out of range check
		return days[dayoftheweek]
	}
	return "INVALID DAY"
}


