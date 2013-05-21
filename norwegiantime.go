package norwegiantime

import (
	"errors"
	"time"
	"fmt"
)

// Gauss's algorithm for finding easter day for a given year
func easterDay(year int) (month, day int, err error) {
	// Source: http://no.it.programmering.delphi.narkive.com/oDY0xYOW/algoritme-for-norske-bevegelige-helligdager
	if (year < 1583) || (year > 4199) {
		return 0, 0, errors.New("year out of range")
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
	month = 3
	if day > 31 {
		month++
		day -= 31
	}
	return month, day, nil
}

func EasterDate(year int) (time.Time, error) {
	month, day, err := easterDay(year)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), err
}

// Checks if the given time is easter day
func IsEaster(t time.Time) bool {
	eastermonth, easterday, err := easterDay(t.Year())
	if err != nil {
		// If the year is out of range, treat it like there isn't an easter day
		return false
	}
	return (t.Month() == time.Month(eastermonth)) && (t.Day() == easterday)
}

// Checks if the given time is at the given month and day
func At(t time.Time, month, day int) bool {
	return (t.Month() == time.Month(month)) && (t.Day() == day)
}

// Checks if the two given times are at the same months and days
func AtDate(t, when time.Time) bool {
	return (t.Month() == when.Month()) && (t.Day() == when.Day())
}

// Checks if a day is at easterday +- a few days
func AtEasterPlus(t time.Time, days int) bool {
	year := t.Year()
	eastermonth, easterday, err := easterDay(year)
	if err != nil {
		// If the year is out of range, treat it like there isn't an easter day
		return false
	}
	easter := time.Date(year, time.Month(eastermonth), easterday, 0, 0, 0, 0, time.UTC)
	when := easter.AddDate(0, 0, days)
	return AtDate(t, when)
}

// Return the number of sundays from day t, +- a few days
func SundaysInPeriod(t time.Time, days int) int {
	sundayCounter := 0
	when := t
	if days < 0 {
		for i := days; i <= 0; i++ {
			when = t.AddDate(0, 0, i)
			if when.Weekday() == time.Sunday {
				sundayCounter++
			}
		}
	} else {
		for i := 0; i <= days; i++ {
			when = t.AddDate(0, 0, i)
			if when.Weekday() == time.Sunday {
				sundayCounter++
			}
		}
	}
	return sundayCounter
}

// Check if a day is at easterday +- a few days, not counting sundays
// TODO: test
func AtEasterPlusNotCountingSundays(t time.Time, days int) bool {
	adjustedDays := days
	if days < 0 {
		adjustedDays += SundaysInPeriod(t, days)
	} else {
		adjustedDays -= SundaysInPeriod(t, days)
	}
	return AtEasterPlus(t, adjustedDays)
}

