package calendar

// locale: tr_TR

// NOTE: Work in progress!
// This calendar has a corresponding locale code in the NewCalendar function in calendar.go

import (
	"time"
)

type TRCalendar struct{}

// Create a new US calendar
func NewTRCalendar() TRCalendar {
	return TRCalendar{}
}

// Finds the TR name for a day of the week.
// Note that time.Weekday starts at 0 with Sunday, not Monday.
func (tc TRCalendar) DayName(day time.Weekday) string {
	return []string{"pazar", "pazartesi", "salı", "çarşamba", "perşembe", "cuma", "cumartesi"}[int(day)]
}

// Finds the english name for a day of the week.
// Note that time.Weekday starts at 0 with Sunday, not Monday.
func (tc TRCalendar) DayNameInEnglish(day time.Weekday) string {
	return day.String()
}

// Finds the TR name for a given month
func (tc TRCalendar) MonthName(month time.Month) string {
	return []string{"ocak", "şubat", "mart", "nisan", "mayıs", "haziran", "temmuz", "ağustos", "eylül", "ekim", "kasım", "aralık"}[int(month)-1]
}

// Finds the english name for a given month
func (tc TRCalendar) MonthNameInEnglish(month time.Month) string {
	return month.String()
}

// Checks if a given date is a "red day" (public holiday) in the TR calendar.
// Returns true/false, a description and true/false for if it's a flag day.
func (tc TRCalendar) RedDay(date time.Time) (bool, string, bool) {

	// Source: https://en.wikipedia.org/wiki/Public_holidays_in_Turkey

	var (
		desc string
		flag bool
	)

	// Sundays
	if date.Weekday() == time.Sunday {
		desc = "Pazar"
	}

	// New Year's Day
	if atMD(date, 1, 1) {
		desc = "Yılbaşı"
		flag = true
	}

	// National sovereignty and children's day
	if atMD(date, 4, 23) {
		desc = "Ulusal Egemenlik ve Çocuk Bayramı"
		flag = true
	}

	// Labor and Solidarity Day
	if atMD(date, 5, 1) {
		desc = "İşçi Bayramı"
		flag = true
	}

	// Commemoration of Atatürk, Youth and Sports Day
	if atMD(date, 5, 19) {
		desc = "Atatürk'ü Anma, Gençlik ve Spor Bayramı"
		flag = true
	}

	// Democracy and National Unity Day
	if atMD(date, 7, 15) {
		desc = "Demokrasi ve Milli Birlik Günü"
		flag = true
	}

	// Victory Day
	if atMD(date, 8, 30) {
		desc = "Zafer Bayramı"
		flag = true
	}

	// Republic Day
	if atMD(date, 10, 29) {
		desc = "Cumhuriyet Bayramı"
		flag = true
	}

	/**
	 * TODO: calculation for Ramadan Feast and Sacrifice Day will be added.
	 */

	// Red days
	if desc != "" {
		// Then return
		return true, desc, flag
	}

	// Normal days
	return false, desc, false
}

// Some days are not red, but special in one way or another.
// Checks if a given date is notable. Returns true/false if the
// given date is notable, a comma separated description (in case there are more
// than one notable event that day) and true/false depending on if it's a flag
// flying day or not.
func (tc TRCalendar) NotableDay(date time.Time) (bool, string, bool) {

	var (
	//descriptions []string
	//flag         bool
	)

	// Since days may overlap, "flaggdager" must come first for the flag
	// flying days to be correct.

	// --- Flag days ---

	// --- Non-flag days ---

	// No notable events
	return false, "", false
}

// Checks if a given date is in a notable time range (summer holidays, for instance)
func (tc TRCalendar) NotablePeriod(date time.Time) (bool, string) {
	// TODO:
	// summer/winter/spring/autumn time
	// etc
	return false, ""
}

// An ordinary day
func (tc TRCalendar) NormalDay() string {
	return "Sıradan"
}
