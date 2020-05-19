package main

import (
	"crypto/sha256"
	"encoding/hex"
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
