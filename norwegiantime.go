package norwegiantime

import (
	"errors"
	"fmt"
	"strings"
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

// Gauss's method for finding easter day for a given year
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

// Returns the easter day (Første påskedag) for a given year
func EasterDay(year int) time.Time {
	month, day := easterDaySpencerJones(year)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// Checks if the given time is easter day (Første påskedag)
func isEaster(date time.Time) bool {
	eastermonth, easterday := easterDaySpencerJones(date.Year())
	return (date.Month() == time.Month(eastermonth)) && (date.Day() == easterday)
}

// Checks if the given time is at the given month and day
func atMD(date time.Time, month, day int) bool {
	return (date.Month() == time.Month(month)) && (date.Day() == day)
}

// Checks if the two given times are at the same months and days
func atDate(t, when time.Time) bool {
	return (t.Month() == when.Month()) && (t.Day() == when.Day())
}

// Checks if a day is at easterday +- a few days
func atEasterPlus(date time.Time, days int) bool {
	year := date.Year()
	eastermonth, easterday := easterDaySpencerJones(year)
	easter := time.Date(year, time.Month(eastermonth), easterday, 0, 0, 0, 0, time.UTC)
	when := easter.AddDate(0, 0, days)
	return atDate(date, when)
}

// Return the number of sundays from day t, +- a few days
func sundaysInPeriod(date time.Time, days int) int {
	sundayCounter := 0
	when := date
	if days < 0 {
		for i := days; i <= 0; i++ {
			when = date.AddDate(0, 0, i)
			if when.Weekday() == time.Sunday {
				sundayCounter++
			}
		}
	} else {
		for i := 0; i <= days; i++ {
			when = date.AddDate(0, 0, i)
			if when.Weekday() == time.Sunday {
				sundayCounter++
			}
		}
	}
	return sundayCounter
}

// Find a preceeding sunday
func searchBackwardsForSunday(date time.Time) (time.Time, error) {
	// Start with the day before the given date
	current := date.AddDate(0, 0, -1)

	// Stay within the same year
	for current.Year() == date.Year() {
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
	return date, errors.New("Could not find an earlier Sunday!")
}

// Get the week number, from 1 to 53
func WeekNum(date time.Time) int {
	_, weeknum := date.ISOWeek()
	return weeknum
}

// Checks if the given date is at Palmesøndag (the Sunday before easter)
func atPalmSunday(date time.Time) bool {
	easter := EasterDay(date.Year())
	palmSunday, err := searchBackwardsForSunday(easter)
	if err != nil {
		// This should not happen, there should always be a sunday before easter for any given year
		return false
	}
	return atDate(date, palmSunday)
}

// Find the Nth sunday of a given year and month
func nthSundayOfMonth(date time.Time, n int) (time.Time, error) {

	sundaycounter := 0

	// Start at the first day in the given month
	current := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same month
	for current.Month() == date.Month() {

		// Is it a Sunday?
		if current.Weekday() == time.Sunday {
			sundaycounter++
		}

		// Is it the Nth sunday?
		if sundaycounter == n {
			return current, nil
		}

		// If it's a sunday, advance almost one week forward
		if current.Weekday() == time.Sunday {
			current = current.AddDate(0, 0, 7)
			continue
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	return date, errors.New(fmt.Sprintf("Could not find the %dth Sunday in %s", n, date.Month().String()))

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

// Dates that are not red, but not completely ordinary either. Some days may
// overlap, in which case a comma separated (", ") list will be returned.
func NotableDay(date time.Time) (bool, string, bool) {

	// TODO: Caching

	// Source: http://www.timeanddate.no/kalender/merkedag-innhold
	// Source: http://no.wikipedia.org/wiki/Norges_offisielle_flaggdager

	var (
		descriptions []string
		flag         bool
	)

	// Since days may overlap, "flaggdager" must come first for the flag
	// flying days to be correct.

	// --- Flag days ---

	// Frigjøringsdagen
	if atMD(date, 5, 8) {
		// Frigjøringsdag 1945
		descriptions = append(descriptions, "Frigjøringsdagen")
		flag = true
	}

	// Samefolkets dag
	if atMD(date, 2, 6) {
		descriptions = append(descriptions, "Samefolkets dag")
		flag = true
	}

	// 21 januar, H.K.H. Prinsesse Ingrid Alexandras fødselsdag
	if atMD(date, 1, 21) {
		descriptions = append(descriptions, "H.K.H. Prinsesse Ingrid Alexandras fødselsdag")
		flag = true
	}

	// 21 februar, H.M. Kong Harald Vs fødselsdag
	if atMD(date, 2, 21) {
		descriptions = append(descriptions, "H.M. Kong Harald Vs fødselsdag")
		flag = true
	}

	// 7 juni, unionsoppløsningen med Sverige i 1905
	if atMD(date, 6, 7) {
		descriptions = append(descriptions, "Unionsoppløsningen med Sverige i 1905")
		flag = true
	}

	// 4 juli, H.M. Dronning Sonjas fødselsdag
	if atMD(date, 7, 4) {
		descriptions = append(descriptions, "H.M. Dronning Sonjas fødselsdag")
		flag = true
	}

	// 20 juli, H.K.H. Kronprins Haakon Magnus' fødselsdag
	if atMD(date, 7, 20) {
		descriptions = append(descriptions, "H.K.H. Kronprins Haakon Magnus' fødselsdag")
		flag = true
	}

	// 29. juli, Olsokdagen
	if atMD(date, 7, 29) {
		descriptions = append(descriptions, "Olsokdagen")
		flag = true
	}

	// 19. aug, H.K.H. Kronprinsesse Mette Marits fødselsdag
	if atMD(date, 8, 19) {
		descriptions = append(descriptions, "H.K.H. Kronprinsesse Mette Marits fødselsdag")
		flag = true
	}

	// 9. sept hvert 4. år, 2013, 2017 osv, Stortingsvalg-dagen
	if (date.Year()-1)%4 == 0 {
		if atMD(date, 9, 9) {
			descriptions = append(descriptions, "Stortingsvalg-dagen")
			flag = true
		}
	}

	// --- Non-flag days ---

	// Askeonsdag (fasten begynner)
	if atEasterPlus(date, -46) {
		descriptions = append(descriptions, "Askeonsdag")
	}

	// Påskeaften (fasten slutter)
	if atEasterPlus(date, -1) {
		descriptions = append(descriptions, "Påskeaften")
	}

	// Fastelavnssøndag (første dag i fastelavn, festen før fasten)
	// Source: http://www.aktivioslo.no/hvaskjer/fastelavn/
	if atEasterPlus(date, -49) {
		descriptions = append(descriptions, "Fastelavnsøndag")
	}

	// Blåmandag (andre dag i fastelavn)
	if atEasterPlus(date, -48) {
		descriptions = append(descriptions, "Blåmandag")
	}

	// Feitetirsdag (tredje og siste dag i fastelavn, også kjent som Mardi Gras)
	if atEasterPlus(date, -47) {
		descriptions = append(descriptions, "Feitetirsdag (Mardi Gras)")
	}

	// Sankthansaften
	if atMD(date, 6, 23) {
		descriptions = append(descriptions, "Sankthansaften")
	}

	// Nyttårsaften
	if atMD(date, 12, 31) {
		descriptions = append(descriptions, "Nyttårsaften")
	}

	// Morsdag
	if atMorsdag(date) {
		descriptions = append(descriptions, "Morsdag")
	}

	// Farsdag
	if atFarsdag(date) {
		descriptions = append(descriptions, "Farsdag")
	}

	// Valentinsdagen
	if atMD(date, 2, 14) {
		descriptions = append(descriptions, "Valentinsdagen")
	}

	// Allehelgensaften (Halloween)
	if atMD(date, 10, 31) {
		descriptions = append(descriptions, "Allehelgensaften (Halloween)")
	}

	// Allehelgensdag
	if atMD(date, 11, 1) {
		descriptions = append(descriptions, "Allehelgensdag")
	}

	// TODO: Implement these:
	//
	// bevegelig, vårjevndøgn
	// siste søndag i mars, sommertid, klokka stilles 1 time frem
	// bevegelig, sommersolverv
	// uke 28, 29 og 30, fellesferie
	// bevegelig, høstjevndøgn
	// siste søndag i oktober, vintertid, klokka stilles 1 time tilbake
	// bevegelig, vintersolverv

	// If there are notable events, return them as a string
	if len(descriptions) > 0 {
		return true, strings.Join(descriptions, ", "), flag
	}

	// No notable events
	return false, "", false
}

// Returns the third boolean argument given a time.Time value and
// a function that takes a time.Time and returns a bool, a string and a bool
func thirdBool(date time.Time, fn func(time.Time) (bool, string, bool)) bool {
	_, _, b := fn(date)
	return b
}

// Describe what type of day a given date is, in Norwegian.
func Describe(date time.Time) string {
	fulldesc := ""
	if red, desc, _ := RedDay(date); red {
		fulldesc = desc
	}
	if notable, desc, _ := NotableDay(date); notable {
		if fulldesc == "" {
			fulldesc = desc
		} else {
			fulldesc += ", " + desc
		}
	}
	if fulldesc != "" {
		return fulldesc
	}
	// Vanlig hverdag
	return "Hverdag"
}

// Checks if a given date is a flying flag day or not
func FlagDay(date time.Time) bool {
	return thirdBool(date, RedDay) || thirdBool(date, NotableDay)
}

// Checks if a given date is a "red day" in the Norwegian calendar.
// Returns true/false, a description and true/false for if it's a flag day.
// The dates will never overlap.
// Includes the 24th of December, even though only half the day is off.
func RedDay(date time.Time) (bool, string, bool) {

	// TODO: Caching

	// Source: http://www.diskusjon.no/index.php?showtopic=1084239
	// Source: http://no.wikipedia.org/wiki/Helligdager_i_Norge

	var (
		desc string
		flag bool
	)

	// Sundays
	if date.Weekday() == 0 {
		desc = "Søndag"
	}

	// Første nyttårsdag, 1. januar
	if atMD(date, 1, 1) {
		desc = "Første nyttårsdag"
		flag = true
	}

	// Palmesøndag
	if atPalmSunday(date) {
		desc = "Palmesøndag"
	}

	// Skjærtorsdag (easter - 3d)
	if atEasterPlus(date, -3) {
		desc = "Skjærtorsdag"
	}

	// Langfredag (easter - 2d)
	if atEasterPlus(date, -2) {
		desc = "Langfredag"
	}

	// Første påskedag
	if atEasterPlus(date, 0) {
		desc = "Første påskedag"
		flag = true
	}

	// Andre påskedag (easter + 1d)
	if atEasterPlus(date, 1) {
		desc = "Andre påskedag"
	}

	// Arbeidernes internasjonale kampdag, 1. mai
	if atMD(date, 5, 1) {
		// Arbeiderbevegelsens dag
		desc = "Arbeidernes internasjonale kampdag"
		flag = true
	}

	// Grunnlovsdagen, 17. mai
	if atMD(date, 5, 17) {
		// Norges grunnlovsdag/nasjonaldagen
		desc = "Grunnlovsdagen"
		flag = true
	}

	// Kristi himmelfartsdag (40. påskedag: easter + 39d)
	if atEasterPlus(date, 39) {
		desc = "Kristi himmelfartsdag"
	}

	// Første pinsedag (50. påskedag: easter + 49d)
	if atEasterPlus(date, 49) {
		desc = "Første pinsedag"
		flag = true
	}

	// Andre pinsedag (51. påskedag: easter + 50d)
	if atEasterPlus(date, 50) {
		desc = "Andre pinsedag"
	}

	// Julaften, halv-rød dag!
	if atMD(date, 12, 24) {
		desc = "Julaften (halv dag)"
	}

	// Første juledag (25. desember)
	if atMD(date, 12, 25) {
		desc = "Første juledag"
		flag = true
	}

	// Andre juledag (26. desember)
	if atMD(date, 12, 26) {
		desc = "Andre juledag"
	}

	// Red days
	if desc != "" {
		return true, desc, flag
	}

	// Normal days
	return false, desc, false
}

// Finds the norwegian name for a day of the week.
// Note that time.Weekday starts at 0 with Sunday, not Monday.
func DayName(day time.Weekday) string {
	days := []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}
	return days[int(day)]
}
