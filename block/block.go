package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Transactions string
	PreviousHash string
	Hash         string
	Nounce       int
	Timestamp    int64
}

func NewBlock(Transactions string, previousHash string) Block {
	newBlock := Block{Transactions: Transactions, PreviousHash: previousHash, Nounce: 1, Timestamp: time.Now().Unix()}
	return newBlock
}

func (b *Block) ComputeHash() string {
	sha := sha256.New()
	jsonTransaction, _ := json.Marshal(b.Transactions)
	sha.Write([]byte(string(jsonTransaction) + b.PreviousHash + strconv.Itoa(b.Nounce) + strconv.Itoa(int(b.Timestamp))))
	res := sha.Sum(nil)
	return fmt.Sprintf("%x", res)
}

func getAnswer(difficulty int) (ans string) {
	for i := 0; i < difficulty; i++ {
		ans += "0"
	}
	return
}

func (b *Block) Mine(difficulty int) {
	for {
		b.Hash = b.ComputeHash()
		if b.Hash[:difficulty] != getAnswer(difficulty) {
			b.Nounce += 1
		} else {
			break
		}
	}
}
