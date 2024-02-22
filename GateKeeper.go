package main

import (
	"fmt"
	"time"
)

var Message string = "Welkom bij Fonteyn Vakantieparken"
var groet string
var timeSelect int

func timechange(Hour int) {
	if Hour < 7 || Hour > 23 {
		timeSelect = 0
	}

	if Hour < 12 || Hour > 7 {
		timeSelect = 1
	}

	if Hour < 18 || Hour > 12 {
		timeSelect = 2
	}

	if Hour < 23 || Hour > 18 {
		timeSelect = 3
	}

}

func groetchange() {

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

}

func main() {

	TimeNow := time.Now()
	HourNow := TimeNow.Hour()

	timechange(HourNow)
	groetchange()

	fmt.Print(groet)
	fmt.Scanln()
}
