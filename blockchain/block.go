package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/go-ready-blockchain/blockchain-go-core/logger"
)

type Block struct {
	Hash         []byte
	StudentData  []byte
	Signature    []byte
	Company      []byte
	Verification []byte
	PrevHash     []byte
	Nonce        int
}

func CreateBlock(data []byte, signature []byte, company []byte, verification []byte, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, signature, company, verification, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	logger.WriteToFile("Created Block for PoW")
	return block
}

func InitFirstBlock() *Block {
	logger.WriteToFile("Creating Genesis Block for Blockchain")
	return CreateBlock([]byte("Genesis Block"), []byte(""), []byte(""), []byte(""), []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	logger.WriteToFile("Serialising the given Block")
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	Handle(err)

	logger.WriteToFile("Deserialising the given Block")
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
