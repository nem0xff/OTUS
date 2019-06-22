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

type OtusEvent interface{}

func LogOtusEvent(e OtusEvent, w io.Writer) {

	date := time.Now().Format("02.01.2006 - 15:04:05")

	switch val := e.(type) {
	case HwAccepted:
		fmt.Fprintf(w, "%v: accepted %v %v", date, val.Id, val.Grade)
		return
	case HwSubmitted:
		fmt.Fprintf(w, "%v: submitted %v \"%v\"", date, val.Id, val.Comment)
		return
	}

	fmt.Fprintf(w, "%v: got some new type %T", date, e)
}
