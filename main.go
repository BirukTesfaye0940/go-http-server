package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Email:", message)
}

type SMSNotifier struct {
}

func (s SMSNotifier) Send(message string) {
	fmt.Println("SMS:", message)
}

func Notify(n Notifier, msg string) {
	n.Send(msg)
}

func main() {
	email := EmailNotifier{}
	sms := SMSNotifier{}

	Notify(email, "Welcome user")
	Notify(sms, "Your code is 1234")
}
