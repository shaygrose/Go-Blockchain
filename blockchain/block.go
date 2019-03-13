package blockchain

import(
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	//"math"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	first := new(Block)
	first.PrevHash = []byte{}
	for i:=0; i<32; i++{
		first.PrevHash = append(first.PrevHash, '\x00')
	}

	first.Generation = 0
	first.Difficulty = difficulty
	first.Data = ""
	return *first
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	next:= new(Block)
	next.PrevHash = prev_block.Hash
	next.Generation = prev_block.Generation + 1
	next.Difficulty = prev_block.Difficulty
	next.Data = data
	return *next
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {

	prevHash:= hex.EncodeToString(blk.PrevHash)
	gen := fmt.Sprintf(":%d", blk.Generation)
	diff := fmt.Sprintf(":%d", blk.Difficulty)
	data := fmt.Sprintf(":%s", blk.Data)
	proof := fmt.Sprintf(":%d", blk.Proof)
	blk.Hash = []byte(prevHash + gen + diff + data + proof)

	hash := sha256.New()
	hash.Write(blk.Hash)
	sha:= hash.Sum(nil)
	return sha
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	nBytes := (blk.Difficulty)/8
	nBits := (blk.Difficulty)%8
	length := len(blk.Hash)

	if(length<=0){
		return false
	}

	for i:= (length-int(nBytes)); i<length; i++{
		if blk.Hash[i] != '\x00'{
			return false
		}
	}
	//check if next byte from the end is divisible by 2^nBits
	if((blk.Hash[length-int(nBytes)-1] % (1<<nBits)) !=0){
		return false
	}

	return true
}

//Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
