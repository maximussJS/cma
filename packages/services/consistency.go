package services

import (
	"cma/packages/constants"
	"cma/packages/errors"
	"cma/packages/models"
	"fmt"
)

type Consistency struct {
	state *State
}

func NewConsistency() *Consistency {
	return &Consistency{state: StateSingletonInstance()}
}

func (c *Consistency) Check(blockchain constants.BlockchainType, blocks []models.Block, direction constants.DirectionString) error {
	ok, err := c.isFirstBlockNextToState(blockchain, blocks, direction)

	if ok != true || err != nil {
		return err
	}

	if err := c.checkBlocksOrder(blockchain, direction, blocks); err != nil {
		return err
	}

	return c.commit(blockchain, blocks, direction)
}

func (c *Consistency) checkBlocksOrder(blockchain constants.BlockchainType, direction constants.DirectionString, blocks []models.Block) error {
	for i := 1; i < len(blocks); i++ {
		prevBlock := blocks[i-1]
		block := blocks[i]

		if !c.isNextBlock(direction, block, prevBlock) {
			return errors.NewBlockOrderError(blockchain, prevBlock, block, direction)
		}
	}
	return nil
}

func (c *Consistency) isFirstBlockNextToState(blockchain constants.BlockchainType, blocks []models.Block, direction constants.DirectionString) (bool, error) {
	state, err := c.state.Get(blockchain)

	if err != nil {
		return false, fmt.Errorf("Consistency.isFirstBlockNextToState() cannot get state block %w", err)
	}

	firstBlock := blocks[0]
	if state == nil || models.IsEmptyStateBlock(state) {
		if firstBlock.BlockNumber() != 0 {
			return false, errors.NewFirstBlockNotZeroError(blockchain, firstBlock, direction)
		}
		return true, nil
	}

	if state.IsExactlySame(firstBlock) && state.DirectionString() == direction {
		return false, errors.NewFirstBlockIsTheSameAsStateError(blockchain, firstBlock, direction)
	}

	ok, err := c.isNextStateBlock(blockchain, firstBlock, direction)
	if ok == false || err != nil {
		return false, errors.NewBlockOrderError(blockchain, state.ToBlock(), firstBlock, direction)
	}

	return true, nil
}

func (c *Consistency) commit(blockchain constants.BlockchainType, blocks []models.Block, direction constants.DirectionString) error {
	lastBlock := blocks[len(blocks)-1]
	state := models.FromBlock(lastBlock, direction)
	if err := c.state.Set(blockchain, &state); err != nil {
		return fmt.Errorf("Consistency.commit() set state error %w", err)
	}
	fmt.Printf("Committed block %d in blockchain %s with direction %s\n", lastBlock.BlockNumber(), blockchain, direction)
	return nil
}

func (c *Consistency) isNextStateBlock(blockchain constants.BlockchainType, nextBlock models.Block, nextDirection constants.DirectionString) (bool, error) {
	state, err := c.state.Get(blockchain)

	if err != nil {
		return false, fmt.Errorf("Consistency.isNextStateBlock() cannot get state block %w", err)
	}

	prevBlock := state.ToBlock()
	prevDirection := state.DirectionString()

	if nextDirection == constants.Forward && prevDirection == constants.Forward {
		return nextBlock.ParentHash == prevBlock.Hash && nextBlock.BlockNumber() == prevBlock.BlockNumber()+1, nil
	}

	if nextDirection == constants.Backward && prevDirection == constants.Backward {
		return nextBlock.Hash == prevBlock.ParentHash && nextBlock.BlockNumber()+1 == prevBlock.BlockNumber(), nil
	}

	if nextDirection == constants.Backward && prevDirection == constants.Forward {
		return nextBlock.Hash == prevBlock.Hash && nextBlock.BlockNumber() == prevBlock.BlockNumber(), nil
	}

	if nextDirection == constants.Forward && prevDirection == constants.Backward {
		return nextBlock.ParentHash == prevBlock.ParentHash && nextBlock.BlockNumber() == prevBlock.BlockNumber(), nil
	}

	return false, nil
}

func (c *Consistency) isNextBlock(direction constants.DirectionString, nextBlock, prevBlock models.Block) bool {
	if direction == constants.Forward {
		return nextBlock.ParentHash == prevBlock.Hash && nextBlock.BlockNumber() == prevBlock.BlockNumber()+1
	}

	if direction == constants.Backward {
		return nextBlock.Hash == prevBlock.ParentHash && nextBlock.BlockNumber()+1 == prevBlock.BlockNumber()
	}

	return false
}
