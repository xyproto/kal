package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xyproto/env"
	"github.com/xyproto/kal"
	"github.com/xyproto/vt"
)

// centerPad will pad a string with spaces up to the given width,
// so that the string is centered.
func centerPad(s string, width int) string {
	l := len([]rune(s))
	if l < width {
		after := (width - l) / 2
		before := (width - l) - after
		return strings.Repeat(" ", before) + s + strings.Repeat(" ", after)
	}
	return s
}

// rightPad will pad the right side of a string with spaces
func rightPad(s string, width int) string {
	l := len([]rune(s))
	if l < width {
		after := (width - l)
		return s + strings.Repeat(" ", after)
	}
	return s
}

func centeredMonthYearString(cal kal.Calendar, year int, month time.Month, width int) string {
	s := fmt.Sprintf("%s %d", cal.MonthName(month), year)
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

// MonthCalendar returns a string that is a complete overview of the given month
func MonthCalendar(cal *kal.Calendar, givenYear int, givenMonth time.Month) string {

	mondayFirst := (*cal).MondayFirst()

	var descriptions, sb strings.Builder

	now := time.Now()
	current := time.Date(givenYear, givenMonth, 1, 0, 0, 0, 0, now.Location())

	// Add the week, if this is the current month
	var weekString string
	if now.Month() == givenMonth && now.Year() == givenYear {
		_, w := now.ISOWeek()
		weekString = "w" + strconv.Itoa(w)
	}

	// Month and year, centered
	sb.WriteString("<lightblue>" + centeredMonthYearString(*cal, givenYear, givenMonth, 20-len(weekString)) + "</lightblue><darkgray>" + weekString + "</darkgray>\n")

	// The shortened names of the week days
	sb.WriteString("<white>" + kal.TwoLetterDays(*cal, (*cal).MondayFirst()) + "</white>\n")

	// Indentation before the first day of the month
	sb.WriteString(strings.Repeat(" ", weekdayPosition(mondayFirst, current)*3))

	// Output all the numbers of the month, with linebreaks at appropriate locations
	for current.Month() == givenMonth {
		isFlagDay := kal.FlagDay(*cal, current)
		if current.Day() == now.Day() && current.Month() == now.Month() && current.Year() == now.Year() { // Today
			sb.WriteString(fmt.Sprintf(vt.BackgroundBlue.String()+"<lightyellow>%2d</lightyellow> ", current.Day()))
		} else if isRedDay := kal.RedDay(*cal, current); current.Weekday() == time.Sunday || isRedDay { // Red day
			sb.WriteString(fmt.Sprintf("<red>%2d</red> ", current.Day()))
			// Collect descriptions, then print them below
			if isRedDay {
				if isFlagDay {
					if mondayFirst {
						descriptions.WriteString(fmt.Sprintf("<lightblue>%2d. %s</lightblue> - %s (flaggdag)\n", current.Day(), (*cal).MonthName(givenMonth), kal.Describe(*cal, current)))
					} else {
						descriptions.WriteString(fmt.Sprintf("<lightblue>%s %d</lightblue> - %s (flaggdag)\n", (*cal).MonthName(givenMonth), current.Day(), kal.Describe(*cal, current)))
					}
				} else {
					if mondayFirst {
						descriptions.WriteString(fmt.Sprintf("<red>%2d. %s</red> - %s\n", current.Day(), (*cal).MonthName(givenMonth), kal.Describe(*cal, current)))
					} else {
						descriptions.WriteString(fmt.Sprintf("<red>%s %d</red> - %s\n", (*cal).MonthName(givenMonth), current.Day(), kal.Describe(*cal, current)))
					}
				}
			}
		} else if isFlagDay { // Flag day
			sb.WriteString(fmt.Sprintf("<lightblue>%2d</lightblue> ", current.Day()))
			// Collect descriptions, then print them below
			if mondayFirst {
				descriptions.WriteString(fmt.Sprintf("<lightblue>%2d. %s</lightblue> - %s (flaggdag)\n", current.Day(), (*cal).MonthName(givenMonth), kal.Describe(*cal, current)))
			} else {
				descriptions.WriteString(fmt.Sprintf("<lightblue>%s %d</lightblue> - %s (flaggdag)\n", (*cal).MonthName(givenMonth), current.Day(), kal.Describe(*cal, current)))
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
		// Assume that a single argument <= 12 was intended to be a month
		if currentYear >= 1 && currentYear <= 12 {
			currentMonth = time.Month(currentYear)
			currentYear = now.Year()
		}
	}

	langEnv := strings.TrimSuffix(env.Str("LC_ALL", env.Str("LANG")), ".UTF-8")
	if langEnv == "C" || env.Str("LC_ALL") == "C" {
		langEnv = "en_US" // default to en_US
	}

	cal, err := kal.NewCalendar(langEnv, true)
	if err != nil {
		log.Fatalln("could not create a calendar using locale " + langEnv)
	}

	moCal := MonthCalendar(&cal, currentYear, currentMonth)

	vt.New().Print(moCal)
}
