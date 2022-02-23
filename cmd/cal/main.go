package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xyproto/calendar"
	"github.com/xyproto/env"
	"github.com/xyproto/textoutput"
)

// centerText will pad a string with spaces up to the given width,
// so that the string is centered.
func centerText(s string, width int) string {
	l := len(s)
	if l < width {
		after := (width - l) / 2
		before := (width - l) - after
		return strings.Repeat(" ", before) + s + strings.Repeat(" ", after)
	}
	return s
}

func monthToNorwegianString(month time.Month) string {
	var nc calendar.NorwegianCalendar
	return nc.MonthName(month)
}

func centeredMonthYearString(year int, month time.Month, langEnv string, width int) string {
	var s string
	// TODO: Use functions from the calendar package instead
	if langEnv == "nb_NO" {
		s = monthToNorwegianString(month)
	} else {
		s = month.String()
	}
	s += fmt.Sprintf(" %d", year)
	return centerText(s, width)
}

func weekdayHeader(mondayFirst bool, langEnv string) string {
	if langEnv == "nb_NO" {
		// TODO: Use nc.DayName() and extract the first two letters of each
		if mondayFirst {
			return "ma ti on to fr lø sø"
		} else {
			return "sø ma ti on to fr lø"
		}
	} else {
		if mondayFirst {
			return "Mo Tu We Th Fr Sa Su"
		} else {
			return "Su Mo Tu We Th Fr Sa"
		}
	}
}

func weekdayPosition(mondayFirst bool, current time.Time) int {
	weekdayPosition := int(current.Weekday())
	if mondayFirst {
		weekdayPosition--
		if weekdayPosition < 0 {
			weekdayPosition = 6
		}
	}
	return weekdayPosition
}

func MonthCalendar(cal *calendar.Calendar, langEnv string, givenYear int, givenMonth time.Month) string {
	mondayFirst := false
	if langEnv == "nb_NO" {
		mondayFirst = true
	}

	var sb strings.Builder

	now := time.Now()
	current := time.Date(givenYear, givenMonth, 1, 0, 0, 0, 0, now.Location())

	// Month and year, centered
	sb.WriteString(centeredMonthYearString(givenYear, givenMonth, langEnv, 20) + "\n")

	// The shortened names of the week days
	sb.WriteString(weekdayHeader(mondayFirst, langEnv) + "\n")

	// Indentation before the first day of the month
	sb.WriteString(strings.Repeat(" ", weekdayPosition(mondayFirst, current)*3))

	// Output all the numbers of the month, with linebreaks at appropriate locations
	//foundToday := false
	//wroteArrow := false
	for current.Month() == givenMonth {
		if current.Day() == now.Day() {
			sb.WriteString(fmt.Sprintf("<lightblue>%2d</lightblue> ", current.Day()))
			//foundToday = true
		} else if current.Weekday() == time.Sunday || calendar.RedDay(*cal, current) {
			// TODO: Collect descriptions, then print them below
			sb.WriteString(fmt.Sprintf("<red>%2d</red> ", current.Day()))
		} else if calendar.FlagDay(*cal, current) {
			// TODO: Collect descriptions, then print them below
			sb.WriteString(fmt.Sprintf("<yellow>%2d</yellow> ", current.Day()))
		} else {
			sb.WriteString(fmt.Sprintf("%2d ", current.Day()))
		}
		current = current.AddDate(0, 0, 1)

		if mondayFirst {
			if current.Weekday() == time.Monday {
				//if foundToday && !wroteArrow {
				//sb.WriteString("<white><-</white>")
				//wroteArrow = true
				//}
				sb.WriteString("\n")
			}
		} else {
			if current.Weekday() == time.Sunday {
				//if foundToday && !wroteArrow {
				//sb.WriteString("<white><-</white>")
				//wroteArrow = true
				//}
				sb.WriteString("\n")
			}
		}
	}

	// Final newline
	sb.WriteString("\n")

	return sb.String()
}

func main() {
	now := time.Now()

	currentYear := now.Year()
	currentMonth := now.Month()

	// Check if the first given argument is a number. If yes, use that as the current year.
	if len(os.Args) > 2 {
		if m, err := strconv.Atoi(os.Args[1]); err == nil { // success
			currentMonth = time.Month(m)
		}
		if y, err := strconv.Atoi(os.Args[2]); err == nil { // success
			currentYear = y
		}
	} else if len(os.Args) > 1 {
		if y, err := strconv.Atoi(os.Args[1]); err == nil { // success
			currentYear = y
		}
	}

	o := textoutput.New()

	langEnv := strings.TrimSuffix(env.Str("LANG"), ".UTF-8")
	cal, err := calendar.NewCalendar(langEnv, true)
	if err != nil || env.Str("LC_ALL") == "C" {
		langEnv = "en_US" // default to en_US
	}
	mc := MonthCalendar(&cal, langEnv, currentYear, currentMonth)

	o.Println(mc)
}
