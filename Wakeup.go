package main

import (
	"gopkg.in/robfig/cron.v3"
	"log"
	"net/smtp"
	"time"
)

func send(body string) {
	from := "*****0@gmail.com"
	pass := "*****"
	to := "******@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Wake Up \n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Email sent successfuly")
}
func main() {
	c := cron.New()
	c.AddFunc("30 8 * * ?", func() { send("Rise and Shine") })
	c.Start()
	time.Sleep(24 * time.Hour)
	c.Stop()
}
