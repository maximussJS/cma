package errors

import (
	"cma/packages/constants"
	"cma/packages/models"
	"fmt"
)

type FirstBlockIsTheSameAsStateError struct {
	Blockchain constants.BlockchainType
	Block      models.Block
	Direction  constants.DirectionString
}

func NewFirstBlockIsTheSameAsStateError(blockchain constants.BlockchainType, block models.Block, direction constants.DirectionString) *FirstBlockIsTheSameAsStateError {
	return &FirstBlockIsTheSameAsStateError{
		Blockchain: blockchain,
		Block:      block,
		Direction:  direction,
	}
}

func (err *FirstBlockIsTheSameAsStateError) Error() string {
	return fmt.Sprintf(
		":large_yellow_circle:CMA Warning:large_yellow_circle:\n"+
			"First block is the same as the state block:\n"+
			"  • Blockchain: %s\n"+
			"  • Direction: %s\n"+
			"  • Block #%d\n"+
			"  • Hash: %s\n"+
			"  • Parent Hash: %s\n",
		err.Blockchain,
		err.Direction,
		err.Block.Number,
		err.Block.Hash,
		err.Block.ParentHash,
	)
}
