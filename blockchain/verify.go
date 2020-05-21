package blockchain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-ready-blockchain/blockchain-go-core/logger"
	"github.com/go-ready-blockchain/blockchain-go-core/security"
	"github.com/go-ready-blockchain/blockchain-go-core/utils"
)

type Verification struct {
	Verified      string               `json:"Verified"`
	Timestamps    map[string]time.Time `json:"Timestamps"`
	Verifications map[string]string    `json:"Verifications"`
}

func AcademicDeptVerification(name string, company string) bool {

	logger.WriteToFile("Academic Dept Verification Initiated")
	block := GetBlockFromBuffer(name, company)

	studentdata, dflag := security.DecryptMessage(block.StudentData, security.GetUserFromDB("AcademicDept").PrivateKey)
	if dflag == false {
		logger.WriteToFile("Decrytion of Message Failed")
		fmt.Println("Decrytion of Message Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	sflag := security.VerifyPSSSignature(security.GetPublicKeyFromDB(name), block.Signature, studentdata)
	if sflag == false {
		logger.WriteToFile("Signature Verification Failed, Authentication Failed")
		fmt.Println("Signature Verification Failed, Authentication Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	block.StudentData = studentdata
	v, vflag := ValidationByAcademicDept(DecodeToStruct(block.Verification), block)
	if vflag == false {
		logger.WriteToFile("Validation By Academic Dept Failed")
		fmt.Println("Validation By Academic Dept Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	//fmt.Println("Verified: ", v.Verified, "\n", "Verifications: ", v.Verifications, "\n", "Timestamps: ", v.Timestamps, "\n") //Print Struct

	block.Verification = EncodeToBytes(v)
	block.StudentData = security.EncryptMessage(studentdata, security.GetPublicKeyFromDB("PlacementDept"))
	block.Signature = security.PSSSignature(studentdata, security.GetUserFromDB("AcademicDept").PrivateKey)
	logger.WriteToFile("Academic Dept Verification Completed")
	PutBlockIntoBuffer(block, name, company)
	return true
}

func PlacementDeptVerification(name string, company string) bool {

	logger.WriteToFile("Placement Dept Verification Initiated")
	block := GetBlockFromBuffer(name, company)

	studentdata, dflag := security.DecryptMessage(block.StudentData, security.GetUserFromDB("PlacementDept").PrivateKey)
	if dflag == false {
		logger.WriteToFile("Decrytion of Message Failed")
		fmt.Println("Decrytion of Message Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	sflag := security.VerifyPSSSignature(security.GetPublicKeyFromDB("AcademicDept"), block.Signature, studentdata)
	if sflag == false {
		logger.WriteToFile("Signature Verification Failed, Authentication Failed")
		fmt.Println("Signature Verification Failed, Authentication Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	block.StudentData = studentdata
	v, vflag := ValidationByPlacementDept(DecodeToStruct(block.Verification), block)
	if vflag == false {
		logger.WriteToFile("Validation By Academic Dept Failed")
		fmt.Println("Validation By Placement Dept Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	//fmt.Println("Verified: ", v.Verified, "\n", "Verifications: ", v.Verifications, "\n", "Timestamps: ", v.Timestamps, "\n") //Print Struct

	block.Verification = EncodeToBytes(v)
	block.StudentData = security.EncryptMessage(studentdata, security.GetPublicKeyFromDB(company))
	block.Signature = security.PSSSignature(studentdata, security.GetUserFromDB("PlacementDept").PrivateKey)

	PutBlockIntoBuffer(block, name, company)

	//Add block to blockchain as a transaction

	AddBlock(block)
	logger.WriteToFile("Placement Dept Verification Completed")
	fmt.Println("Validation successfully completed.\nCompany can retrieve the data")
	return true
}

func InitVerification() *Verification {
	logger.WriteToFile("Verification Structure Initiated")
	v := &Verification{"", make(map[string]time.Time), make(map[string]string)}

	v.Verified = "Not Done Yet"

	v.Verifications["Academic Dept"] = "Not Done Yet"
	v.Verifications["Placement Dept"] = "Not Done Yet"

	v.Timestamps["Created At"] = time.Now()
	v.Timestamps["Academic Dept"] = time.Time{}  //zero value
	v.Timestamps["Placement Dept"] = time.Time{} //zero value

	return v
}

func CheckIfVerifiedByAll(v *Verification) bool {
	logger.WriteToFile("Verification Status Check")
	if v.Verified == "True" {
		return true
	} else if v.Verified == "Not Done Yet" {
		return false
	} else {
		fmt.Println("Error")
	}
	return false
}

func CheckIfVerifiedByAcademicDept(v *Verification) bool {
	logger.WriteToFile("Check if Verified By Academic Dept")
	if v.Verifications["Academic Dept"] == "True" {
		return true
	} else if v.Verifications["Academic Dept"] == "Not Done Yet" {
		return false
	} else {
		fmt.Println("Error")
	}
	return false
}

func ValidationByAcademicDept(v *Verification, block *Block) (*Verification, bool) {
	logger.WriteToFile("Validation By Academic Dept")
	if CheckIfVerifiedByAll(v) {
		fmt.Println("Already Verified")
		return v, true
	}

	//TODO: validateBlockchain()
	pow := NewProof(block)
	vflag := pow.Validate()
	if vflag == false {
		logger.WriteToFile("Validation By Academic Dept of Proof of Work Failed ")
		fmt.Println("Academic Dept Validation of Proof Of Work Failed")
		return v, false
	}
	logger.WriteToFile("Academic Dept Successfully completed Validation of Proof Of Work!")
	fmt.Println("Academic Dept Successfully completed Validation of Proof Of Work!")

	tflag := ProofOfElapsedTime(v.Timestamps["Created At"])
	if tflag == false {
		logger.WriteToFile("Proof Of Elapsed Time failed")
		fmt.Println("Proof Of Elapsed Time failed")
		return v, false
	}

	v.Verifications["Academic Dept"] = "True"
	v.Timestamps["Academic Dept"] = time.Now()

	logger.WriteToFile("Academic Dept Successfully completed Validation")
	return v, true
}

func ValidationByPlacementDept(v *Verification, block *Block) (*Verification, bool) {
	logger.WriteToFile("Validation By Placement Dept")
	if CheckIfVerifiedByAll(v) {
		fmt.Println("Already Verified")
		return v, true
	}

	if CheckIfVerifiedByAcademicDept(v) == false {
		logger.WriteToFile("Verification Not Yet Done by Academic Dept")
		fmt.Println("Verification Not Yet Done by Academic Dept")
		return v, false
	}

	//TODO: validateBlockchain()
	pow := NewProof(block)
	vflag := pow.Validate()
	if vflag == false {
		logger.WriteToFile("Placement Dept Validation of Proof Of Work Failed")
		fmt.Println("Placement Dept Validation of Proof Of Work Failed")
		return v, false
	}
	logger.WriteToFile("Placement Dept Successfully completed Validation Proof Of Work!")
	fmt.Println("Placement Dept Successfully completed Validation Proof Of Work!")

	tflag := ProofOfElapsedTime(v.Timestamps["Created At"])
	if tflag == false {
		logger.WriteToFile("Proof Of Elapsed Time failed")
		fmt.Println("Proof Of Elapsed Time failed")
		return v, false
	}

	v.Verifications["Placement Dept"] = "True"
	v.Timestamps["Placement Dept"] = time.Now()
	v.Verified = "True"

	logger.WriteToFile("Placement Dept Successfully completed Validation")
	return v, true
}

func EncodeToBytes(v *Verification) []byte {
	vbytes, _ := json.Marshal(v)
	return vbytes
}

func DecodeToStruct(vbytes []byte) *Verification {
	result := &Verification{}
	json.Unmarshal(vbytes, &result)
	return result
}

func ProofOfElapsedTime(creation time.Time) bool {
	logger.WriteToFile("Proof Of Elapsed Time Initiated")
	limit := creation.Add(24 * time.Hour) //24 Hr Limit
	now := time.Now()
	return now.After(creation) && now.Before(limit)
}

// func TestVerify() {

// 	//Initialize the Verification struct for a new student
// 	v := InitVerification()

// 	fmt.Println("Verified: ", v.Verified, "\n", "Verifications: ", v.Verifications, "\n", "Timestamps: ", v.Timestamps, "\n") //Print Struct

// 	//Function Call For Academic Dept to Verify
// 	v, flag := ValidationByAcademicDept(v)

// 	fmt.Println("Verified: ", v.Verified, "\n", "Verifications: ", v.Verifications, "\n", "Timestamps: ", v.Timestamps, "\n") //Print Struct

// 	//Function Call For Placement Dept to Verify
// 	v, flag = ValidationByPlacementDept(v)

// 	fmt.Println("Verified: ", v.Verified, "\n", "Verifications: ", v.Verifications, "\n", "Timestamps: ", v.Timestamps, "\n") //Print Struct

// 	fmt.Println(flag)
// 	//Encode it to Bytes
// 	b := EncodeToBytes(v)
// 	//fmt.Println(b)

// 	//Store it in the block
// 	block := &TestBlock{b}

// 	//fmt.Println(block)

// 	//Fetch it from Block and Decode it
// 	nb := block.Verify
// 	nv := DecodeToStruct(nb)

// 	fmt.Println("Verified: ", nv.Verified, "\n", nv.Timestamps, "\n", nv.Verifications) //Print Struct

// }

// func main() {

// }
