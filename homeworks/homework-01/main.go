package main

import "fmt"

type Notification struct {
	Recipient string
	Subject   string
	Body      string
	Channel   string
	IsUrgent  bool
}

func (n Notification) Format() string {
	return fmt.Sprintf(
		"to: %s | subject: %s | body: %s | channel: %s | urgent: %v",
		n.Recipient,
		n.Subject,
		n.Body,
		n.Channel,
		n.IsUrgent,
	)
}

func (n Notification) Send() (string, string) {
	message := n.Format()
	status := n.Urgent()
	return message, status
}

func (n Notification) Validate() error {
	if n.Recipient == "" {
		return fmt.Errorf("recipient is required")
	}

	if n.Subject == "" {
		return fmt.Errorf("subject is required")
	}

	return nil
}

func (n Notification) Urgent() string {
	if !n.IsUrgent {
		return "queued"
	} else {
		return "queued-priority"
	}
}

func runCase(name string, n Notification) {
	fmt.Println("CASE:", name)

	err := n.Validate()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("sent: no")
		fmt.Println()
		return
	}

	message, status := n.Send()
	fmt.Println(message)
	fmt.Println("status:", status)
	fmt.Println("sent: yes")
	fmt.Println()
}

func main() {
	regular := Notification{
		Recipient: "user@example.com",
		Subject:   "Deploy done",
		Body:      "v1.2.3 deployed",
		Channel:   "email",
		IsUrgent:  false,
	}

	urgent := Notification{
		Recipient: "ops@example.com",
		Subject:   "Database unavailable",
		Body:      "connection refused",
		Channel:   "telegram",
		IsUrgent:  true,
	}

	emptyRecipient := Notification{
		Subject:  "Deploy done",
		Body:     "v1.2.3 deployed",
		Channel:  "email",
		IsUrgent: false,
	}

	emptySubject := Notification{
		Recipient: "user@example.com",
		Body:      "v1.2.3 deployed",
		Channel:   "email",
		IsUrgent:  true,
	}

	bothEmpty := Notification{
		Body:     "v1.2.3 deployed",
		Channel:  "email",
		IsUrgent: true,
	}

	runCase("regular", regular)
	runCase("urgent", urgent)
	runCase("empty recipient", emptyRecipient)
	runCase("empty subject", emptySubject)
	runCase("both empty", bothEmpty)
}
