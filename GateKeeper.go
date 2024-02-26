package main

import (
	"fmt"
	"time"
)

var Message string = "Welkom bij Fonteyn Vakantieparken"
var MessageGeenToegang string = "U heeft helaas geen toegang tot het parkeerterrein"
var timeSelect int
var kenteken string
var groet string

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

func kentekencontrole(geldigKentekens []string) bool {

	for i := 0; i < len(geldigKentekens); i++ {
		if kenteken == geldigKentekens[i] {
			return true
		}
	}
	return false
}

func main() {
	TimeNow := time.Now()
	HourNow := TimeNow.Hour()

	timechange(HourNow)
	groetchange()

	fmt.Print("Type je kenteken in: ")
	fmt.Scan(&kenteken)

	geldigKentekens := []string{"AB-12-CD", "XY-34-ZZ", "GH-56-IJ", "KL-78-MN"}

	if kentekencontrole(geldigKentekens) {
		fmt.Println(groet)
	} else {
		fmt.Println(MessageGeenToegang)
	}
	fmt.Scanln()
}
