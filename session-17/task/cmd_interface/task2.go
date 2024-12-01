package cmd_interface

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func handleAdd() {
	fs := flag.NewFlagSet("add", flag.ExitOnError)

	err := fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}

	if fs.NArg() < 2 {
		fmt.Printf("Error: Invalid number of arguments for add. Expected 2, got %d.\n", fs.NArg())
	} else {
		num1, _ := strconv.ParseFloat(fs.Arg(0), 64)
		num2, _ := strconv.ParseFloat(fs.Arg(1), 64)
		fmt.Println("Result:", num1+num2)
	}
}

func handleConcat() {
	fs := flag.NewFlagSet("concat", flag.ExitOnError)

	err := fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}

	if fs.NArg() < 2 {
		fmt.Printf("Error: Invalid number of arguments for concat. Expected 2, got %d.\n", fs.NArg())
	} else {
		fmt.Println("Result:", fs.Arg(0)+fs.Arg(1))
	}
}

func Task2() {
	fmt.Println("Task-2   ******************")

	if len(os.Args) < 2 {
		fmt.Println("command is missing")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		handleAdd()
	case "concat":
		handleConcat()
	}
}
