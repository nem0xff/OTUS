package parallelfunc

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func TestExecuter(t *testing.T) {

	pf := getFuncs(20)
	_ = Executer(pf[:])

}

func getFuncs(n int) []func() error {

	var result []func() error
	for i := 0; i < n; i++ {
		z := i

		f := func() error {
			fmt.Println("Message from func #", z)

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
