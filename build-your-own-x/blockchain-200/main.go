package main

// Block is a block of data written to the blockchain
type Block struct {
	Index     int    // the position of the data record in the blockchain
	Timestamp string // automatically determined and is the time the data is written
	BPM       int    // "beats per minute" - pulse rate
	Hash      string // SHA256 identifier representing this data record
	PrevHash  string // SHA256 identifier of the previous record in the chain
}

func main() {

}
