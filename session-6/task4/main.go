package main

import "fmt"

var countries = map[string]string{
	"Azerbaijan": "Baku",
	"USA":        "Washington",
	"Germany":    "Berlin",
	"Japan":      "Tokyo",
}

func main() {
	findCapital("Russia")
	countries["Russia"] = "Moscow"
	countries["Albania"] = "Durres"
	for country := range countries {
		fmt.Printf("Capital of %s is %s\n", country, countries[country])
	}
}

func findCapital(country string) {
	if value, ok := countries[country]; ok {
		fmt.Printf("Capital of %s is %s\n", country, value)
	} else {
		fmt.Printf("Capital of %s is not found\n", country)
	}
}
