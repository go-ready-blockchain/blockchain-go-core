package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-ready-blockchain/blockchain-go-core/Init"
	"github.com/go-ready-blockchain/blockchain-go-core/blockchain"
	"github.com/go-ready-blockchain/blockchain-go-core/company"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("createBlockChain \tTo Create a new Block Chain")
	fmt.Println("student -usn USN -branch BRANCH -name NAME -gender GENDER -dob DOB -perc10th PERC10TH -perc12th PERC12TH -cgpa CGPA -backlog BACKLOG -email EMAIL -mobile MOBILE -staroffer STAROFFER\tTo Add a New Student")
	fmt.Println("company -name NAME \tAddCompany")
	fmt.Println("request -company COMPANY -student USN \tCompany requests for Student's Data")
	fmt.Println("verify-AcademicDept -student USN \tAcademicDept Verifies Student's data")
	fmt.Println("verify-PlacementDept -student USN \tPlacementDept Verifies Student's data")
	fmt.Println("companyRetrieveData -student USN \tCompany retrieves Student's data")
	fmt.Println("print - Prints the blocks in the chain")
}

func createBlockChain() {
	Init.InitializeBlockChain()
	fmt.Println("BlockChain Initialized!")
}

func addStudent(usn string, branch string, name string, gender string, dob string, perc10th string, perc12th string, cgpa string, backlog bool, email string, mobile string, staroffer bool) {
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
		Usn       string `json:"Usn"`
		Branch    string `json:"Branch"`
		Name      string `json:"Name"`
		Gender    string `json:"Gender"`
		Dob       string `json:"Dob"`
		Perc10th  string `json:"Perc10th"`
		Perc12th  string `json:"Perc12th"`
		Cgpa      string `json:"Cgpa"`
		Backlog   bool   `json:"Backlog"`
		Email     string `json:"Email"`
		Mobile    string `json:"Mobile"`
		StarOffer bool   `json:"StarOffer"`
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
		message = "Company retrieved the data!"
	} else {
		message = "Company failed to retrieve the data!"
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

func main() {
	port := "5000"
	http.HandleFunc("/createBlockChain", callcreateBlockChain)
	http.HandleFunc("/student", calladdStudent)
	http.HandleFunc("/company", calladdCompany)
	http.HandleFunc("/request", callrequestBlock)
	http.HandleFunc("/verify-AcademicDept", callverificationByAcademicDept)
	http.HandleFunc("/verify-PlacementDept", callverificationByPlacementDept)
	http.HandleFunc("/companyRetrieveData", callcompanyRetrieveData)
	http.HandleFunc("/print", callprintChain)
	http.HandleFunc("/usage", callprintUsage)
	fmt.Printf("Server listening on localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
