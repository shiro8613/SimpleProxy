package modules

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Bind map[string]interface{} `yaml:"bind"`
	Server []map[string]interface{} `yaml:"server"`
}

var configer []config

func LoadConfig() {
	config := config{}

	b, err := os.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	yaml.Unmarshal(b, &config)

	configer = append(configer, config)
	
}

func GetConfig() config{
	return configer[0]
}