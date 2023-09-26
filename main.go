package main

import (
	"fmt"

	"github.com/pengzuhao/golang-utils/chinesecalendar"
	"github.com/pengzuhao/golang-utils/sshremotecmd"
	// "golang-utils/chinesecalendar"
	// "golang-utils/sshremotecmd"
)

func main() {
	// 1
	dateStr := "2023-09-29"
	property, isHoliday := chinesecalendar.Bmcx(dateStr)
	fmt.Println(property, isHoliday)

	// 2
	remoteAddr, userName, passwd, cmd := "192.168.131.129", "root", "1", "ls /root"
	port := 22
	sshremotecmd.CmdWithOutput(remoteAddr, userName, passwd, cmd, port)
}
