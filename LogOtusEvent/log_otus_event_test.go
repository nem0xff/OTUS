package LogOutusEvent

import (
	"fmt"
	"testing"
)

type HwAcceptedEvent struct {
	HwAccepted
}

func (e HwAcceptedEvent) logMessage() (result string) {
	result = fmt.Sprintf("accepted %v %v", e.Id, e.Grade)
	return
}

type HwSubmittedEvent struct {
	HwSubmitted
}

func (e HwSubmittedEvent) logMessage() (result string) {
	result = fmt.Sprintf("submitted %v \"%v\"", e.Id, e.Comment)
	return
}

func TestSendHwAccepted(t *testing.T) {

	var event HwAcceptedEvent

	event.Id = 3456
	event.Grade = 4

	LogOtusEvent(event, logger)
}

func TestSendHwSubmitted(t *testing.T) {

	var event HwSubmittedEvent

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
