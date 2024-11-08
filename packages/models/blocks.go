package models

import (
	"cma/packages/constants"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

type Blocks struct {
	Data      []Block `json:"data" validate:"required,dive"`
	Direction string  `json:"direction" validate:"required,oneof=-> <-"`
}

func (b *Blocks) DirectionString() constants.DirectionString {
	direction, err := constants.FromString(b.Direction)

	if err != nil {
		log.Fatalf("Invalid direction string %s", b.Direction)
	}

	return direction
}

func (b *Blocks) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Direction: %s\n", b.Direction))

	for i, block := range b.Data {
		builder.WriteString(fmt.Sprintf("Block #%d:\n", i+1))
		builder.WriteString(fmt.Sprintf("  - Number: %d\n", block.Number))
		builder.WriteString(fmt.Sprintf("  - Hash: %s\n", block.Hash))
		builder.WriteString(fmt.Sprintf("  - ParentHash: %s\n", block.ParentHash))
		builder.WriteString(fmt.Sprintf("  - Timestamp: %d\n", block.Timestamp))
		builder.WriteString("\n")
	}

	return builder.String()
}

func (b *Blocks) UnmarshalJSON(data []byte) error {
	var temp struct {
		Data      []json.RawMessage `json:"data"`
		Direction string            `json:"direction"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Direction != "->" && temp.Direction != "<-" {
		return errors.New("invalid direction value; must be '->' or '<-'")
	}

	var blocks []Block
	for _, raw := range temp.Data {
		var block Block
		if err := json.Unmarshal(raw, &block); err != nil {
			return err
		}
		blocks = append(blocks, block)
	}

	b.Data = blocks
	b.Direction = temp.Direction

	return nil
}
