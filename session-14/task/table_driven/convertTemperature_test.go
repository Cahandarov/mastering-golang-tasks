package table_driven

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCasesTemp = []struct {
	value       float64
	scale       string
	expected    float64
	expectedErr error
	message     string
}{
	{28, "Celsius", 82, nil, "Converting from Celsius to Fahrenheit"},
	{33.8, "Fahrenheit", 1, nil, "Converting from Fahrenheit to Celsius"},
	{56, "Fahrenheit", 13, nil, "Converting from Fahrenheit to Celsius"},
	{28, "Celsiusss", 0, ErrorWrongInput, "Error occurred due to wrong input"},
	{28, "Fahreinheittr", 0, ErrorWrongInput, "Error occurred due to wrong input"},
}

func TestConvertTemperature(t *testing.T) {
	for _, tc := range testCasesTemp {
		t.Run(tc.message, func(t *testing.T) {
			result, err := ConvertTemperature(tc.value, tc.scale)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}
