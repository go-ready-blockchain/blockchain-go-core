package main

//ONLY FOR TESTING PURPOSES
//MANUAL PIPELINE

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-ready-blockchain/blockchain-go-core/Init"
	"github.com/go-ready-blockchain/blockchain-go-core/blockchain"
	"github.com/go-ready-blockchain/blockchain-go-core/company"
	"github.com/go-ready-blockchain/blockchain-go-core/logger"
	"github.com/go-ready-blockchain/blockchain-go-core/notification"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("Make POST request to /createBlockChain \tTo Create a new Block Chain")
	fmt.Println("Make POST request to /student \tTo Add a New Student")
	fmt.Println("Make POST request to /company \tTo Add Company")
	fmt.Println("Make POST request to /send \tTo Send Email to Eligible Students based on Eligibility Criteria")
	fmt.Println("Make GET request to /handlerequest \tHandle Request and Initiate Creation of Request Block")
	fmt.Println("Make POST request to /verify-AcademicDept \tAcademicDept Verifies Student's data")
	fmt.Println("Make POST request to /verify-PlacementDept \tPlacementDept Verifies Student's data")
	fmt.Println("Make POST request to /companyRetrieveData \tCompany retrieves Student's data")
	fmt.Println("Make POST request to /print \t Prints the blocks in the chain")
	fmt.Println("Make POST request to /request-student \t Test Direct Request to Student")
}

func createBlockChain() {
	Init.InitializeBlockChain()
	fmt.Println("BlockChain Initialized!")
}

func addStudent(usn string, branch string, name string, gender string, dob string, perc10th float32, perc12th float32, cgpa float32, backlog bool, email string, mobile string, staroffer bool) {
	Init.InitStudentNode(usn, branch, name, gender, dob, perc10th, perc12th, cgpa, backlog, email, mobile, staroffer)
	fmt.Println("Student Added!")

}
func addCompany(company string) {
	Init.InitCompanyNode(company)
	fmt.Println("Company Added!")

}

func requestBlock(name string, company string) {
	blockchain.InitBlockInBuffer(name, company)
	fmt.Println("Requested Block Initialized!")
}

func verificationByAcademicDept(name string, company string) bool {
	flag := blockchain.AcademicDeptVerification(name, company)
	if flag == true {
		fmt.Println("Verification By Academic Dept Successfully completed!")
		return true
	} else {
		fmt.Println("Verification By Academic Dept Failed!")
		return false
	}
}

func verificationByPlacementDept(name string, company string) bool {
	flag := blockchain.PlacementDeptVerification(name, company)
	if flag == true {
		fmt.Println("Verification by Placement Dept Successfully completed!")
		return true
	} else {
		fmt.Println("Verification by Placement Dept Failed!")
		return false
	}
}

func companyRetrieveData(name string, companyname string) bool {
	flag := company.RetrieveData(name, companyname)
	if flag == true {
		fmt.Println("Company retrieved the data!")
		return true
	} else {
		fmt.Println("Company failed to retrieve the data!")
		return false
	}
}

