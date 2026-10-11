package main

import (
	"fmt"
)

const (
	appVersion = "0.1.0"
	appName    = "notifier"
)

type Notification struct {
	Recipient string
	Subject   string
	Body      string
	Channel   string
	isUrgent  bool
}

func (n Notification) Send() (string, string) {
	message := n.Format()
	status := "queued"
	return message, status
}

func (n Notification) Format() string {
	return fmt.Sprintf("to: %s | subject %s | body: %s | channel %s",
		n.Recipient, n.Subject, n.Body, n.Channel)
}

func (n Notification) Validate() error {
	if n.Recipient == "" {
		return fmt.Errorf("получатель не указан")
	}

	return nil
}

func main() {
	n := Notification{
		Recipient: "",
		Subject:   "Deply done",
		Body:      "v1.2.3 deplyed",
		Channel:   "email",
	}

	err := n.Validate()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	message, status := n.Send()

	fmt.Println(message, status)
}
