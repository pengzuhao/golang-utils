package main

import (
	"fmt"

	"golang-utils/gochinesecalendar"
)

func main() {
	dateStr := "2023-09-29"
	property, isHoliday := gochinesecalendar.IsYesterdayWorkDay(dateStr)
	fmt.Println(property, isHoliday)
}
