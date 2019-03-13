package blockchain

import (
	"work_queue"
)

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

type miningWorker struct { //implements work_queue.Worker interface
  blk Block
	start_pos uint64
  end_pos uint64
}

func (mw miningWorker) Run() interface{} {
  res:=new(MiningResult)
  for i:= mw.start_pos; i<= mw.end_pos; i++{
    mw.blk.SetProof(i)
    if mw.blk.ValidHash(){ //we found a valid hash
      res.Proof = i
      res.Found = true
      return res  //returning MiningResult
    }
  }
  res.Found = false
  return res
}

func min(a,b uint64) uint64{
  if a<b {
    return a
  }else{
    return b
  }
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
  chunk_size := (end-start) / chunks //size of each individual chunk
  q := work_queue.Create(uint(workers), uint(chunks)) //create queue capabale of doing chunks tasks with workers threads

	startPos := start
  for i:=uint64(0); i<chunks; i++{
    endPos:= min(startPos+chunk_size, end)
    section := miningWorker{blk, startPos, endPos}
    q.Enqueue(section)
    startPos = endPos+1
  }
	res:=new(MiningResult)
	for r:= range q.Results{
		res = r.(*MiningResult)
		if(res.Found == true ){ //if we find a valid proof
			q.Shutdown()
			return *res
		}
	}
	return *res
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
