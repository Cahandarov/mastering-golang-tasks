package mocking

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockGetDataBase struct {
	mock.Mock
}

func (m *MockGetDataBase) GetUserByID(id int) (User, error) {

	args := m.Called(id)

	if err := args.Error(1); err != nil {
		return User{}, ErrorUserNotFound
	}

	user, ok := args.Get(0).(User)
	if !ok {
		return User{}, ErrorUserNotFound
	}
	return user, nil
}

func TestGetDataBase(t *testing.T) {

	testObj := new(MockGetDataBase)

	testObj.On("GetUserByID", 1).Return(User{Id: 1, Name: "John"}, nil)

	testObj.On("GetUserByID", 3).Return(User{}, ErrorUserNotFound)

	t.Run("Testing getting user by ID", func(t *testing.T) {
		user, err := GetUser(testObj, 1)
		assert.Nil(t, err)
		assert.Equal(t, "John", user.Name, "they should be equal")
	})

	t.Run("Testing getting user by wrong ID", func(t *testing.T) {
		_, err := GetUser(testObj, 3)
		assert.Equal(t, ErrorUserNotFound, err, "should be error")
	})

	testObj.AssertExpectations(t)
}
