### 声明：此仓库仅用于工作和学习，禁止用于非法攻击，非法传播。一切遵守《网络安全法》。
# 1、chinesecalendar: 节假日判断 
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
	property, isHoliday := chinesecalendar.Bmcx(dateStr)
	fmt.Println(property, isHoliday)
}
```

```
return: (string, bool, err)
```
参数说明：
- dateStr中，月份、日期小于10时十位须补0

返回值说明：
- workday, false, err           // 工作日
- wnrl_riqi_ban, false, err     // 调休
- wnrl_riqi_mo, true, err       // 周末
- wnrl_riqi_xiu, true, err      // 休假
##
# 2、sshremotecmd：ssh执行远程命令
```
import (
	"fmt"
	"github.com/pengzuhao/golang-utils/sshremotecmd"
)

var remoteAddr, userName, passwd, cmd = "192.168.131.129", "root", "1", "ls /root"
var port = 22

func main() {
	resWithOutput := sshremotecmd.CmdWithOutput(remoteAddr, userName, passwd, cmd, port)
	fmt.Println(resWithOutput)
	resWithOutOutput := sshremotecmd.CmdWithOutOutput(remoteAddr, userName, passwd, cmd, port)
	fmt.Printf("Val: %v, Type: %T", resWithOutOutput, resWithOutOutput)
}

```
返回值说明：
- resWithOutput: string, err
- resWithOutOutput: bool, err
##
# 3、yaml文件读写
```
说明：须自行根据yaml数据结构定义结构体。
例：文件 parseyaml.yaml 内容如下，
eipaddress: 1.1.1.1
allocationid: aa
recordid: bb
```
```
import (
	"fmt"
	"github.com/pengzuhao/golang-utils/parseyaml"
)

func main() {
	newData := &parseyaml.YamlStruct{
		EipAddress:   "1.1.1.1",
		AllocationId: "aa",
		RecordId:     "bb",
	}
	reads, err := parseyaml.ReadYaml()
	if err != nil {
		return
	}
	fmt.Println(reads.AllocationId, reads.EipAddress, reads.RecordId)
	err = parseyaml.WriteYaml(newData)
	if err != nil {
		return
	}
}
```
##
# 4、二维码生成与解析
```
import (
	"fmt"
	"github.com/pengzuhao/golang-utils/qrcode"
)
func main() {
	var content = "https://www.bilibili.com/video/BV1Cx411J7pt/?share_source=copy_web&vd_source=93c783b46e7446e13f3f91a996ca06f9"
	var fileName = "qrcode.png"
	err = qrcode.QRCEncode(content, fileName)
	if err != nil {
		return
	}
	contentRead, err := qrcode.QRCDecode(fileName)
	if err != nil {
		return
	}
	fmt.Println(contentRead)
}
```
##
# 5、自定义日志
- info、warn、err 日志，输出到logs目录下不同的文件中
- debug 日志，作为标准输出
```
import (
	"github.com/pengzuhao/golang-utils/customlog"
)
func main() {
	var mylog = &customlog.MyLog{Logger: log.Default()}
	mylog.Debug("db")
	mylog.Info("if")
	mylog.Warn("ww")
	mylog.Error("ee")
}
```
##
# 6、redis 
```
import (
	"github.com/pengzuhao/golang-utils/redisopera"
)

func main(){
	var ipaddr, passwd, key, value, expired = "192.168.131.129:6379", "123456", "a", 1, 0 * time.Second

	redisopera.RedisSet(ipaddr, passwd, key, value, expired)	// set

	res, err := redisopera.RedisGet(ipaddr, passwd, key)		// get
	if err != nil {
		return
	}
	fmt.Println(res)

	res, err = redisopera.RedisConnPoolGet(ipaddr, passwd, key)	// redis pool get
	if err != nil {
		return
	}
	fmt.Println(res)
}
```
##
# 7、文件监控
```
import (
	"github.com/pengzuhao/golang-utils/watchdir"
)

func main(){
	dir := "./"
	watchdir.WatchDir(dir)
}
```
## 
# 8、单位转换（4舍5入）
```
import (
	"github.com/pengzuhao/golang-utils/unitconv"
)

func main(){
	var num, unit, unitType = 156789, "KB", "byte"
	numLast, unitLast := unitconv.Compute(num, unit, unitType)
	fmt.Println(numLast, unitLast)
}
```
### 参数说明：
- UnitType == "bandwidth" // 带宽
1. bps
2. Kbps
3. Mbps
4. Gbps
5. Tbps
6. Pbps
- UnitType == byte	// 字节
1. Byte
2. KB
3. MB
4. GB
5. TB
6. PB
- UnitType == meter	// 长度
1. nm
2. um
3. mm
4. m
5. km
- UnitType == weight 	// 质量
1. mg
2. g
3. kg
4. t
- UnitType == count 	// 数量
1. count
2. thousand
3. million
4. billion
### 返回值说明
```
// 四舍五入保留2位小数float64， string
153.11 MB
```