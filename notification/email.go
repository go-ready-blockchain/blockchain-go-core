package notification

import (
	"log"
	"net/smtp"

	"github.com/go-ready-blockchain/blockchain-go-core/logger"
)

func sendEmail(subject string, studentName string, Usn string, companyName string, acceptlink string, rejectlink string, To string) {
	logger.WriteToFile("Sending Email to Student: " + studentName)
	from := "placementblk@gmail.com"
	pass := "consensusproject"
	msg := "From: " + from + "\n" +
		"To: " + To + "\n" +
		"Subject: " + subject + " \n\n" +
		"Hi " + studentName + "\n\n" +
		companyName + " is visiting your campus. You fit the eligibility criteria set by the company." + "\n\n" +
		"If you wish to register for this company,\nHit this link: " + acceptlink + "\n\n" +
		"If you wish to reject this company,\nHit this link: " + rejectlink + "\n\n" +
		"Thanks and Regards, \nPlacement Dept"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{To}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

// func main() {
// 	sendEmail("Send your details", "Ralph", "Cisco", "link: ", "pkgauravkarkal@gmail.com")
// }
