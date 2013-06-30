package moskus

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
// This date is for the Norwegian transition to summertime
func atSommertid(date time.Time) bool {
	afterMarch := time.Date(date.Year(), time.April, 1, 0, 0, 0, 0, time.UTC)
	lastSundayInMarch, _ := searchBackwardsForSunday(afterMarch)
	return atDate(date, lastSundayInMarch)
}

// Checks if the given date is at the last sunday in October.
// (Transition to wintertime, adjust watches one hour back)
// This date is for the Norwegian transition to wintertime
func atVintertid(date time.Time) bool {
	afterOctober := time.Date(date.Year(), time.November, 1, 0, 0, 0, 0, time.UTC)
	lastSundayInOctober, _ := searchBackwardsForSunday(afterOctober)
	return atDate(date, lastSundayInOctober)
}

// Norwegian Mother's day, 2nd Sunday in February
func atMorsdag(date time.Time) bool {
	return atNthWeekdayOfMonth(date, 2, time.Sunday, time.February)
}

// Norwegian Father's day, 2nd Sunday in November
func atFarsdag(date time.Time) bool {
	return atNthWeekdayOfMonth(date, 2, time.Sunday, time.November)
}

// Spring equinox
func atNorthwardEquinox(date time.Time) bool {
	if date.Month() != time.March {
		return false
	}
	return atDate(date, northwardEquinox(date.Year()))
}

// Summer solstice
func atNorthernSolstice(date time.Time) bool {
	if date.Month() != time.June {
		return false
	}
	return atDate(date, northernSolstice(date.Year()))
}

// Autumn equinox
func atSouthwardEquinox(date time.Time) bool {
	if date.Month() != time.September {
		return false
	}
	return atDate(date, southwardEquinox(date.Year()))
}

// Winter solstice
func atSouthernSolstice(date time.Time) bool {
	if date.Month() != time.December {
		return false
	}
	return atDate(date, southernSolstice(date.Year()))
}

// Check if the given date is at the Nth weekday (for istance Sunday) of a given month
func atNthWeekdayOfMonth(date time.Time, n int, weekday time.Weekday, month time.Month) bool {
	if date.Month() != month {
		return false
	}
	nthDay, err := nthWeekdayOfMonth(date, n, weekday)
	if err != nil {
		return false
	}
	if atDate(date, nthDay) {
		return true
	}
	return false
}

// Birthday of Dr. Martin Luther King, Jr.
// Third monday in January
func atDrMarthinLutherKingJrBirthday(date time.Time) bool {
	return atNthWeekdayOfMonth(date, 3, time.Monday, time.January)
}

