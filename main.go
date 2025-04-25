package main

import (
	"caln/calendar"
	"fmt"
)

func main() {
	data, err := calendar.ParseJson("data/2082.json")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	calendar.PrintNepaliCalendar(data)

}
