package assctins

import (
	"fmt"
	"session-20/cn_to_dbase"
)

type Customer struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:255;not null"`
	Orders []Order
}

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Number     uint   `gorm:"not null"`
	Name       string `gorm:"size:255;not null"`
}

func migrateCustomerAndOrder() {
	err := cn_to_dbase.DB.AutoMigrate(&Customer{}, &Order{})
	if err != nil {
		fmt.Println(err)
	}
}

func createCustomer() {
	customer := Customer{Name: "Hasan", Orders: []Order{
		{Number: 1234456, Name: "Laptop"},
		{Number: 12784456, Name: "Iphone"},
		{Number: 15674456, Name: "Adaptor"},
	}}
	cn_to_dbase.DB.Create(&customer)
}

func readCustomers() {
	customers := []Customer{}
	result := cn_to_dbase.DB.Preload("Orders").Find(&customers)
	if result.Error == nil {

		for _, customer := range customers {
			output := ""
			output += fmt.Sprintf("Customer: %s ", customer.Name)
			for _, order := range customer.Orders {
				output += fmt.Sprintf("Order: %s ", order.Name)
			}
			fmt.Println(output)
		}

	}
}

func Task5() {
	migrateCustomerAndOrder()
	createCustomer()
	readCustomers()
}
