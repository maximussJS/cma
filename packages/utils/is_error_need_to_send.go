package utils

import (
	custom_errors "cma/packages/errors"
	"errors"
)

var blockOrderError *custom_errors.BlockOrderError
var firstBlockIsTheSameAsStateError *custom_errors.FirstBlockIsTheSameAsStateError
var firstBlockNotZeroError *custom_errors.FirstBlockNotZeroError

func IsErrorNeedToSend(err error) bool {
	return errors.As(err, &blockOrderError) || errors.As(err, &firstBlockNotZeroError) || errors.As(err, &firstBlockIsTheSameAsStateError)
}
