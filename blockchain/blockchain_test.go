package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"crypto/sha256"
	"encoding/hex"
	//"fmt"
)

// TODO: some useful tests of Blocks

func TestCreatingBlocks(t *testing.T){
	b0:=Initial(uint8(16))
	b1:= b0.Next("message")

	assert.Equal(t, uint8(16), b0.Difficulty, "Difficulty is wrong in Intial Block")
	assert.Equal(t, uint64(0), b0.Generation, "Generation is wrong in Initial Block")
	//assert.Equal(t, []byte("\x00"), b0.PrevHash, "PrevHash is wrong in Next Block")
	assert.Equal(t, uint8(16), b1.Difficulty, "Difficulty is wrong in Next Block")
	assert.Equal(t, uint64(1), b1.Generation, "Generation is wrong in Next Block")
	assert.Equal(t, "message", b1.Data, "Data is wrong in Next Block")

}

func TestCalcHash(t *testing.T){

	b0:=Initial(uint8(16))
	b0.SetProof(56231)
	res:="6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000"

	assert.Equal(t, res, hex.EncodeToString(b0.CalcHash()))

	b1:=b0.Next("message")
	b1.SetProof(2159)
	res2:="9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000"

	assert.Equal(t, res2, hex.EncodeToString(b1.CalcHash()))


}


func TestValidHash(t *testing.T){
	b0 := Initial(19)
	b0.SetProof(87745)

	//b0 := Initial(16)
	//b0.SetProof(56231)

	assert.Equal(t, true, b0.ValidHash())

	b1 := b0.Next("hash example 1234")
	b1.SetProof(1407891)

	assert.Equal(t, true, b1.ValidHash())

	//b1.SetProof(346082)
	//assert.Equal(t, false, b1.ValidHash())

}


func TestMine(t *testing.T){

	b0 := Initial(20)
	b0.Mine(1)
	//fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	//fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	//fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))

}


func TestAdd(t *testing.T){
	b0 := Initial(7)
	b0.Mine(1)

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)

	blockchain := new(Blockchain)
	blockchain.Add(b0)
	blockchain.Add(b1)
	blockchain.Add(b2)

	assert.Equal(t, blockchain.Chain[1].Data, "this is an interesting message")
	assert.Equal(t, blockchain.Chain[2].Difficulty, uint8(7))
	assert.Equal(t, len(blockchain.Chain), 3)
}


func TestIsValid(t *testing.T){

	b0 := Initial(7)
	b0.Mine(1)

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b1.Generation = 3

	badBlockchain := new(Blockchain)
	badBlockchain.Add(b0)
	badBlockchain.Add(b1)

	assert.Equal(t, false, badBlockchain.IsValid())

	bl0 := Initial(7)
	bl0.Mine(1)

	bl1 := bl0.Next("this is an interesting message")
	bl1.Mine(1)

	bl2 := bl1.Next("this is not interesting")
	bl2.Mine(1)

	goodBlockchain := new(Blockchain)
	goodBlockchain.Add(bl0)
	goodBlockchain.Add(bl1)
	goodBlockchain.Add(bl2)

	assert.Equal(t, true, goodBlockchain.IsValid())

}