func printChain() {
	iter := blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Student Data: %x\n", block.StudentData)
		fmt.Printf("Signature: %x\n", block.Signature)
		fmt.Printf("Company: %s\n", block.Company)
		fmt.Printf("Verification: %s\n", block.Verification)
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func callcreateBlockChain(w http.ResponseWriter, r *http.Request) {

	createBlockChain()

	w.Header().Set("Content-Type", "application/json")
	message := "BlockChain Initialized!"
	w.Write([]byte(message))
}

func calladdStudent(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Usn       string  `json:"Usn"`
		Branch    string  `json:"Branch"`
		Name      string  `json:"Name"`
		Gender    string  `json:"Gender"`
		Dob       string  `json:"Dob"`
		Perc10th  float32 `json:"Perc10th"`
		Perc12th  float32 `json:"Perc12th"`
		Cgpa      float32 `json:"Cgpa"`
		Backlog   bool    `json:"Backlog"`
		Email     string  `json:"Email"`
		Mobile    string  `json:"Mobile"`
		StarOffer bool    `json:"StarOffer"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	addStudent(b.Usn, b.Branch, b.Name, b.Gender, b.Dob, b.Perc10th, b.Perc12th, b.Cgpa, b.Backlog, b.Email, b.Mobile, b.StarOffer)

	message := "Student Added!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func calladdCompany(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	addCompany(b.Company)

	message := "Company Added!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func sendNotification(w http.ResponseWriter, r *http.Request) {

	type jsonBody struct {
		Company      string   `json:"company"`
		Backlog      string   `json:"backlog"`
		StarOffer    string   `json:"starOffer"`
		Branch       []string `json:"branch"`
		Gender       string   `json:"gender"`
		CgpaCond     string   `json:"cgpaCond"`
		Cgpa         string   `json:"cgpa"`
		Perc10thCond string   `json:"perc10thCond"`
		Perc10th     string   `json:"perc10th"`
		Perc12thCond string   `json:"perc12thCond"`
		Perc12th     string   `json:"perc12th"`
	}

	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	message := ""
	flag := notification.SendNotification("http://localhost:5000", b.Company, b.Backlog, b.StarOffer, b.Branch, b.Gender, b.CgpaCond, b.Cgpa, b.Perc10thCond, b.Perc10th, b.Perc12thCond, b.Perc12th)

	if flag == true {
		message = "Notification sent successfully to Students!"
	} else {
		message = "Sending Notification to Student Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func callrequestBlock(w http.ResponseWriter, r *http.Request) {

	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	requestBlock(b.Name, b.Company)

	message := "Requested Block Initialized!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Approval bool   `json:"approval"`
		Name     string `json:"name"`
		Company  string `json:"company"`
	}

	decoder := r.URL.Query()
	approval := decoder["approval"][0]
	Approval, _ := strconv.ParseBool(approval)
	Company := decoder["company"][0]
	Name := decoder["name"][0]
	fmt.Println(Approval, Company)

	if !Approval {
		fmt.Println("Student :", Name, "Rejected Request for Data for Company: ", Company)
		w.Write([]byte(string("Student : " + Name + " Rejected Request for Data for Company: " + Company)))
		return
	}
	requestBlock(Name, Company)

	message := "Requested Block Initialized!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))

	fmt.Println("\n\nSending Notification to Academic Dept for Verification\n\n")

}

func callverificationByAcademicDept(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	message := ""
	flag := verificationByAcademicDept(b.Name, b.Company)
	if flag == true {
		message = "Verification By Academic Dept Successfully completed!"
	} else {
		message = "Verification By Academic Dept Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func callverificationByPlacementDept(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	message := ""
	flag := verificationByPlacementDept(b.Name, b.Company)
	if flag == true {
		message = "Verification by Placement Dept Successfully completed!"
	} else {
		message = "Verification by Placement Dept Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func callcompanyRetrieveData(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	message := ""
	flag := companyRetrieveData(b.Name, b.Company)
	if flag == true {
		message = "Company retrieved the data!"
	} else {
		message = "Company failed to retrieve the data!"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func callprintChain(w http.ResponseWriter, r *http.Request) {

	printChain()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Chain!!"
	w.Write([]byte(message))
}

func callprintUsage(w http.ResponseWriter, r *http.Request) {

	printUsage()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Usage!!"
	w.Write([]byte(message))
}

func TestLog() {

	name := time.Now().String()
	logger.FileName = "Test" + name + ".log"
	logger.NodeName = "All Nodes"
	logger.CreateFile()
}

func main() {
	//notification.Test_main()
	TestLog()
	port := "5000"
	http.HandleFunc("/createBlockChain", callcreateBlockChain)
	http.HandleFunc("/student", calladdStudent)
	http.HandleFunc("/company", calladdCompany)
	http.HandleFunc("/send", sendNotification)
	http.HandleFunc("/request-student", callrequestBlock)
	http.HandleFunc("/handlerequest", handlerequest)
	http.HandleFunc("/verify-AcademicDept", callverificationByAcademicDept)
	http.HandleFunc("/verify-PlacementDept", callverificationByPlacementDept)
	http.HandleFunc("/companyRetrieveData", callcompanyRetrieveData)
	http.HandleFunc("/print", callprintChain)
	http.HandleFunc("/usage", callprintUsage)
	fmt.Printf("Server listening on localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
