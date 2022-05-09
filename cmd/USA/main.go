package main

import (
	"fmt"
	"time"

	"github.com/xyproto/kal"
)

// List all the days of a given month
func datebonanza(cal kal.Calendar, year int, month time.Month) {
	fmt.Println(month.String(), year)
	fmt.Println("====================")

	current := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same month
	for current.Month() == month {

		if red, desc, flag := cal.RedDay(current); red {
			fmt.Printf("%s is red: %s (flag: %v)\n", current.String()[:10], desc, flag)
		}

		if notable, desc, flag := cal.NotableDay(current); notable {
			fmt.Printf("%s is notable: %s (flag: %v)\n", current.String()[:10], desc, flag)
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}
	fmt.Println()
}

// List notable days
func notable(cal kal.Calendar, year int) {
	current := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same year
	for current.Year() == year {

		if notable, desc, flag := cal.NotableDay(current); notable {
			fmt.Printf("%s %d is at %s (flag: %v)\n", desc, year, current.String()[:10], flag)
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	fmt.Println()
}

// List flag days
func flag(cal kal.Calendar, year int) {
	current := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same year
	for current.Year() == year {

		if flag := kal.FlagDay(cal, current); flag {
			fmt.Printf("%s (%s) is a flag flying day\n", kal.Describe(cal, current), current.String()[:10])
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	fmt.Println()
}

func main() {
	cal, err := kal.NewCalendar("en_US", true)
	if err != nil {
		panic("Could not create a US calendar!")
	}

	//year := time.Now().Year()
	year := 2016

	// When is easter this year?
	easter := kal.EasterDay(year)
	fmt.Printf("Easter %d is at %s\n", year, easter.String()[:10])

	// Show some info for March this year
	//datebonanza(year, time.Month(3))

	// Show some info for March 2000
	//datebonanza(2000, time.Month(3))

	// Show some info for the current month
	datebonanza(cal, year, time.Month(time.Now().Month()))

	notable(cal, year)
	notable(cal, year+1)

	flag(cal, year)

	fmt.Println(cal.DayName(1))
	fmt.Println(cal.MonthName(time.Month(12)))
}
