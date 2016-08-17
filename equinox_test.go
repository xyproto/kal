package calendar

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	// TODO: Check the values against a table
	for i := 2010; i < 2020; i++ {
		fmt.Println()
		fmt.Println(northwardEquinox(i))
		fmt.Println(northernSolstice(i))
		fmt.Println(southwardEquinox(i))
		fmt.Println(southernSolstice(i))
	}
}
