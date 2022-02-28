package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xyproto/calendar"
	"github.com/xyproto/env"
	"github.com/xyproto/textoutput"
	"github.com/xyproto/vt100"
)

// centerPad will pad a string with spaces up to the given width,
// so that the string is centered.
func centerPad(s string, width int) string {
	l := len(s)
	if l < width {
		after := (width - l) / 2
		before := (width - l) - after
		return strings.Repeat(" ", before) + s + strings.Repeat(" ", after)
	}
	return s
}

// rightPad will pad the right side of a string with spaces
func rightPad(s string, width int) string {
	l := len(s)
	if l < width {
		after := (width - l)
		return s + strings.Repeat(" ", after)
	}
	return s
}

func centeredMonthYearString(cal calendar.Calendar, year int, month time.Month, width int) string {
	s := cal.MonthName(month)
	s += fmt.Sprintf(" %d", year)
	return centerPad(s, width)
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

func MonthCalendar(cal *calendar.Calendar, givenYear int, givenMonth time.Month) string {

	mondayFirst := (*cal).MondayFirst()

	var descriptions, sb strings.Builder

	now := time.Now()
	current := time.Date(givenYear, givenMonth, 1, 0, 0, 0, 0, now.Location())

	// Month and year, centered
	sb.WriteString("<lightblue>" + centeredMonthYearString(*cal, givenYear, givenMonth, 20) + "</lightblue>\n")

	// The shortened names of the week days
	sb.WriteString(calendar.TwoLetterDays(*cal, (*cal).MondayFirst()) + "\n")

	// Indentation before the first day of the month
	sb.WriteString(strings.Repeat(" ", weekdayPosition(mondayFirst, current)*3))

	// Output all the numbers of the month, with linebreaks at appropriate locations
	for current.Month() == givenMonth {
		if current.Day() == now.Day() { // Today
			sb.WriteString(fmt.Sprintf(vt100.BackgroundBlue.String()+"<white>%2d</white> ", current.Day()))
		} else if isRedDay := calendar.RedDay(*cal, current); current.Weekday() == time.Sunday || isRedDay { // Red day
			// TODO: Collect descriptions, then print them below
			sb.WriteString(fmt.Sprintf("<red>%2d</red> ", current.Day()))
			if isRedDay {
				if mondayFirst {
					descriptions.WriteString(fmt.Sprintf("<red>%2d. %s</red> - %s\n", current.Day(), (*cal).MonthName(givenMonth), calendar.Describe(*cal, current)))
				} else {
					descriptions.WriteString(fmt.Sprintf("<red>%s %d</red> - %s\n", (*cal).MonthName(givenMonth), current.Day(), calendar.Describe(*cal, current)))
				}
			}
		} else if calendar.FlagDay(*cal, current) { // Flag day
			// TODO: Collect descriptions, then print them below
			sb.WriteString(fmt.Sprintf("<lightblue>%2d</lightblue> ", current.Day()))
			if mondayFirst {
				descriptions.WriteString(fmt.Sprintf("<lightblue>%2d. %s</lightblue> - %s (flaggdag)\n", current.Day(), (*cal).MonthName(givenMonth), calendar.Describe(*cal, current)))
			} else {
				descriptions.WriteString(fmt.Sprintf("<lightblue>%s %d</lightblue> - %s (flaggdag)\n", (*cal).MonthName(givenMonth), current.Day(), calendar.Describe(*cal, current)))
			}
		} else { // Ordinary day
			sb.WriteString(fmt.Sprintf("%2d ", current.Day()))
		}
		current = current.AddDate(0, 0, 1)

		if (mondayFirst && (current.Weekday() == time.Monday)) || (!mondayFirst && (current.Weekday() == time.Sunday)) {
			sb.WriteString("\n")
		}
	}

	calendarString := sb.String()
	if !strings.HasSuffix(calendarString, "\n") {
		calendarString += "\n"
	}

	if ds := descriptions.String(); len(ds) > 0 {
		return calendarString + "\n" + ds
	}

	return calendarString
}

func main() {
	now := time.Now()

	currentYear := now.Year()
	currentMonth := now.Month()

	// Check if the first given argument is a number. If yes, use that as the current year.
	if len(os.Args) > 2 {
		if m, err := strconv.Atoi(os.Args[1]); err == nil && m >= 1 && m <= 12 { // success
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

	langEnv := strings.TrimSuffix(env.Str("LC_ALL", env.Str("LANG")), ".UTF-8")
	if langEnv == "C" || env.Str("LC_ALL") == "C" {
		langEnv = "en_US" // default to en_US
	}

	cal, err := calendar.NewCalendar(langEnv, true)
	if err != nil {
		log.Fatalln("could not create a calendar with langEnv " + langEnv)
	}

	moCal := MonthCalendar(&cal, currentYear, currentMonth)

	o := textoutput.New()
	o.Print(moCal)
}
