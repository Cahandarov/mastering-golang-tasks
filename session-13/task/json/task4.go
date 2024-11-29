package json

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AppName   string   `json:"app_name"`
	Version   string   `json:"version"`
	DebugMode bool     `json:"debug_mode"`
	Services  []string `json:"services"`
}

var config = Config{
	AppName:   "MyApp",
	Version:   "1.0",
	DebugMode: true,
	Services:  []string{"service1", "service2"},
}
var marshalledData Config

func createFile() {
	file, err := os.OpenFile("task/json/config.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error happened on creating file:", err)
	}
	defer file.Close()

	j, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Error happened on marshalling:", err)
	}

	_, err = file.Write(j)
	if err != nil {
		fmt.Println("Error happened on writing to file:", err)
	}
}

func readConfigFile() {

	data, err := os.ReadFile("task/json/config.json")
	if err != nil {
		fmt.Println("Error happened on reading file:", err)
	}
	err = json.Unmarshal(data, &marshalledData)
	if err != nil {
		fmt.Println("Error happened on unmarshalling json:", err)
	}
	fmt.Println("Configuration Details:")
	fmt.Printf("- App Name: %s\n", config.AppName)
	fmt.Printf("- Version: %s\n", config.Version)
	fmt.Printf("- Debug Mode: %v\n", config.DebugMode)
	fmt.Printf("- Services: %v\n", strings.Join(config.Services, ","))

}

func Task4() {
	fmt.Println("Task-4   ********************")

	createFile()
	readConfigFile()
}

//{"app_name":"MyApp","version":"1.0","debug_mode":true,"services":["service1","service2"]}
