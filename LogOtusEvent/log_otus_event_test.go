package LogOutusEvent

import (
	"fmt"
	"testing"
)

func TestSendHwAccepted(t *testing.T) {

	var event HwAccepted

	event.Id = 3456
	event.Grade = 4

	LogOtusEvent(event, logger)
}

func TestSendHwSubmitted(t *testing.T) {

	var event HwSubmitted

	event.Id = 3576
	event.Comment = "Some comment text"
	event.Code = "Some code"

	LogOtusEvent(event, logger)
}

// Helpers

type logWrite struct{}

func (logWrite) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return n, err
}

var logger logWrite
