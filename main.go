package main

import (
	"fmt"

	"github.com/pengzuhao/golang-utils/chinesecalendar"
	// "golang-utils/chinesecalendar"
)

func main() {
	dateStr := "2023-09-29"
	property, isHoliday := chinesecalendar.Bmcx(dateStr)
	fmt.Println(property, isHoliday)
}
