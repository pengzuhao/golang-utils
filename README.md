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