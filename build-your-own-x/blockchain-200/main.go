package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block is a block of data written to the blockchain
type Block struct {
	Index     int    // the position of the data record in the blockchain
	Timestamp string // automatically determined and is the time the data is written
	BPM       int    // "beats per minute" - pulse rate
	Hash      string // SHA256 identifier representing this data record
	PrevHash  string // SHA256 identifier of the previous record in the chain
}

// Blockchain represents the set of blocks
var Blockchain []Block

func main() {

}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
