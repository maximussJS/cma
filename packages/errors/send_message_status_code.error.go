package errors

import "fmt"

type SendMessageStatusCode struct {
	statusCode int
}

func NewSendMessageStatusCode(statusCode int) *SendMessageStatusCode {
	return &SendMessageStatusCode{statusCode: statusCode}
}

func (err *SendMessageStatusCode) Error() string {
	return fmt.Sprintf("Status code is not 200, it is %d", err.statusCode)
}
