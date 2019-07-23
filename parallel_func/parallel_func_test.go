package parallelfunc

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestExecuter(t *testing.T) {
	var executer Executer

	pf := getFuncs(100)
	_ = executer.startTasks(pf[:], 2)

}

func getFuncs(n int) []Task {
	runtime.GOMAXPROCS(4)
	var result []Task
	for i := 0; i < n; i++ {
		z := i
		timeExec := time.Millisecond * time.Duration(rand.Intn(1000))
		f := func() error {
			time.Sleep(timeExec)
			fmt.Printf("Message from func #%v, time exection %v \n", z, timeExec)
			if rand.Intn(100) > 80 {
				return errors.New("error in function")
			} else {
				return nil
			}

		}
		result = append(result, f)
	}

	return result
}
