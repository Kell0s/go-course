package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	appVersion = "0.1.0"
	appName    = "notifier"
)

type Notification struct {
	Recipient string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Channel   string `json:"channel"`
	isUrgent  bool   `json:"urgent"`
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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		App     string `json:"app"`
		Version string `json:"version"`
		Status  string `json:"status"`
	}{
		App:     appName,
		Version: appVersion,
		Status:  "available",
	}

	js, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	n := Notification{
		Recipient: "example@gmail.com",
		Subject:   "Deply done",
		Body:      "v1.2.3 deplyed",
		Channel:   "email",
	}

	w.Header().Set("Content-type", "application/json")

	err := n.Validate()
	if err != nil {
		errResp := struct {
			Error string `json: "error"`
		}{
			Error: err.Error(),
		}
		js, _ := json.Marshal(errResp)
		w.Header().Set("Content-type", "application/json")
		w.Write(js)
		return
	}

	js, _ := json.Marshal(n)
	w.Write(js)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/notification", notificationHandler)

	fmt.Println("starting server on: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
