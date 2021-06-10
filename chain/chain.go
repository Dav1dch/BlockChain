package chain

import (
	"blockChain/block"
	"blockChain/transaction"
	"crypto/ecdsa"
	"encoding/json"
	"strconv"
)

type Chain struct {
	Ch              []block.Block
	TransactionPool []transaction.Transaction_struct
	MinerReward     int
	Difficulty      int
}

func (c *Chain) bigBang() {
	c.Ch = make([]block.Block, 0)
	genesisBlock := block.NewBlock("我是祖先", "")
	genesisBlock.Hash = genesisBlock.ComputeHash()
	c.TransactionPool = make([]transaction.Transaction_struct, 0)
	c.MinerReward = 50
	c.Ch = append(c.Ch, genesisBlock)
}

func (c *Chain) getLatestBlock() block.Block {
	return c.Ch[len(c.Ch)-1]
}

func (c *Chain) AddTransaction(transaction transaction.Transaction_struct) {
	c.TransactionPool = append(c.TransactionPool, transaction)
}
func (c *Chain) AddBlockToChain(newBlock block.Block) {
	latestBlock := c.getLatestBlock()
	newBlock.PreviousHash = latestBlock.Hash
	newBlock.Mine(c.Difficulty)
	c.Ch = append(c.Ch, newBlock)
}

func (c *Chain) MineTransactionPool(minerRewardAddress *ecdsa.PublicKey) {
	minerRewardTransaction := transaction.NewTransaction(nil, minerRewardAddress, strconv.Itoa(c.MinerReward))
	c.TransactionPool = append(c.TransactionPool, minerRewardTransaction)
	latestBlockHash := c.getLatestBlock().Hash
	jsonTransactionPool, _ := json.Marshal(c.TransactionPool)
	newBlock := block.NewBlock(string(jsonTransactionPool), latestBlockHash)
	newBlock.Mine(c.Difficulty)
	c.Ch = append(c.Ch, newBlock)
	c.TransactionPool = make([]transaction.Transaction_struct, 0)
}

func (c *Chain) Validate() bool {
	if len(c.Ch) == 1 {
		if c.Ch[0].Hash != c.Ch[0].ComputeHash() {
			return false
		}
	}
	for index, blockToCheck := range c.Ch[1:] {
		if blockToCheck.Hash != blockToCheck.ComputeHash() || blockToCheck.PreviousHash != c.Ch[index].Hash {
			return false
		}
	}
	return true
}

func NewChain(difficulty int) Chain {
	bc := Chain{Difficulty: difficulty}
	bc.bigBang()
	return bc
}
