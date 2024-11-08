package errors

import (
	"cma/packages/constants"
	"cma/packages/models"
	"fmt"
)

type BlockOrderError struct {
	Blockchain constants.BlockchainType
	PrevBlock  models.Block
	CurBlock   models.Block
	Direction  constants.DirectionString
}

func NewBlockOrderError(blockchain constants.BlockchainType, prevBlock, curBlock models.Block, direction constants.DirectionString) *BlockOrderError {
	return &BlockOrderError{
		Blockchain: blockchain,
		PrevBlock:  prevBlock,
		CurBlock:   curBlock,
		Direction:  direction,
	}
}

func (err *BlockOrderError) Error() string {
	return fmt.Sprintf(
		":red_circle:CMA Alert:red_circle:\n"+
			"Block State Order Mismatch Details:\n"+
			"• Blockchain: %s\n"+
			"• Block #%d\n"+
			"  • Hash: %s\n"+
			"  • Parent Hash: %s\n"+
			"  • Direction: %s\n"+
			"• Block in State #%d\n"+
			"  • Hash: %s\n"+
			"  • Parent Hash: %s\n"+
			"  • Direction: %s\n"+
			"Please investigate this issue.",
		err.Blockchain,
		err.PrevBlock.Number,
		err.PrevBlock.Hash,
		err.PrevBlock.ParentHash,
		err.Direction,
		err.CurBlock.Number,
		err.CurBlock.Hash,
		err.CurBlock.ParentHash,
		err.Direction,
	)
}
