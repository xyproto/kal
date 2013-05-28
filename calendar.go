package moskus

import (
	"errors"
	"time"
)

// TODO: Add support for other languages and locales

// A common interface for calendars for all languages and locales
type Calendar interface {
	DayName(time.Weekday) string
	RedDay(time.Time) (bool, string, bool)
	NotableDay(time.Time) (bool, string, bool)
	NormalDay() string
	NotablePeriod(date time.Time) (bool, string)
}

/* Creates a new calendar based on a given langauge string
 * Supported strings:
 *   nb_NO (Norwegian Bokmål)
 * The calendar can be cached for faster lookups
 */
func NewCalendar(locCode string, cache bool) (Calendar, error) {
	var (
		cal       Calendar
		supported bool = true
	)

	// Find the corresponding calendar struct for the given locale
	switch locCode {
	case "nb_NO":
		cal = NewNorwegianCalendar()
	default:
		supported = false
	}

	if !supported {
		// Return an error
		return cal, errors.New("Locale not supported: " + locCode)
	}
	if !cache {
		// Return a calendar without cache
		return cal, nil
	}
	// Return a cached calendar cache
	return NewCalendarCache(cal), nil
}

// Checks if a given date is a flying flag day or not
func FlagDay(cal Calendar, date time.Time) bool {
	return thirdBool(date, cal.RedDay) || thirdBool(date, cal.NotableDay)
}

// Describe what type of day a given date is
func Describe(cal Calendar, date time.Time) string {
	fulldesc := ""
	if red, desc, _ := cal.RedDay(date); red {
		fulldesc = desc
	}
	if notable, desc, _ := cal.NotableDay(date); notable {
		if fulldesc == "" {
			fulldesc = desc
		} else {
			fulldesc += ", " + desc
		}
	}
	if fulldesc != "" {
		return fulldesc
	}
	return cal.NormalDay()
}

// Get the week number, from 1 to 53
func WeekNum(date time.Time) int {
	_, weeknum := date.ISOWeek()
	return weeknum
}
