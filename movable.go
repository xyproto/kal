package norwegiantime

// Movable dates

import (
	"time"
)

// Checks if the given date is at Palm Sunday (the Sunday before easter)
func atPalmSunday(date time.Time) bool {
	easter := EasterDay(date.Year())
	// There will always be a sunday before easter for any given year,
	// so we don't need to check the value of err
	palmSunday, _ := searchBackwardsForSunday(easter)
	return atDate(date, palmSunday)
}

// Checks if the given date is at the last sunday in March.
// (Transition to summertime, adjust watches one hour ahead)
func atSummerTime(date time.Time) bool {
	afterMarch := time.Date(date.Year(), time.April, 1, 0, 0, 0, 0, time.UTC)
	lastSundayInMarch, _ := searchBackwardsForSunday(afterMarch)
	return atDate(date, lastSundayInMarch)
}

// Checks if the given date is at the last sunday in October.
// (Transition to wintertime, adjust watches one hour back)
func atWinterTime(date time.Time) bool {
	afterOctober := time.Date(date.Year(), time.November, 1, 0, 0, 0, 0, time.UTC)
	lastSundayInOctober, _ := searchBackwardsForSunday(afterOctober)
	return atDate(date, lastSundayInOctober)
}

// Morsdag, 2nd Sunday in February
func atMorsdag(date time.Time) bool {
	if date.Month() != time.February {
		return false
	}
	nthSunday, err := nthSundayOfMonth(date, 2)
	if err != nil {
		return false
	}
	if atDate(date, nthSunday) {
		return true
	}
	return false
}

// Farsdag, 2nd Sunday in November
func atFarsdag(date time.Time) bool {
	if date.Month() != time.November {
		return false
	}
	nthSunday, err := nthSundayOfMonth(date, 2)
	if err != nil {
		return false
	}
	if atDate(date, nthSunday) {
		return true
	}
	return false
}

// Vårjevndøgn
func atNorthwardEquinox(date time.Time) bool {
	if date.Month() != time.March {
		return false
	}
	return atDate(date, northwardEquinox(date.Year()))
}

// Sommersolverv
func atNorthernSolstice(date time.Time) bool {
	if date.Month() != time.June {
		return false
	}
	return atDate(date, northernSolstice(date.Year()))
}

// Høstjevndøgn
func atSouthwardEquinox(date time.Time) bool {
	if date.Month() != time.September {
		return false
	}
	return atDate(date, southwardEquinox(date.Year()))
}

// Vintersolverv
func atSouthernSolstice(date time.Time) bool {
	if date.Month() != time.December {
		return false
	}
	return atDate(date, southernSolstice(date.Year()))
}
