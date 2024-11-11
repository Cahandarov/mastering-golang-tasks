package table_driven

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	str         string
	expected    string
	expectedErr error
	message     string
}{
	{"level", "", ErrorPalindrome, "Should be return error"},
	{"example@ string# with- special$ chars%", "example string with special chars", nil, "This case is ok"},
	{"", "", nil, "This case is ok"},
	{" ", "", nil, "This case is ok"},
	{" 7789", "7789", nil, "This case is ok"},
	{"7789", "7789", nil, "This case is ok"},
}

func TestReverseString(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			str, err := ReverseString(testCase.str)
			if testCase.expectedErr != nil {
				assert.Equal(t, testCase.expectedErr, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, testCase.expected, str)
		})
	}
}
