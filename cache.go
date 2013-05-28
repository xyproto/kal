package moskus

import (
	"time"
)

// A CalendarCache wraps and caches a Calendar
type CalendarCache struct {
	cal          Calendar
	cacheRed     map[time.Time]string // red day description
	cacheNotable map[time.Time]string // notable day description
	cacheFlag    map[time.Time]bool   // flag flying day
}

// Creates a new CalendarCache that wraps and caches the given Calendar.
// A CalendarCache is also a Calendar itself, since it implements the
// Calendar interface.
func NewCalendarCache(cal Calendar) CalendarCache {
	var calca CalendarCache
	calca.cal = cal
	calca.cacheRed = make(map[time.Time]string)
	calca.cacheNotable = make(map[time.Time]string)
	calca.cacheFlag = make(map[time.Time]bool)
	return calca
}

// Wraps the RedDay function and caches some of the results
func (calca CalendarCache) RedDay(date time.Time) (bool, string, bool) {
	// Return from cache, if it's there
	desc, ok := calca.cacheRed[date]
	if ok {
		return ok, desc, calca.cacheFlag[date]
	}

	// Get the information from the calendar
	red, desc, flag := calca.cal.RedDay(date)

	// Add red days to the cache
	// TODO: Also cache non-red days
	if red {
		calca.cacheRed[date] = desc
		calca.cacheFlag[date] = flag
	}

	return red, desc, flag
}

// Wraps the NotableDay function and caches some of the results
func (calca CalendarCache) NotableDay(date time.Time) (bool, string, bool) {
	// Return from cache, if it's there
	desc, ok := calca.cacheNotable[date]
	if ok {
		return ok, desc, calca.cacheFlag[date]
	}

	// Get the information from the calendar
	notable, desc, flag := calca.cal.NotableDay(date)

	// Add notable days to cache
	// TODO: Also cache non-notable days
	if notable {
		calca.cacheNotable[date] = desc
		calca.cacheFlag[date] = flag
	}

	return notable, desc, flag
}

// --- These are here just to satisfy the Calendar interface ---

// Wraps the NotablePeriod function
func (calca CalendarCache) NotablePeriod(date time.Time) (bool, string) {
	return calca.cal.NotablePeriod(date)
}

// Wraps the DayName function
func (calca CalendarCache) DayName(date time.Weekday) string {
	return calca.cal.DayName(date)
}

// Wraps the NormalDay function
func (calca CalendarCache) NormalDay() string {
	return calca.cal.NormalDay()
}
