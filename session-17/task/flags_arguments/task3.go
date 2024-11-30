package flags_arguments

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func handleGreet() {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	name := fs.String("name", "", "Write name")
	uprCse := fs.Bool("uppercase", false, "Make uppercase")

	err := fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}

	if *name == "" {
		fmt.Println("Name is required")
		os.Exit(1)
	} else {
		if *uprCse {
			fmt.Printf("HELLO, %s!\n", strings.ToUpper(*name))
		} else {
			fmt.Printf("Hello, %s!\n", *name)
		}

	}
}

func Task3() {
	fmt.Println("Task-3   ******************")

	if len(os.Args) < 2 {
		fmt.Println("command is missing")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "greet":
		handleGreet()
	}
}
