package cmd_interface

import (
	"flag"
	"fmt"
	"os"
)

func handleHello() {
	fs := flag.NewFlagSet("hello", flag.ExitOnError)

	err := fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}

	if fs.NArg() < 1 {
		fmt.Println("Please enter name also")
	} else {
		fmt.Printf("Hello, %s!\n", fs.Arg(0))
	}
}

func handleVersion() {
	fs := flag.NewFlagSet("version", flag.ExitOnError)

	err := fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}

	fmt.Println("Version: 1.0.0")
}
func handleHelp() {
	fs := flag.NewFlagSet("help", flag.ExitOnError)

	err := fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}
	fmt.Println("Usage: ")
	fmt.Println("hello <name> - Greet the user.")
	fmt.Println("version      - Display the version.")
	fmt.Println("help         - Display this help.")

}
func Task1() {
	fmt.Println("Task-1   ******************")

	if len(os.Args) < 2 {
		fmt.Println("command is missing")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hello":
		handleHello()
	case "version":
		handleVersion()
	case "help":
		handleHelp()
		//default:
		//	fmt.Println("Please enter valid command!")
	}

}
