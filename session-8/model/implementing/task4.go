package implementing

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	message string
}

type SMSNotifier struct {
	message string
}

func (e *EmailNotifier) Notify(message string) {
	e.message = message
	fmt.Printf("Sending email notification: %s\n", e.message)
}
func (s *SMSNotifier) Notify(message string) {
	s.message = message
	fmt.Printf("Sending SMS notification: %s\n", s.message)
}

func Task4() {
	fmt.Println("Task 4  ****************")
	var email Notifier = &EmailNotifier{}
	var sms Notifier = &SMSNotifier{}
	email.Notify("Your item has shipped")
	sms.Notify("Your item has shipped")
}
