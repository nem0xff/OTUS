package LogOutusEvent

import (
	"fmt"
	"testing"
)

type logWrite struct{}

func (logWrite) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return n, err
}

var logger logWrite

func TestSendHwAccepted(t *testing.T) {

	var hwAccept HwAccepted

	hwAccept.Grade = 5
	hwAccept.Id = 1

	LogOtusEvent(hwAccept, logger)
}

func TestSendHwSubmitted(t *testing.T) {

	var hwSubmitted HwSubmitted

	hwSubmitted.Id = 3576
	hwSubmitted.Comment = "Some comment text"
	hwSubmitted.Code = "Some code"

	LogOtusEvent(hwSubmitted, logger)
}
