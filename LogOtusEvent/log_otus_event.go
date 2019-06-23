package LogOutusEvent

import (
	"fmt"
	"io"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
}

func (e HwAccepted) logMessage() (result string) {
	result = fmt.Sprintf("accepted %v %v", e.Id, e.Grade)
	return
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (e HwSubmitted) logMessage() (result string) {
	result = fmt.Sprintf("submitted %v \"%v\"", e.Id, e.Comment)
	return
}

type OtusEvent interface {
	logMessage() string
}

func LogOtusEvent(e OtusEvent, w io.Writer) {

	date := time.Now().Format("02.01.2006 - 15:04:05")
	message := e.logMessage()
	w.Write([]byte(date + ": " + message))
}
