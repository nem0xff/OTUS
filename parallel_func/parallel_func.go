package parallelfunc

import (
	"fmt"
	"sync"
)

func Executer(functions []func() error) error {
	var wg sync.WaitGroup

	var mu sync.Mutex
	var errorCount int

	for _, function := range functions {
		fmt.Println(errorCount)
		if errorCount > 2 {
			break
		}
		wg.Add(1)
		go func(funcForExec func() error) {
			if errorCount > 2 {
				fmt.Println("Кол-во ошибок: ", errorCount)
				wg.Done()
				return
			}
			err := funcForExec()
			if err != nil {
				mu.Lock()
				errorCount++
				mu.Unlock()
			}

			wg.Done()

		}(function)

	}
	wg.Wait()
	return nil
}
