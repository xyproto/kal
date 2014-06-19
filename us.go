package moskus

// locale: en_US

// NOTE: Work in progress!

// Anything that's specific to the US, with the exception of some movable dates which are in movable.go
// This calendar has a corresponding locale code in the NewCalendar function in calendar.go

import (
	"time"
)

type USCalendar struct{}

// Create a new US calendar
func NewUSCalendar() USCalendar {
	return USCalendar{}
}

// Finds the US name for a day of the week.
// Note that time.Weekday starts at 0 with Sunday, not Monday.
func (nc USCalendar) DayName(day time.Weekday) string {
	return day.String()
}

// Finds the US name for a given month
func (nc USCalendar) MonthName(month time.Month) string {
	return month.String()
}

// Checks if a given date is a "red day" (public holiday) in the US calendar.
// Returns true/false, a description and true/false for if it's a flag day.
// The dates will never overlap.
func (nc USCalendar) RedDay(date time.Time) (bool, string, bool) {

	// Source: http://en.wikipedia.org/wiki/Public_holidays_in_the_United_States
	// Source: http://timpanogos.wordpress.com/flag-fly-dates/

	var (
		desc string
		flag bool
	)

	// Sundays
	if date.Weekday() == 0 {
		desc = "Sunday"
	}

	// New Year's Day
	if atMD(date, 1, 1) {
		desc = "New Year's Day"
		flag = true
	}

	// Birthday of Dr. Martin Luther King, Jr.
	if atDrMartinLutherKingJrBirthday(date) {
		desc = "Birthday of Dr. Martin Luther King, Jr."
		flag = true
	}

	// TODO: The rest

	// Elections: 2000, 2004, 2008, 2012 etc

	/*
		First January 20 following a Presidential electionInauguration DayObserved only by federal government employees in Washington, D.C., and the border counties of Maryland and Virginia to relieve congestion that occurs with this major event. Swearing-in of President of the United States and Vice President of the United States. Celebrated every fourth year. Note: Takes place on January 21 if the 20th is a Sunday (although the President is still privately inaugurated on the 20th). If Inauguration Day falls on a Saturday, the preceding Friday is not a federal holiday.

		Third Monday in FebruaryWashington's Birthday/Presidents' DayWashington's Birthday was first declared a federal holiday by an 1879 act of Congress. The Uniform Holidays Act, 1968, shifted the date of the commemoration of Washington's Birthday from February 22 to the third Monday in February (between February 15 and 21, meaning the observed holiday never falls on Washington's actual birthday). Because of this, combined with the fact that President Abraham Lincoln's birthday falls on February 12, many people now refer to this holiday as "Presidents' Day" and consider it a day honoring all American presidents. However, neither the Uniform Holidays Act nor any subsequent law changed the name of the holiday from Washington's Birthday to Presidents' Day.[8]

		Last Monday in MayMemorial DayHonors the nation's war dead from the Civil War onwards; marks the unofficial beginning of the summer season. (traditionally May 30, shifted by the Uniform Holidays Act 1968)

		July 4Independence DayHonorsCelebrates the signing of the Declaration of Independence from British rule, also called the Fourth of July. Firework celebrations are held in many cities throughout the nation.

		First Monday in SeptemberLabor DayHonorsCelebratesCelebrates the achievements of workers and the labor movement; marks the unofficial end of the summer season.

		Second Monday in OctoberColumbus DayHonorsCelebratesCelebratesHonors Christopher Columbus, traditional discoverer of the Americas. In some areas it is also a celebration of Italian culture and heritage. (traditionally October 12)

		November 11Veterans DayHonorsCelebratesCelebratesHonorsHonors all veterans of the United States armed forces. It is observed on November 11 to recall the end of World War I on that date in 1918 (major hostilities of World War I were formally ended at the 11th hour of the 11th day of the 11th month of 1918 with the German signing of the Armistice).

		Fourth Thursday in NovemberThanksgiving DayHonorsCelebratesCelebratesHonorsHonorsTraditionally celebrates the giving of thanks for the autumn harvest. Traditionally includes the sharing of a turkey dinner. Traditional start of the Christmas and holiday season.

		December 25ChristmasThe most widely celebrated holiday of the Christian year, Christmas is observed as a commemoration of the birth of Jesus of Nazareth.
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
func (nc USCalendar) NotableDay(date time.Time) (bool, string, bool) {

	// Source: http://www.timeanddate.no/kalender/merkedag-innhold
	// Source: http://no.wikipedia.org/wiki/Norges_offisielle_flaggdager

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
func (nc USCalendar) NotablePeriod(date time.Time) (bool, string) {
	// TODO:
	// uke 28, 29 og 30, fellesferie
	// jul
	// fastelavn
	// faste
	// sommer/vinter/vår/høst (legg denne først i listen)
	return false, ""
}

// An ordinary day
func (nc USCalendar) NormalDay() string {
	return "Ordinary"
}
