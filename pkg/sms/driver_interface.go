package sms

type Driver interface {
	Send(phone string, message Message, config map[string]string) bool
}
