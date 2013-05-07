package main

import (
	"fmt"
	"time"

	"github.com/xyproto/norwegiantime"
)

func main() {
	month := 3
	year := time.Now().Year()
	fmt.Println(time.Month(month).String(), year)
	fmt.Println("====================")
	for day := 1; day <= 31; day++ {
		if norwegiantime.RedDay(year, month, day) {
			if day == 31 {
				fmt.Println(day, ": "+"RØD (hvis den finnes)")
			} else {
				fmt.Println(day, ": "+"RØD!")
			}
		} else {
			if day == 31 {
				fmt.Println(day, ": "+"vanlig (hvis den finnes)")
			} else {
				fmt.Println(day, ": "+"vanlig")
			}
		}
	}

	fmt.Println()
	m, d := norwegiantime.EasterDay(year)
	fmt.Println("Påsken:", d, " ", time.Month(m).String())
}
