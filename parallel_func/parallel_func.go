package parallelfunc

import (
	"fmt"
	"sync"
)

type Task func() error

type Executer struct {
	mu            sync.Mutex
	errorCount    int
	maxErrorCount int
	taskCount     int
	wg            sync.WaitGroup
}

func (e *Executer) startTask(task Task) {
	defer e.wg.Done()
	if e.allowWork() {
		err := task()
		fmt.Printf("Return from task \"%v\" \n", err)
		if err != nil {
			e.mu.Lock()
			e.errorCount += 1
			e.mu.Unlock()
		}
	}
}

func (e *Executer) startTasks(tasks []Task, maxErrorCount int) error {

	e.maxErrorCount = maxErrorCount

	for _, task := range tasks {
		if e.allowWork() {
			e.wg.Add(1)
			go e.startTask(task)
			e.taskCount++
		}

	}

	e.wg.Wait()

	fmt.Println(e.errorCount)

	return nil

}

func (e *Executer) allowWork() bool {
	e.mu.Lock()
	result := e.errorCount < e.maxErrorCount
	e.mu.Unlock()

	return result
}
