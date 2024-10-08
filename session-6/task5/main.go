package main

import "fmt"

func main() {
	var langs = []string{"go", "java", "go", "python", "go", "java"} //Instructor code
	var result = map[string]int{}
	for _, value := range langs {
		result[value] += 1
	}
	fmt.Println("Input:", langs)
	fmt.Println("Output:", result)

	//	counts := countWords("go", "java", "go", "python", "go", "java")  //My code
	//	fmt.Println("Output", counts)
	//	for key, value := range counts {
	//		fmt.Printf("%s : %d\n", key, value)
	//	}
	//}
	//
	//func countWords(words ...string) map[string]int {
	//	fmt.Println("Input:", words)
	//	counts := make(map[string]int)
	//	receivedStrings := strings.Join(words, ", ")
	//	for _, subString := range words {
	//		matchedWord := strings.Count(receivedStrings, subString)
	//		counts[subString] = matchedWord
	//	}
	//	return counts
}
