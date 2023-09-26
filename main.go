package main

// "github.com/pengzuhao/golang-utils/chinesecalendar"
// "github.com/pengzuhao/golang-utils/sshremotecmd"

import (
	"fmt"
	"golang-utils/chinesecalendar"
	"golang-utils/parseyaml"
	"golang-utils/sshremotecmd"
)

var remoteAddr, userName, passwd, cmd = "192.168.131.129", "root", "1", "ls /opt"
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
	fmt.Printf("Val: %v, Type: %T\n", resWithOutOutput, resWithOutOutput)

	// 3
	var yamlFile = "parseyaml.yaml"
	newData := &parseyaml.YamlStruct{
		EipAddress:   "1.1.1.1",
		AllocationId: "aa",
		RecordId:     "bb",
	}
	err := parseyaml.WriteYaml(yamlFile, newData)
	if err != nil {
		return
	}
	reads, err := parseyaml.ReadYaml(yamlFile)
	if err != nil {
		return
	}
	fmt.Println(reads.AllocationId, reads.EipAddress, reads.RecordId)
}
