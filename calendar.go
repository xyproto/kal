// Kal is a calendar package for finding public holidays ("red days"), Easter, notable days, equinoxes, solstices and flag flying days.
package kal

import (
	"errors"
	"time"
)

// Calendar provides a common interface for calendars of all languages
// and locales
type Calendar interface {
	DayName(time.Weekday) string
	RedDay(time.Time) (bool, string, bool)
	NotableDay(time.Time) (bool, string, bool)
	NormalDay() string
	NotablePeriod(time.Time) (bool, string)
	MonthName(time.Month) string
	MondayFirst() bool
}

/* Create a new calendar based on a given language string.
 *
 *  Supported strings:
 *  nb_NO (Norwegian Bokmål)
 *  en_US (US English)
 *  tr_TR (Turkish)
 *
 *  The calendar can be cached for faster lookups
 */
func NewCalendar(locCode string, cache bool) (cal Calendar, err error) {
	// Find the corresponding calendar struct for the given locale
	switch locCode {
	case "nb_NO":
		cal = NewNorwegianCalendar()
	case "en_US":
		cal = NewUSCalendar()
	case "tr_TR":
		cal = NewTRCalendar()
	default:
		return cal, errors.New("Locale not supported: " + locCode)
	}
	if cache {
		// Return a calendar with cache
		return NewCachedCalendar(cal), nil
	}
	return cal, nil
}

// Returns the third boolean argument given a time.Time value and
// a function that takes a time.Time and returns a bool, a string and a bool
func thirdBool(date time.Time, fn func(time.Time) (bool, string, bool)) bool {
	_, _, b := fn(date)
	return b
}

// Checks if a given date is a flag flying day or not
func FlagDay(cal Calendar, date time.Time) bool {
	return thirdBool(date, cal.RedDay) || thirdBool(date, cal.NotableDay)
}

// Checks if a given date is a "red" day or not
func RedDay(cal Calendar, date time.Time) bool {
	return thirdBool(date, cal.RedDay)
}

// Checks if a given date is a notable day or not
func NotableDay(cal Calendar, date time.Time) bool {
	return thirdBool(date, cal.NotableDay)
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

// Return a space separated string of the two first letters of every weekday
func TwoLetterDays(cal Calendar, mondayFirst bool) string {
	var (
		i time.Weekday
		s string
	)
	if !mondayFirst {
		for i = 0; i < 7; i++ {
			if i != 0 {
				s += " "
			}
			s += string([]rune(cal.DayName(i))[:2])
		}
	} else {
		for i = 1; i < 7; i++ {
			if i != 1 {
				s += " "
			}
			s += string([]rune(cal.DayName(i))[:2])
		}
		s += " " + string([]rune(cal.DayName(time.Weekday(0)))[:2])
	}
	return s
}

// Get the week number, from 1 to 53
func WeekNum(date time.Time) int {
	_, weeknum := date.ISOWeek()
	return weeknum
}
