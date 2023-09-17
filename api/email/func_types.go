package email

type MailUserExistsFunc func(MailService, string) (EmailService, error)
