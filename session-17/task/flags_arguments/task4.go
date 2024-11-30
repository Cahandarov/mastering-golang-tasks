package flags_arguments

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func handleMath() {

	fs := flag.NewFlagSet("math", flag.ExitOnError)
	operation := fs.String("operation", "", "Math operation: add, subtract, multiply, divide")

	err := fs.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	if fs.NArg() < 2 {
		fmt.Println("Error: Two operands are required")
		os.Exit(1)
	}

	arg1, err1 := strconv.Atoi(fs.Arg(0))
	arg2, err2 := strconv.Atoi(fs.Arg(1))
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Operands must be integers")
		os.Exit(1)
	}

	switch *operation {
	case "add":
		fmt.Printf("Result: %d\n", arg1+arg2)
	case "subtract":
		fmt.Printf("Result: %d\n", arg1-arg2)
	case "multiply":
		fmt.Printf("Result: %d\n", arg1*arg2)
	case "divide":
		if arg2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
		} else {
			fmt.Printf("Result: %d\n", arg1/arg2)
		}
	default:
		fmt.Println("Error: Invalid operation. Supported operations are add, subtract, multiply, divide")
	}
}

func Task4() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Command is missing")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "math":
		handleMath()
	default:
		fmt.Println("Error: Unknown command")
	}
}
