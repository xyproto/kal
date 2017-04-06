package calendar

// Movable dates

import (
	"time"
)

// Checks if the given date is at Palm Sunday (the Sunday before Easter)
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

// Inauguration day. 21st of January, unless if it is a sunday, then it's the 20th.
func atInaugurationDay(date time.Time) bool {
	// Election day, 2000, 2004, 2008, 2012 etc
	if (date.Year() % 4) == 0 {
		// Normally on the 21st
		if atMD(date, 1, 21) {
			if date.Weekday() != time.Sunday { // not a sunday
				return true
			}
		}
		// The day before, if the 21st is a sunday
		if atMD(date, 1, 20) {
			// if tomorrow is a sunday, this is election day
			if date.AddDate(0, 0, 1).Weekday() == time.Sunday {
				return true
			}
		}
	}
	return false
}

// The Tuesday following the first Monday in November
func atElectionDay(date time.Time) bool {
	// Ensure that we are in November
	if date.Month() != time.November {
		return false
	}

	// Find the first Monday in November
	monday, err := nthWeekdayOfMonth(date, 1, time.Monday)
	if err != nil {
		return false
	}

	// Find the following Tuesday
	tuesday := monday.AddDate(0, 0, 1)

	// Compare with given date
	return atDate(date, tuesday)
}
