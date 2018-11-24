package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Config struct {
	FileLocation      string `json:"FileLocation"`
	BlankFileLocation string `json:"BlankFileLocation"`
	MaxFileSize       int64  `json:"MaxFileSize"`
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
	config := LoadConfiguration("config.json")
	fmt.Println(config.FileLocation)
	fileInfo, err := os.Stat(config.FileLocation)
	if err == nil {
		fmt.Println(fileInfo)
		fmt.Println(fileInfo.Size())
		if fileInfo.Size() > config.MaxFileSize {
			fmt.Println(fileInfo.Size())
			fmt.Println("larger than max ")
			t := time.Now()
			tf := t.Format("2006-01-02")
			fmt.Println("renaming  file ")
			os.Rename(config.FileLocation, tf+"_"+config.FileLocation)
			originalFile, err := os.Open(config.BlankFileLocation)
			defer originalFile.Close()
			if err != nil {
				print(err)
			}
			newFile, _ := os.Create(config.FileLocation)
			io.Copy(newFile, originalFile)
		} else {
			fmt.Println("file smaller than max")
			fmt.Println("exiting ")
		}
	} else {
		fmt.Println(err)
	}

}
