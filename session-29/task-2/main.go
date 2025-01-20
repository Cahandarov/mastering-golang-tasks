package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

/*
Tapşırıq:
Bir XML faylından məlumatları oxuyun, average product price statistik göstəriciləri hesablamaq üçün məlumatları işləyin və nəticəni yeni bir XML faylına yazın.


statistics.xml:
<statistics>
  <average_price>15.166666666666666</average_price>
  <total_products>3</total_products>
</statistics>

*/

type Product struct {
	Name  string  `xml:"name"`
	Price float64 `xml:"price"`
}

type Products struct {
	XMLName  xml.Name  `xml:"products"`
	Products []Product `xml:"product"`
}
type Statistics struct {
	AveragePrice  float64 `xml:"average_price"`
	TotalProducts int     `xml:"total_products"`
}

func main() {
	// XML faylını oxuyuruq
	buf, err := os.ReadFile("sessions/29/students/shamkhal/task-2/products.xml")
	if err != nil {
		fmt.Println("error happened while reading file")
	}

	// XML məlumatlarını struct-a çeviririk

	var products Products
	errJ := xml.Unmarshal([]byte(buf), &products)
	if errJ != nil {
		fmt.Println("error happened while unmarshalling file")
	}
	fmt.Println(products)

	// Statistik hesablamalar
	var total float64
	for _, v := range products.Products {
		total = total + v.Price
	}
	average := total / float64(len(products.Products))
	fmt.Println(average)
	statisticts := Statistics{
		AveragePrice:  average,
		TotalProducts: len(products.Products),
	}
	fmt.Println(statisticts)

	// Yeni XML faylı yazırıq
	// İndentasiya üçün MarshalIndent istifadə olunur: https://pkg.go.dev/encoding/xml#MarshalIndent
	mar, errJs := xml.MarshalIndent(statisticts, "", "  ")
	if errJs != nil {
		fmt.Println("error happened on marshalIntend")
	}

	errmIn := os.WriteFile("sessions/29/students/shamkhal/task-2/statisticts.xml", mar, 0644)
	if errmIn != nil {
		fmt.Println("error on writing file")
	}

	// Yeni məlumatları fayla yazırıq
	// 0644: https://pkg.go.dev/os#FileMode
	// WriteFile: https://pkg.go.dev/os#WriteFile

	fmt.Println("Statistics XML file created successfully.")
}
