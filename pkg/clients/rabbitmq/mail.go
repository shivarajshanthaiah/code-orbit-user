package rabbitmq

func SendMail(email, message, subject string) error{
	msgs := Messages{
		Email:    email,
		Messages: message,
		Subject:  subject,
	}
	return PublishNotifications(msgs)
}