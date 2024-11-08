package models

import (
	"encoding/json"
	"fmt"
)

type Block struct {
	Number     int64  `json:"number" validate:"required,min=0"`
	Hash       string `json:"hash" validate:"required"`
	ParentHash string `json:"parentHash" validate:"required"`
	Timestamp  int64  `json:"timestamp" validate:"required,min=0"`
}

func (b *Block) ToString() string {
	return fmt.Sprintf("Block #%d hash=%s parentHash=%s", b.BlockNumber(), b.Hash, b.ParentHash)
}

func (b *Block) ToJSON() (string, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (b *Block) BlockNumber() int64 {
	return b.Number
}

func (b *Block) IsExactlySame(other *Block) bool {
	return b.BlockNumber() == other.BlockNumber() &&
		b.Hash == other.Hash &&
		b.ParentHash == other.ParentHash
}
