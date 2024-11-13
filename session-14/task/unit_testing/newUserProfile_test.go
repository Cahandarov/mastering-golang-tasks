package unit_testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserProfile(t *testing.T) {
	t.Run("Testing true version", func(t *testing.T) {
		result := NewUserProfile("Emil", 16)
		assert.Equal(t, result.IsMinor, true, "IsMinor should be true")
	})
	t.Run("Testing false version", func(t *testing.T) {
		result := NewUserProfile("Shamkhal", 42)
		assert.Equal(t, result.IsMinor, false, "IsMinor should be false")
	})
	t.Run("Testing when age is 18", func(t *testing.T) {
		result := NewUserProfile("Hasan", 18)
		assert.Equal(t, result.IsMinor, false, "IsMinor should be false again")
	})
}
