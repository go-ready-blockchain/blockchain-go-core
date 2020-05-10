package Init

import (
	"fmt"

	"github.com/go-ready-blockchain/blockchain-go-core/blockchain"
	"github.com/go-ready-blockchain/blockchain-go-core/security"
	"github.com/go-ready-blockchain/blockchain-go-core/student"
	"github.com/go-ready-blockchain/blockchain-go-core/utils"
)

func InitializeBlockChain() {
	blockchain.InitBlockChain()
	InitNodes()
}

func InitNodes() {

	security.GenerateAcademicDeptKeys()

	security.GeneratePlacementDeptKeys()

}
func InitCompanyNode(company string) {
	security.GenerateCompanyKeys(company)
}

func InitStudentNode(usn string, branch string, name string, gender string, dob string, perc10th float32, perc12th float32, cgpa float32, backlog bool, email string, mobile string, staroffer bool) {

	security.GenerateStudentKeys(usn)

	stud := student.EnterStudentData(usn, branch, name, gender, dob, perc10th, perc12th, cgpa, backlog, email, mobile, staroffer)
	fmt.Println(stud)

	utils.StoreStudentData(usn, branch, name, gender, dob, perc10th, perc12th, cgpa, backlog, email, mobile, staroffer)

	//StoreStudentDataInDb(student.EncodeToBytes(stud))

}
