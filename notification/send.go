package notification

import (
	"fmt"
)

func SendNotification(link string, companyName string, Backlog string, StarOffer string, Branch []string, Gender string, CgpaCond string, Cgpa string, Perc10thCond string, Perc10th string, Perc12thCond string, Perc12th string) bool {
	//logger.WriteToFile("Sending notification for Email")
	emailitems := ApplyFilter(Backlog, StarOffer, Branch, Gender, CgpaCond, Cgpa, Perc10thCond, Perc10th, Perc12thCond, Perc12th)

	for _, emailitem := range emailitems {
		email := emailitem.Email
		usn := emailitem.Usn
		name := emailitem.Name
		fmt.Println(email, usn, name)
		acceptlink := link + "/handlerequest?" + "approval=" + "true" + "&company=" + companyName + "&name=" + usn
		rejectlink := link + "/handlerequest?" + "approval=" + "false" + "&company=" + companyName + "&name=" + usn
		sendEmail("Request for Student Data", name, usn, companyName, acceptlink, rejectlink, email)

	}
	return true
}

// func Test_main() {
// 	Backlog := "true"
// 	StarOffer := ""
// 	Branch := ""
// 	Gender := ""
// 	CgpaCond := ""
// 	//CgpaCond := "GreaterThan"
// 	Cgpa := "5"
// 	//Perc10thCond := ""
// 	Perc10thCond := "GreaterThan"
// 	Perc10th := "30"
// 	//Perc12thCond := ""
// 	Perc12thCond := "GreaterThan"
// 	Perc12th := "90"

// 	SendNotification("link", "JPMC", Backlog, StarOffer, Branch, Gender, CgpaCond, Cgpa, Perc10thCond, Perc10th, Perc12thCond, Perc12th)

// }
