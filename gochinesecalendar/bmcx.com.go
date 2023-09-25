package gochinesecalendar

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

func IsYesterdayWorkDay(dateStr string) (property string, isHoliday bool) {
	// layout := "2006-01-02"
	// dateFormat, err := time.ParseInLocation(layout, dateStr, time.Local)
	// if err != nil {
	// 	panic(err)
	// }
	client := &http.Client{}
	client.Timeout = time.Second
	res, err := client.Get("https://wannianrili.bmcx.com/" + dateStr + "__wannianrili/")
	res.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	res.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	res.Header.Set("Accept-Encoding", "gzip, deflate, br")
	res.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	res.Header.Set("Cookie:", "c_y_g_j=36; Hm_lvt_bd706f26d2267b54fd3543ceaea48e96=1692344207; __gads=ID=de73f9cfc58d87e9-22b122abe6e2003f:T=1692344208:RT=1692345564:S=ALNI_MYOegKhyMhfd3Ko7iBuby0DaF6Bdg; __gpi=UID=00000c2e3c2105b5:T=1692344208:RT=1692345564:S=ALNI_MZWtO0xcsIifnQ7_TG567AO4Km3cA; Hm_lpvt_bd706f26d2267b54fd3543ceaea48e96=1692345567")
	res.Header.Set("Pragma", "no-cache")
	res.Header.Set("Upgrade-Insecure-Requests", "1")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	reStr := fmt.Sprintf(`.*class="(.*)" href="/%v__wannianrili/".*`, dateStr)
	re := regexp.MustCompile(reStr)
	bodySli := re.FindStringSubmatch(string(body))
	if len(bodySli) == 0 {
		property = "workday"
		return property, false
	}
	property = bodySli[len(bodySli)-1]
	switch property {
	case "wnrl_riqi_mo":
		isHoliday = true
	case "wnrl_riqi_xiu":
		isHoliday = true
	case "wnrl_riqi_ban":
		isHoliday = false
	default:
		isHoliday = false
	}
	fmt.Println(isHoliday)
	return property, isHoliday
}
