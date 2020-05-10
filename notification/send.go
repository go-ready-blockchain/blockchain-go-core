package notification

import "fmt"

func sendNotification(companyName string, Backlog string, StarOffer string, Branch string, Gender string, CgpaCond string, Cgpa string, Perc10thCond string, Perc10th string, Perc12thCond string, Perc12th string) {
	emailitems := ApplyFilter(Backlog, StarOffer, Branch, Gender, CgpaCond, Cgpa, Perc10thCond, Perc10th, Perc12thCond, Perc12th)
	fmt.Println(emailitems)

	for name, email := range emailitems {
		fmt.Println(name, email)
		//sendEmail("Request for Student Data", name, companyName, "link: ", email)
	}
}

func test_main() {
	Backlog := "true"
	StarOffer := ""
	Branch := ""
	Gender := ""
	CgpaCond := ""
	//CgpaCond := "GreaterThan"
	Cgpa := "5"
	//Perc10thCond := ""
	Perc10thCond := "GreaterThan"
	Perc10th := "30"
	//Perc12thCond := ""
	Perc12thCond := "GreaterThan"
	Perc12th := "90"

	sendNotification("JPMC", Backlog, StarOffer, Branch, Gender, CgpaCond, Cgpa, Perc10thCond, Perc10th, Perc12thCond, Perc12th)

}
