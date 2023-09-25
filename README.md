# golang-utils

## 1、gochinesecalendar: 节假日判断
usage:
import (
	"fmt"

	"github.com/pengzuhao/golang-utils/gochinesecalendar"
)

func main() {
	dateStr := "2023-09-29"
	property, isHoliday := gochinesecalendar.IsYesterdayWorkDay(dateStr)
	fmt.Println(property, isHoliday)
}
return: (string, bool)
1、workday, false           // 工作日
2、wnrl_riqi_ban, false     // 调休
3、wnrl_riqi_mo, true       // 周末
4、wnrl_riqi_xiu, true      // 休假
