package constants

import (
	"fmt"
)

type DirectionString string

const (
	Forward  DirectionString = "FORWARD"
	Backward DirectionString = "BACKWARD"
)

func FromString(direction string) (DirectionString, error) {
	switch direction {
	case "->":
		return Forward, nil
	case "<-":
		return Backward, nil
	default:
		return "", fmt.Errorf("invalid direction: %s", direction)
	}
}

func ToString(direction DirectionString) (string, error) {
	switch direction {
	case Forward:
		return "->", nil
	case Backward:
		return "<-", nil
	default:
		return "", fmt.Errorf("invalid direction: %s", direction)
	}
}
