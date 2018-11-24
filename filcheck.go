package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {

	t := time.Now()
	tf := t.Format("2006/01/02")
	fmt.Println(tf)
	config := LoadConfiguration("config.json")
	fmt.Println(config)
	info, _ := os.Stat("test.txt")
	fmt.Println(info)
}
