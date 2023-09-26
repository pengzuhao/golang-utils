package main

import (
	"fmt"
	"golang-utils/chinesecalendar"
	"golang-utils/sshremotecmd"
	// "github.com/pengzuhao/golang-utils/chinesecalendar"
	// "github.com/pengzuhao/golang-utils/sshremotecmd"
	// "golang-utils/chinesecalendar"
	// "golang-utils/sshremotecmd"
)

var remoteAddr, userName, passwd, cmd = "192.168.131.129", "root", "1", "ls /root"
var port = 22

func main() {
	// 1
	dateStr := "2023-09-29"
	property, isHoliday, _ := chinesecalendar.Bmcx(dateStr)
	fmt.Println(property, isHoliday)

	// 2
	resWithOutput, _ := sshremotecmd.CmdWithOutput(remoteAddr, userName, passwd, cmd, port)
	fmt.Println(resWithOutput)
	resWithOutOutput, _ := sshremotecmd.CmdWithOutOutput(remoteAddr, userName, passwd, cmd, port)
	fmt.Printf("Val: %v, Type: %T", resWithOutOutput, resWithOutOutput)
}
