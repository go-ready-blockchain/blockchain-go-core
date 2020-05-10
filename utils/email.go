package utils

import (
	"log"
	"net/smtp"
)

func send(subject string, studentName string, companyName string, To string) {
	from := "placementblk@gmail.com"
	pass := "consensusproject"
	msg := "From: " + from + "\n" +
		"To: " + To + "\n" +
		"Subject: " + subject + " \n\n" +
		studentName + " " + companyName

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{To}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

// func main() {
// 	send("Send your details", "Ralph", "Cisco", "dhanushkr42@gmail.com")
// }
