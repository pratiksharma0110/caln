package calendar

import (
	"fmt"
	"strconv"

	"strings"
	"time"
)

var month = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}

// this should be updated every year as monthDays are not fixed for BS
var monthDays = []int{
	31,
	31,
	32,
	31,
	31,
	31,
	30,
	29,
	30,
	29,
	30,
	30,
}

func findDay(month int, date int) (int, error) {

	if month < 1 || month > 12 {
		return -1, fmt.Errorf("invalid month")
	}
	max_days := monthDays[month-1]
	if date < 1 || date > max_days {
		return -1, fmt.Errorf("invalid date")
	}
	//baisakh 1 was on monday,
	// 1 - sund, 2- monday, 6-friday , 0-sat
	start_day := 2

	total_days := 0
	for i := 0; i < month-1; i++ {

		total_days += monthDays[i]

	}

	var day_index int = (start_day + total_days + date - 1) % 7

	return day_index, nil

}

func currentDate(data [][]map[string]any) (string, int, []any) {
	englishDate := time.Now().Format("2006/01/02")

	var nepYear string
	var nepMonthIndex int
	var nepaliDate string
	var today []any
	for index, items := range data {
		for _, m := range items {
			if ad, exists := m["ad"].(string); exists && ad == englishDate {
				nepMonthIndex = index + 1
				nepaliDate = m["bs"].(string)
				nepYear = strings.Split(nepaliDate, "-")[0]
				nepDate := strings.Split(nepaliDate, "-")[2]

				today = append(today, nepDate, m["events"].(string), m["holiday"].(bool), m["tithi"].(string))

			}

		}
	}
	return nepYear, nepMonthIndex, today

}

func getNepaliMonthName(month int) string {
	switch month {
	case 1:
		return "Baisakh"
	case 2:
		return "Jestha"
	case 3:
		return "Ashadh"
	case 4:
		return "Shrawan"
	case 5:
		return "Bhadra"
	case 6:
		return "Ashwin"
	case 7:
		return "Kartik"
	case 8:
		return "Mangsir"
	case 9:
		return "Poush"
	case 10:
		return "Magh"
	case 11:
		return "Falgun"
	case 12:
		return "Chaitra"
	default:
		return "Invalid month"
	}
}

func PrintNepaliCalendar(data [][]map[string]any) {

	nepYear, nepMonth, today := currentDate(data)

	start_day, err := findDay(1, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return

	}
	month := getNepaliMonthName(nepMonth)

	fmt.Printf("🗓️%-10s\t\t%-4v\n", strings.ToUpper(month), nepYear)

	fmt.Printf("%4s%4s%4s%4s%4s%4s \033[31m%-4s\033[0m\n", "SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT")

	spaceCount := (start_day - 1) % 7

	currentMonth := data[nepMonth-1]

	for i := 0; i < spaceCount; i++ {
		fmt.Print("    ")
	}

	for _, m := range currentMonth {

		date := m["bs"].(string)
		date = strings.Split(date, "-")[2]

		holiday := m["holiday"].(bool)

		if holiday {
			fmt.Printf("\033[31m%4v\033[0m", date)
		} else {
			fmt.Printf("%4v", date)
		}
		dateInt, err := strconv.Atoi(date)
		if err != nil {

			fmt.Println("Error converting date:", err)
			continue
		}

		if (spaceCount+dateInt)%7 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf(" Today's Info:\n")
	fmt.Printf(" Date:%v\n Events: %v\n Holiday:%v\n Tithi:%v\n", today[0], today[1], today[2], today[3])

}
