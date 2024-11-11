package table_driven

import (
	"errors"
	"strings"
)

var ErrorPalindrome = errors.New("don't use palindrome words")

func ReverseString(str string) (string, error) {
	var reversedString string
	var stringWithoutSpecialChars string
	var slice []string
	exclude := map[rune]bool{
		'@': true,
		'#': true,
		'-': true,
		'$': true,
		'%': true,
		'&': true,
		'*': true,
	}

	trimmedAndLowercased := strings.ToLower(strings.Trim(str, " "))

	for _, l := range trimmedAndLowercased {
		if !exclude[l] {
			slice = append(slice, string(l))
		}
	}
	stringWithoutSpecialChars = strings.Join(slice, "")

	for _, v := range stringWithoutSpecialChars {
		reversedString = string(v) + reversedString
	}

	if stringWithoutSpecialChars == reversedString && stringWithoutSpecialChars != "" {
		return "", ErrorPalindrome
	} else {
		return stringWithoutSpecialChars, nil
	}

}

//func Task3() {
//	fmt.Println("Task-3    **************8")
//
//	fmt.Println(ReverseString("example@ string# with- special$ chars%"))
//	fmt.Println(ReverseString(" level"))
//}
