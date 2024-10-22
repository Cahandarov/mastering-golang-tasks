package implementing

import "fmt"

type Notifier interface {
	notify(message string)
}
type EmailNotifier struct {
	message string
}
type SMSNotifier struct {
	message string
}

func (e *EmailNotifier) notify(message string) {
	e.message = message
	fmt.Printf("Sending email notification: %s\n", e.message)
}
func (s *SMSNotifier) notify(message string) {
	s.message = message
	fmt.Printf("Sending SMS notification: %s\n", s.message)
}

func Task4() {
	fmt.Println("Task 4  ****************")
	email := &EmailNotifier{}
	sms := &SMSNotifier{}
	email.notify("Your item has shipped")
	sms.notify("Your item has shipped")
}
