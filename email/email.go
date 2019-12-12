package email

import (
	"bytes"
	"os"
	"github.com/joho/godotenv"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)

	//* load dotenv file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

}

func mailer(receipients []string) {
	emailUsername := os.Getenv("email_user")
	emailPassword := os.Getenv("email_password")
	emailServer := os.Getenv("email_server")
	// emailPort := os.Getenv("email_port")

	message := &EmailMessage{
		From:    "sajioloye@gmail.com",
		To:      receipients,
		Subject: "Sign up successful",
		Body:    "Welcome to the Go Bank. Let's grow together.",
	}

	//* populate a buffer with the rendered message text from the template
	var body bytes.Buffer
	t.Execute(&body, message)

	//* set up the smtp mail client
	authCreds := &EmailCredentials{
		Username: emailUsername,
		Password: emailPassword,
		Server:   emailServer,
		Port:     2525,
	}

	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)

	//* sends the email
	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		//* the bytes from the message buffer are passed in when the message is sent
		body.Bytes())
}
