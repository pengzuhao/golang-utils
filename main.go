package main

// "github.com/pengzuhao/golang-utils/chinesecalendar"
// "github.com/pengzuhao/golang-utils/sshremotecmd"

import (
	"fmt"
	"golang-utils/unitconv"
)

// var remoteAddr, userName, passwd, cmd = "192.168.131.129", "root", "1", "ls /opt"
// var port = 22

func main() {
	// // 1
	// dateStr := "2023-09-29"
	// property, isHoliday, _ := chinesecalendar.Bmcx(dateStr)
	// fmt.Println(property, isHoliday)

	// // 2
	// resWithOutput, _ := sshremotecmd.CmdWithOutput(remoteAddr, userName, passwd, cmd, port)
	// fmt.Println(resWithOutput)
	// resWithOutOutput, _ := sshremotecmd.CmdWithOutOutput(remoteAddr, userName, passwd, cmd, port)
	// fmt.Printf("Val: %v, Type: %T\n", resWithOutOutput, resWithOutOutput)

	// // 3
	// var yamlFile = "parseyaml.yaml"
	// newData := &parseyaml.YamlStruct{
	// 	EipAddress:   "1.1.1.1",
	// 	AllocationId: "aa",
	// 	RecordId:     "bb",
	// }
	// err := parseyaml.WriteYaml(yamlFile, newData)
	// if err != nil {
	// 	return
	// }
	// reads, err := parseyaml.ReadYaml(yamlFile)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(reads.AllocationId, reads.EipAddress, reads.RecordId)

	// // 4
	// var content = "https://www.bilibili.com/video/BV1Cx411J7pt/?share_source=copy_web&vd_source=93c783b46e7446e13f3f91a996ca06f9"
	// var fileName = "qrcode.png"
	// err = qrcode.QRCEncode(content, fileName)
	// if err != nil {
	// 	return
	// }
	// contentRead, err := qrcode.QRCDecode(fileName)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(contentRead)

	// // 5
	// var mylog = &customlog.MyLog{Logger: log.Default()}
	// mylog.Debug("db")
	// mylog.Info("if")
	// mylog.Warn("ww")
	// mylog.Error("ee")

	// // 6
	// var ipaddr, passwd, key, value, expired = "192.168.131.129:6379", "123456", "a", 1, 0 * time.Second

	// redisopera.RedisSet(ipaddr, passwd, key, value, expired)

	// res, err := redisopera.RedisGet(ipaddr, passwd, key)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(res)

	// res, err = redisopera.RedisConnPoolGet(ipaddr, passwd, key)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(res)

	// // 7
	// dir := "./"
	// watchdir.WatchDir(dir)

	// 8
	var num, unit, unitType = 156789, "KB", "byte"
	numLast, unitLast := unitconv.Compute(num, unit, unitType)
	fmt.Println(numLast, unitLast)
}
