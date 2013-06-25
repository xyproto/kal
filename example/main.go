package main

import (
	"fmt"
	"time"

	"github.com/xyproto/moskus"
)

// List all the days of a given month
func datebonanza(cal moskus.Calendar, year int, month time.Month) {
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
func notable(cal moskus.Calendar, year int) {
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
func flag(cal moskus.Calendar, year int) {
	current := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	// As long as we are in the same year
	for current.Year() == year {

		if flag := moskus.FlagDay(cal, current); flag {
			fmt.Printf("%s (%s) is a flaggdag\n", moskus.Describe(cal, current), current.String()[:10])
		}

		// Advance to the next day
		current = current.AddDate(0, 0, 1)
	}

	fmt.Println()
}

func main() {
	cal, err := moskus.NewCalendar("nb_NO", true)
	if err != nil {
		panic("Could not create a Norwegian calendar!")
	}

	//year := time.Now().Year()
	year := 2013

	// When is easter this year?
	easter := moskus.EasterDay(year)
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
