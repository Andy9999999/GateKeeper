package main

import (
	"fmt"
	"time"
)

func main() {

	TimeNow := time.Now()
	HourNow := TimeNow.Hour()
	var Message string = "Welkom bij Fonteyn Vakantieparken"

	var groet string

	var timeSelect int

	if HourNow < 7 || HourNow > 23 {
		timeSelect = 0
	}

	if HourNow < 12 || HourNow > 7 {
		timeSelect = 1
	}

	if HourNow < 18 || HourNow > 12 {
		timeSelect = 2
	}

	if HourNow < 23 || HourNow > 18 {
		timeSelect = 3
	}

	switch timeSelect {
	case 0:
		groet = "Sorry, de parkeerplaats is ’s nachts gesloten"

	case 1:
		groet = fmt.Sprintf("Goedemorgen, %v\n", Message)

	case 2:
		groet = fmt.Sprintf("Goedemiddag, %v\n", Message)

	case 3:
		groet = fmt.Sprintf("Goedenavond, %v\n", Message)
	}

	fmt.Print(groet)
	fmt.Scanln()
}
