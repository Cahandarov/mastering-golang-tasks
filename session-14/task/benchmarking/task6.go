package benchmarking

import (
	"strings"
)

func ConcatenateWithPlus(strs []string) string {
	var strDefault string
	for _, str := range strs {
		strDefault += str
	}
	return strDefault
}

func ConcatenateWithJoin(strs []string) string {
	return strings.Join(strs, "")
}

//func Task6() {
//	fmt.Println("Task-6    **************8")
//	fmt.Println(ConcatenateWithPlus([]string{"I", "hope", "one", "day", "i", "will", "be", "developer", "in", "Pasha"}))
//	fmt.Println(ConcatenateWithJoin([]string{"I", "hope", "one", "day", "i", "will", "be", "developer", "in", "Pasha"}))
//}
