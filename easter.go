package calendar

import (
	"errors"
	"time"
)

// Spencer Jones' formula from 1922
func easterDaySpencerJones(year int) (month, day int) {
	// Source: http://no.wikipedia.org/wiki/Påskeformelen
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	n := (h + l - 7*m + 114) / 31
	p := (h + l - 7*m + 114) % 31
	return n, p + 1
}

// Gauss's method for finding Easter day for a given year
func easterDayGauss(year int) (month, day int, err error) {
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

// Checks if a day is at Easter day +- a few days
func atEasterPlus(date time.Time, days int) bool {
	year := date.Year()
	eastermonth, easterday := easterDaySpencerJones(year)
	easter := time.Date(year, time.Month(eastermonth), easterday, 0, 0, 0, 0, time.UTC)
	when := easter.AddDate(0, 0, days)
	return atDate(date, when)
}

// Returns the Easter day for any given year
func EasterDay(year int) time.Time {
	month, day := easterDaySpencerJones(year)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
