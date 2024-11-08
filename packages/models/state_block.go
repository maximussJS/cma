package models

import (
	"cma/packages/constants"
	"encoding/json"
	"fmt"
	"log"
)

type StateBlock struct {
	Number     int64  `json:"number" validate:"required,min=0"`
	Hash       string `json:"hash" validate:"required"`
	ParentHash string `json:"parentHash" validate:"required"`
	Timestamp  int64  `json:"timestamp" validate:"required,min=0"`
	Direction  string `json:"direction" validate:"required,oneof=-> <-"`
}

func (s StateBlock) String() string {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("Error marshalling StateBlock: %v", err)
	}
	return string(jsonData)
}

func IsEmptyStateBlock(block *StateBlock) bool {
	return block == nil
}

func (s *StateBlock) BlockNumber() int64 {
	return s.Number
}

func (s *StateBlock) IsExactlySame(other Block) bool {
	return s.BlockNumber() == other.BlockNumber() &&
		s.Hash == other.Hash &&
		s.ParentHash == other.ParentHash
}

func (s *StateBlock) DirectionString() constants.DirectionString {
	direction, err := constants.FromString(s.Direction)

	if err != nil {
		log.Fatalf("Invalid direction string %s", s.Direction)
	}

	return direction
}

func FromBlock(block Block, directionString constants.DirectionString) StateBlock {
	direction, err := constants.ToString(directionString)

	if err != nil {
		log.Fatalf("Invalid direction string %s", directionString)
	}

	return StateBlock{
		Number:     block.Number,
		Hash:       block.Hash,
		ParentHash: block.ParentHash,
		Timestamp:  block.Timestamp,
		Direction:  direction,
	}
}

func (s *StateBlock) ToBlock() Block {
	return Block{
		Number:     s.Number,
		Hash:       s.Hash,
		ParentHash: s.ParentHash,
		Timestamp:  s.Timestamp,
	}
}
