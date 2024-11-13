package mocking

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockNotify struct {
	mock.Mock
}

func (m *MockNotify) Send(message string) error {

	args := m.Called(message)

	err := args.Error(0)

	if err != nil {
		return ErrorMessageNotSent
	} else {
		return nil
	}
}

func TestSendMessage(t *testing.T) {

	testObj := new(MockNotify)

	testObj.On("Send", "Message 1").Return(nil)
	testObj.On("Send", "Message 2").Return(ErrorMessageNotSent)
	t.Run("Message should be sent", func(t *testing.T) {
		err := NotifyUser(testObj, "Message 1")
		assert.Nil(t, err)
	})
	t.Run("Should be error on sending message", func(t *testing.T) {
		err := NotifyUser(testObj, "Message 2")
		assert.Equal(t, ErrorMessageNotSent, err)
	})

	testObj.AssertExpectations(t)

}
