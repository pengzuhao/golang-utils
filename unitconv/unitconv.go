package unitconv

import (
	"fmt"
	"math"
	"strconv"
)

var UnitType = map[string]map[int]string{
	"bandwidth": Bandwidth,
	"byte":      Byte,
	"meter":     Meter,
	"weight":    Weight,
	"count":     Count,
}

var Bandwidth = map[int]string{
	0: "bps",
	1: "Kbps",
	2: "Mbps",
	3: "Gbps",
	4: "Tbps",
	5: "Pbps",
}
var Byte = map[int]string{
	0: "Byte",
	1: "KB",
	2: "MB",
	3: "GB",
	4: "TB",
	5: "PB",
}
var Meter = map[int]string{
	0: "nm",
	1: "um",
	2: "mm",
	3: "m",
	4: "km",
}
var Weight = map[int]string{
	0: "mg",
	1: "g",
	2: "kg",
	3: "t",
}
var Count = map[int]string{
	0: "count",
	1: "thousand",
	2: "million",
	3: "billion",
}

var Base = map[string]int{
	"bandwidth": 1024,
	"byte":      1024,
	"meter":     1000,
	"weight":    1000,
	"count":     1000,
}

func Compute(num int, unit, unitType string) (numLast float64, unitLast string, err error) {
	var j int = 10
	_, isExist := UnitType[unitType]
	if !isExist {
		err = fmt.Errorf("unknown unittype")
		return
	}
	for k, v := range UnitType[unitType] {
		if v == unit {
			j = k
			break
		}
	}
	if j == 10 {
		err = fmt.Errorf("unknown unit")
		return
	}
	base := Base[unitType]
	if num < base {
		numLast = float64(num)
		unitLast = unit
		return
	}
	for i := 0; i < len(UnitType[unitType]); i++ {
		numNew, errFloat := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(num)/float64(math.Pow(float64(base), float64(i)))), 64)
		if errFloat != nil {
			err = fmt.Errorf(errFloat.Error())
			return
		}
		if int(numNew) < base {
			numLast = numNew
			unitLast = UnitType[unitType][i+j]
			break
		}
	}
	return
}
