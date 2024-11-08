package constants

import (
	"encoding/json"
	"fmt"
)

type BlockchainType string

const (
	Base     BlockchainType = "base"
	Ethereum BlockchainType = "ethereum"
)

func NewBlockchainType(value string) (BlockchainType, error) {
	bt := BlockchainType(value)
	if !bt.IsValid() {
		return "", fmt.Errorf("invalid blockchain type: %s", value)
	}
	return bt, nil
}

func (b BlockchainType) IsValid() bool {
	switch b {
	case Base, Ethereum:
		return true
	default:
		return false
	}
}

func (b BlockchainType) String() string {
	return string(b)
}

func (b *BlockchainType) UnmarshalJSON(data []byte) error {
	value := string(data[1 : len(data)-1]) // Trim quotes
	bt := BlockchainType(value)
	if !bt.IsValid() {
		return fmt.Errorf("invalid blockchain type: %s", value)
	}
	*b = bt
	return nil
}

func (b BlockchainType) MarshalJSON() ([]byte, error) {
	if !b.IsValid() {
		return nil, fmt.Errorf("cannot marshal invalid blockchain type: %s", b)
	}
	return json.Marshal(string(b))
}
