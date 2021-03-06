package blockchain

import (
	"fmt"

	"github.com/go-ready-blockchain/blockchain-go-core/student"
	"github.com/go-ready-blockchain/blockchain-go-core/utils"
)

type BlockChainIterator struct {
	CurrentHash []byte
}

func InitBlockChain() {
	//logger.WriteToFile("Intialising the blockchain")
	firstblock := InitFirstBlock()

	//logger.WriteToFile("Storing the blockchain in database")
	utils.StoreInBLOCKCHAIN(firstblock.Hash, firstblock.Serialize())

	utils.StoreLastHash(firstblock.Hash)

	//logger.WriteToFile("First Block Added to blockchain")
	fmt.Println("First Block Added to BlockChain!")

}

func AddBlock(block *Block) {

	//logger.WriteToFile("Fetching last hash")
	block.PrevHash = utils.GetLastHash()

	//logger.WriteToFile("Stroing in Blockchain DB")
	utils.StoreInBLOCKCHAIN(block.Hash, block.Serialize())

	utils.StoreLastHash(block.Hash)

	//logger.WriteToFile("Block Added to Blockchain")
	fmt.Println("Added to BlockChain")

}

func Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{utils.GetLastHash()}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	var encodedBlock []byte
	currenthash := iter.CurrentHash
	encodedBlock = utils.GetFromBLOCKCHAIN(currenthash)

	block = Deserialize(encodedBlock)

	iter.CurrentHash = block.PrevHash

	return block
}

//TODO : Change this to get all blocks in one shot

func InitBlockInBuffer(name string, Company string) {

	lastHash := utils.GetLastHash()

	v := InitVerification()
	verification := EncodeToBytes(v)

	var studentdata []byte = []byte("StudentData")
	studentdata = utils.GetStudentData(name)

	newBlock := CreateBlock(studentdata, []byte{}, []byte(Company), []byte{}, lastHash)
	fmt.Println("New Block created!")
	newBlock.StudentData = student.EncryptStudentData("AcademicDept", studentdata)

	newBlock.Signature = student.GenerateStudentSignature(name, studentdata)

	newBlock.Verification = verification

	PutBlockIntoBuffer(newBlock, name, Company)
	fmt.Println("Block added to Buffer")
	//logger.WriteToFile("Block Added to Buffer")

}

func GetBlockFromBuffer(name string, company string) *Block {
	//logger.WriteToFile("Fetching block from buffer")
	namecompany := name + "/" + company
	var encodedBlock []byte = []byte("BufferBlock")
	encodedBlock = utils.FetchBlockFromBuffer(namecompany)
	var block *Block = Deserialize(encodedBlock)

	return block

}

func PutBlockIntoBuffer(block *Block, name string, company string) {
	//logger.WriteToFile("Storing temporary block into the buffer")
	namecompany := name + "/" + company

	var encodedblock []byte = block.Serialize()

	utils.StoreInBuffer(encodedblock, namecompany)

}
