package errors

import (
	"cma/packages/constants"
	"cma/packages/models"
	"fmt"
)

type FirstBlockNotZeroError struct {
	Blockchain constants.BlockchainType
	FirstBlock models.Block
	Direction  constants.DirectionString
}

func NewFirstBlockNotZeroError(blockchain constants.BlockchainType, firstBlock models.Block, direction constants.DirectionString) *FirstBlockNotZeroError {
	return &FirstBlockNotZeroError{
		Blockchain: blockchain,
		FirstBlock: firstBlock,
		Direction:  direction,
	}
}

func (err *FirstBlockNotZeroError) Error() string {
	return fmt.Sprintf(
		"*🔴CMA Alert🔴*\n"+
			"*State is empty, but the first block number is not 0.*\n"+
			"• *Blockchain:* `%s`\n"+
			"• *Block #%d*\n"+
			"  • *Hash:* `%s`\n"+
			"  • *Parent Hash:* `%s`\n"+
			"  • *Direction:* `%s`\n\n"+
			"Please investigate this issue.",
		err.Blockchain,
		err.FirstBlock.Number,
		err.FirstBlock.Hash,
		err.FirstBlock.ParentHash,
		err.Direction,
	)
}
