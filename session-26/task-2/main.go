package main

import "fmt"

/*
Factory Method Pattern
*/
type Product struct {
}
type Service struct {
}
type Custom struct {
}

func (o Product) Process(orderID string) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("enter order")
	} else {
		return orderID, nil
	}
}
func (o Service) Process(orderID string) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("enter order")
	} else {
		return orderID, nil
	}
}
func (o Custom) Process(orderID string) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("enter order")
	} else {
		return orderID, nil
	}
}

type Processor interface {
	Process(string) (string, error)
}

func GetOrderProcessor(orderType string) (Processor, error) {
	if orderType == "product" {
		return &Product{}, nil
	} else if orderType == "service" {
		return &Service{}, nil
	} else if orderType == "custom" {
		return &Custom{}, nil
	} else {
		return nil, fmt.Errorf("invalid order type")
	}
}

func main() {
	fmt.Println("Enter order type (product, service, custom):")
	var orderType string
	fmt.Scan(&orderType)

	fmt.Println("Enter order ID:")
	var orderID string
	fmt.Scan(&orderID)

	processor, err := GetOrderProcessor(orderType)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	result, err := processor.Process(orderID)
	if err != nil {
		fmt.Printf("Error processing order: %s\n", err)
		return
	}

	fmt.Println(result)
}
