package user_inputs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task6() {
	fmt.Println("Enter text (type STOP to end):")

	var lines []string
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "STOP") {
			break
		}

		lines = append(lines, input)
	}

	fmt.Println("Captured lines in reverse order:")
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}
}
