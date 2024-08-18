package utils

import (
	"github.com/adityarizkyramadhan/emailer"
	"os"
	"strconv"
)

func MailClient() *emailer.Mailer {
	port := os.Getenv("EMAIL_PORT")
	portInt, _ := strconv.Atoi(port)
	return emailer.New(
		os.Getenv("EMAIL_NAME"),
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_HOST"),
		portInt,
	)
}
