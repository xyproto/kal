package norwegiantime

import (
	"time"
)

// A common interface for calendars for all languages and locales
type Calendar interface {
	DayName(time.Weekday) string
	RedDay(time.Time) (bool, string, bool)
	NotableDay(time.Time) (bool, string, bool)
	NormalDay() string
	NotablePeriod(date time.Time) (bool, string)
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
