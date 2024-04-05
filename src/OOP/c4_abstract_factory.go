/**
*@description This program implements the abstract factory, creating
* instances of senders messages.
**/

package main

import (
	"errors"
	"fmt"
)

type INotificationFactory interface {
	SendNotification(message string)
	GetSender() ISender
}
type ISender interface {
	GeSenderMehtod() string
	GetSenderChannel() string
}
type SMSNotitification struct {
}

func (SMSNotitification) SendNotification(message string) {
	fmt.Printf("%s\n", message)
}

func (SMSNotitification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GeSenderMehtod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct{}

func (EmailNotification) SendNotification(message string) {
	fmt.Printf("%s\n", message)
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (EmailNotificationSender) GeSenderMehtod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Anymail"
}

type Employee struct {
	notificationType string
	name             string
}

func (employee *Employee) getNotificationFactory() (INotificationFactory, error) {
	switch {
	case employee.notificationType == "Email":
		return &EmailNotification{}, nil
	case employee.notificationType == "SMS":
		return &SMSNotitification{}, nil
	default:
		return nil, errors.New("Does not exists the notification ")
	}

}
func NewEmplouee(name, notificationType string) *Employee {
	return &Employee{
		name:             name,
		notificationType: notificationType,
	}
}
func sendNotification(f INotificationFactory, message string) {
	fmt.Println("Sending notification via: ", f.GetSender().GeSenderMehtod())
	f.SendNotification(fmt.Sprintf("Dear %s, you were hired by our company ", message))
}

func main() {
	var employees = []*Employee{
		NewEmplouee("John", "Email"),
		NewEmplouee("Elen", "SMS"),
	}
	for _, employee := range employees {
		notification, err := employee.getNotificationFactory()
		if err != nil {
			panic(err)
		}
		sendNotification(notification, employee.name)
	}
}
