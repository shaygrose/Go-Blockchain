package blockchain

import(
	"encoding/hex"
)


type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}else if blk.ValidHash(){
		//add the block
		chain.Chain = append(chain.Chain, blk)
	}
}

func (chain Blockchain) IsValid() bool {

	//initial block has generation 0
	if chain.Chain[0].Generation != 0{
		return false
	}

	//initial block has previous hash all null bybtes
	for i:= 0; i<32; i++{
		if chain.Chain[0].PrevHash[i] != '\x00'{
			return false
		}
	}

	length := len(chain.Chain)

	for i:=1; i<length; i++{
		if chain.Chain[i].Generation != chain.Chain[i-1].Generation + 1 ||
		chain.Chain[i].Difficulty != chain.Chain[0].Difficulty ||
		hex.EncodeToString(chain.Chain[i].PrevHash) != hex.EncodeToString(chain.Chain[i-1].Hash) ||
		hex.EncodeToString(chain.Chain[i].Hash) != hex.EncodeToString(chain.Chain[i].CalcHash()) ||
		!chain.Chain[i].ValidHash(){
			return false
		}
	}
	return true
}