// Checks if the given year, month and day is a "red" in the Norwegian calendar
// Returns true/false and a description
func RedDay(year int, month time.Month, day int) (bool, string) {
	return RedDate(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// Find a preceeding sunday
func searchBackwardsForSunday(t time.Time) (time.Time, error) {
	// Start with the day before the given date
	current := t.AddDate(0, 0, -1)

	// Stay within the same year
	for current.Year() == t.Year() {
		// Check if it's a Sunday
		if current.Weekday() == time.Sunday {
			// Found one
			return current, nil
		}

		// Go the previous day
		current = current.AddDate(0, 0, -1)
	}

	// This should never happen, since we will only be searching
	// from easter day and backwards
	return t, errors.New("Could not find an earlier Sunday!")
}

// Checks if the given date is at Palmesøndag
// Palmesøndag is the Sunday before easter (Første Påskedag)
func AtPalmesøndag(t time.Time) bool {
	easter, err := EasterDate(t.Year())
	if err != nil {
		// If the year is out of range, treat it like there isn't an easter day
		return false
	}
	palmesøndag, err := searchBackwardsForSunday(easter)
	return AtDate(t, palmesøndag)
}

// Find the Nth sunday of a given year and month
func nthSundayOfMonth(t time.Time, n int) (time.Time, error) {

	sundaycounter := 0

	// Start at the first day in the given month
	current := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same month
	for current.Month() == t.Month() {

		// Is it a Sunday?
		if current.Weekday() == time.Sunday {
			sundaycounter++
		}

		// Is it the Nth sunday?
		if sundaycounter == n {
			return current, nil
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	return t, errors.New(fmt.Sprintf("Could not find the %dth Sunday in %s", n, t.Month().String()))

}

// Morsdag, 2nd Sunday in February
func AtMorsdag(t time.Time) bool {
	if t.Month() != time.February {
		return false
	}
	nthSunday, err := nthSundayOfMonth(t, 2)
	if err != nil {
		return false
	}
	if AtDate(t, nthSunday) {
		return true
	}
	return false
}

// Farsdag, 2nd Sunday in November
func AtFarsdag(t time.Time) bool {
	if t.Month() != time.November {
		return false
	}
	nthSunday, err := nthSundayOfMonth(t, 2)
	if err != nil {
		return false
	}
	if AtDate(t, nthSunday) {
		return true
	}
	return false
}

// Dates that are not red, but not completely ordinary either
func NotableDate(date time.Time) (bool, string) {

	// Askeonsdag (fasten begynner)
	if AtEasterPlus(date, -46) {
		return true, "Askeonsdag"
	}

	// Påskeaften (fasten slutter)
	if AtEasterPlus(date, -1) {
		return true, "Påskeaften"
	}

	// Fetetirsdag
	if AtEasterPlus(date, -47) {
		return true, "Fetetirsdag"
	}

	// Fastelavn
	if AtEasterPlus(date, -49) {
		return true, "Fastelavn"
	}

	// Fastelavnsøndag

	// Blåmandag

	// Morsdag
	if AtMorsdag(date) {
		return true, "Morsdag"
	}

	// Farsdag
	if AtFarsdag(date) {
		return true, "Farsdag"
	}

	// Sankthansaften
	if At(date, 6, 23) {
		return true, "Sankthansaften"
	}

	return false, ""
}

// Checks if a given date is a "red day" in the Norwegian calendar
// Returns true/false and a description
func RedDate(date time.Time) (bool, string) {
	// Source: http://www.diskusjon.no/index.php?showtopic=1084239
	// Source: http://no.wikipedia.org/wiki/Helligdager_i_Norge

	// Første nyttårsdag, 1. januar
	if At(date, 1, 1) {
		return true, "Første nyttårsdag"
	}

	// Palmesøndag
	if AtPalmesøndag(date) {
		return true, "Palmesøndag"
	}

	// Skjærtorsdag (easter - 3d)
	if AtEasterPlus(date, -3) {
		return true, "Skjærtorsdag"
	}

	// Langfredag (easter - 2d)
	if AtEasterPlus(date, -2) {
		return true, "Langfredag"
	}

	// Første påskedag
	if AtEasterPlus(date, 0) {
		return true, "Første påskedag"
	}

	// Andre påskedag (easter + 1d)
	if AtEasterPlus(date, 1) {
		return true, "Andre påskedag"
	}

	// Arbeidernes internasjonale kampdag, 1. mai
	if At(date, 5, 1) {
		return true, "Arbeidernes internasjonale kampdag"
	}

	// Grunnlovsdagen, 17. mai
	if At(date, 5, 17) {
		return true, "Grunnlovsdagen"
	}

	// Kristi himmelfartsdag (40. påskedag: easter + 39d)
	if AtEasterPlus(date, 39) {
		return true, "Kristi himmelfartsdag"
	}

	// Første pinsedag (50. påskedag: easter + 49d)
	if AtEasterPlus(date, 49) {
		return true, "Første pinsedag"
	}

	// Andre pinsedag (51. påskedag: easter + 50d)
	if AtEasterPlus(date, 50) {
		return true, "Andre pinsedag"
	}

	// Første juledag (25. desember)
	if At(date, 12, 25) {
		return true, "Første juledag"
	}

	// Andre juledag (26. desember)
	if At(date, 12, 26) {
		return true, "Andre juledag"
	}

	// Sundays
	if date.Weekday() == 0 {
		return true, "Søndag"
	}

	// Normal days
	return false, ""
}

// Finds the norwegian name for a day of the week
// Note that time.Weekday starts at 0 with Sunday, not Monday
func NorwegianName(dayoftheweek time.Weekday) string {
	days := []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}
	return days[int(dayoftheweek)]
}
