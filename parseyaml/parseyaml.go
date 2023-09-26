package parseyaml

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlStruct struct {
	EipAddress   string `yaml:"eipaddress"`
	AllocationId string `yaml:"allocationid"`
	RecordId     string `yaml:"recordid"`
}

func ReadYaml(yamlFile string) (reads *YamlStruct, err error) {
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(data, &reads)
	return
}
func WriteYaml(yamlFile string, newData *YamlStruct) (err error) {
	data, err := yaml.Marshal(newData)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(yamlFile, data, 0777)
	return
}
