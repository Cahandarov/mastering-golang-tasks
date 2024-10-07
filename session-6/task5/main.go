package main

import (
	"fmt"
	"strings"
)

func main() {

	counts := countWords("go", "java", "go", "python", "go", "java")
	fmt.Println("Output", counts)
	for key, value := range counts {
		fmt.Printf("%s : %d\n", key, value)
	}
}

func countWords(words ...string) map[string]int {
	fmt.Println("Input:", words)
	counts := make(map[string]int)
	receivedStrings := strings.Join(words, ", ")
	for _, subString := range words {
		matchedWord := strings.Count(receivedStrings, subString)
		counts[subString] = matchedWord
	}
	return counts
}
