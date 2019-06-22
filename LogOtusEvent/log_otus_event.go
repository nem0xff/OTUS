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

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type OtusEvent interface {
	logMessage() string
}

func LogOtusEvent(e OtusEvent, w io.Writer) {

	date := time.Now().Format("02.01.2006 - 15:04:05")
	message := e.logMessage()

	fmt.Fprintf(w, "%v: %v", date, message)
}
