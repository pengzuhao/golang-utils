### 声明：此仓库仅用于工作和学习，禁止用于非法攻击，非法传播。一切遵守《网络安全法》。
# 1、gochinesecalendar: 节假日判断
用法： 
```
# go get github.com/pengzuhao/golang-utils
```

```
import (
	"fmt"

	"github.com/pengzuhao/golang-utils/chinesecalendar"
)

func main() {
	dateStr := "2023-09-29"
	property, isHoliday := chinesecalendar.IsYesterdayWorkDay(dateStr)
	fmt.Println(property, isHoliday)
}
```
```
return: (string, bool)
```

返回值说明：
- workday, false           // 工作日
- wnrl_riqi_ban, false     // 调休
- wnrl_riqi_mo, true       // 周末
- wnrl_riqi_xiu, true      // 休假
##