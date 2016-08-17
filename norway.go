package calendar

// locale: nb_NO

// Anything that's specific to Norway, with the exception of some movable dates which are in movable.go
// This calendar has a corresponding locale code in the NewCalendar function in calendar.go
// Use this as a template for implementing other languages and locales

import (
	"strings"
	"time"
)

type NorwegianCalendar struct{}

// Create a new Norwegian calendar
func NewNorwegianCalendar() NorwegianCalendar {
	return NorwegianCalendar{}
}

// Finds the Norwegian name for a day of the week.
// Note that time.Weekday starts at 0 with Sunday, not Monday.
func (nc NorwegianCalendar) DayName(day time.Weekday) string {
	return []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}[int(day)]
}

// Finds the Norwegian name for a given month
func (nc NorwegianCalendar) MonthName(month time.Month) string {
	return []string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"}[int(month)-1]
}

// Checks if a given date is a "red day" (public holiday) in the Norwegian calendar.
// Returns true/false, a description and true/false for if it's a flag day.
// The dates will never overlap.
// Includes the 24th of December, even though only half the day is red.
func (nc NorwegianCalendar) RedDay(date time.Time) (bool, string, bool) {

	// Source: http://www.diskusjon.no/index.php?showtopic=1084239
	// Source: http://no.wikipedia.org/wiki/Helligdager_i_Norge

	var (
		desc string
		flag bool
	)

	// Sundays
	if date.Weekday() == time.Sunday {
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
func (nc NorwegianCalendar) NotableDay(date time.Time) (bool, string, bool) {

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

	// Vårjevndøgn
	if atNorthwardEquinox(date) {
		descriptions = append(descriptions, "Vårjevndøgn")
	}

	// Sommersolverv
	if atNorthernSolstice(date) {
		descriptions = append(descriptions, "Sommersolverv")
	}

	// Høstjevndøgn
	if atSouthwardEquinox(date) {
		descriptions = append(descriptions, "Høstjevndøgn")
	}

	// Vintersolverv
	if atSouthernSolstice(date) {
		descriptions = append(descriptions, "Vintersolverv")
	}

	// Siste søndag i mars, sommertid, klokka stilles 1 time frem
	if atSommertid(date) {
		descriptions = append(descriptions, "Sommertid (+1t)")
	}

	// Siste søndag i oktober, vintertid, klokka stilles 1 time tilbake
	if atVintertid(date) {
		descriptions = append(descriptions, "Vintertid (-1t)")
	}

	// If there are notable events, return them as a string
	if len(descriptions) > 0 {
		desc := strings.Join(descriptions, ", ")
		// Then return
		return true, desc, flag
	}

	// No notable events
	return false, "", false
}

// Checks if a given date is in a notable time range (summer holidays, for instance)
func (nc NorwegianCalendar) NotablePeriod(date time.Time) (bool, string) {
	// TODO:
	// uke 28, 29 og 30, fellesferie
	// jul
	// fastelavn
	// faste
	// sommer/vinter/vår/høst (legg denne først i listen)
	return false, ""
}

// An ordinary day
func (nc NorwegianCalendar) NormalDay() string {
	return "Hverdag"
}
