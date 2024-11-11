package mocking

import (
	"errors"
)

type Notifier interface {
	Send(message string) error
}

var ErrorMessageNotSent = errors.New("not sent")

func NotifyUser(n Notifier, message string) error {
	err := n.Send(message)
	if err != nil {
		return ErrorMessageNotSent
	} else {
		return nil
	}
}

// func Task8() {
// 	fmt.Println("Task-8    **************8")
// }
