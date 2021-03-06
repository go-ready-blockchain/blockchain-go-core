package company

import (
	"fmt"

	"github.com/go-ready-blockchain/blockchain-go-core/blockchain"
	"github.com/go-ready-blockchain/blockchain-go-core/security"
	"github.com/go-ready-blockchain/blockchain-go-core/student"
	"github.com/go-ready-blockchain/blockchain-go-core/utils"
)

func RetrieveData(name string, company string) bool {

	//logger.WriteToFile("Company Retriving Data")
	block := blockchain.GetBlockFromBuffer(name, company)

	studentdata, dflag := security.DecryptMessage(block.StudentData, security.GetUserFromDB(company).PrivateKey)
	if dflag == false {
		fmt.Println("Decrytion of Message Failed")
		//logger.WriteToFile("Decrytion of Message Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	sflag := security.VerifyPSSSignature(security.GetPublicKeyFromDB("PlacementDept"), block.Signature, studentdata)
	if sflag == false {
		fmt.Println("Signature Verification Failed, Authentication Failed")
		//logger.WriteToFile("Signature Verification Failed, Authentication Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	v := blockchain.DecodeToStruct(block.Verification)
	vflag := blockchain.CheckIfVerifiedByAll(v)
	if vflag == false {
		fmt.Println("Verification Not Yet Done. Company not allowed to retrieve Data")
		//logger.WriteToFile("Verification Not Yet Done. Company not allowed to retrieve Data")
		return false
	}
	fmt.Println("Student Data:\n")
	studentstruct := student.DecodeToStruct(studentdata)
	student.PrintStudentData(studentstruct)
	fmt.Println("\nRetrieved Data Successfully")
	//logger.WriteToFile("Retrieved Data Successfully")

	return true
}
