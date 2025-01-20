package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Tapşırıq:
Bir sistem yaradın ki, bir JSON faylından konfiqurasiya məlumatlarını oxusun və müəyyən şərtlərə uyğun olaraq
fayldakı features melumatlarina "featuresC" məlumatlara elave edib yenidən yazsın.
*/

type Config struct {
	AppName         string   `json:"appName"`
	Version         float32  `json:"version"`
	EnabledFeatures []string `json:"features"`
}

func main() {
	// JSON faylını oxuyuruq

	// Mellim bu yuxaridakilar lazim deil uje, ReadFile kifayet edir byte kodlari lde etmek ucun

	buf, err := os.ReadFile("sessions/29/students/shamkhal/task-1/config.json")
	if err != nil {
		fmt.Println("error happened while reading file")
	}
	// JSON məlumatlarını struct-a çeviririk
	var config Config
	errJ := json.Unmarshal([]byte(buf), &config)
	if errJ != nil {
		fmt.Println("error happened while unmarshalling file")
	}
	fmt.Println(config)

	// Yeni "featureC" xüsusiyyətini EnabledFeatures-ə əlavə edirik
	config.EnabledFeatures = append(config.EnabledFeatures, "featureD")
	fmt.Println(&config.EnabledFeatures)

	/*
		Modifikasiya olunmuş məlumatları yazırıq
		İndentasiya üçün MarshalIndent istifadə olunur: https://pkg.go.dev/encoding/json#MarshalIndent

	*/
	mar, errJs := json.MarshalIndent(config, "", "  ")
	if errJs != nil {
		fmt.Println("error happened on marshalIntend")
	}

	/*
		Yeni məlumatları fayla yazırıq
		0644: https://pkg.go.dev/os#FileMode
		WriteFile: https://pkg.go.dev/os#WriteFile
	*/

	errmIn := os.WriteFile("sessions/29/students/shamkhal/task-1/config.json", mar, 0644)
	if errmIn != nil {
		fmt.Println("error on writing file")
	}

	fmt.Println("Configuration updated successfully.")
}
