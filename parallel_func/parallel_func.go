package main

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

func (e *Executer) startTask(task Task, syncChannel chan interface{}) {
	defer e.wg.Done()
	var emptyVar interface{}
	syncChannel <- emptyVar
	if e.allowWork() {
		err := task()
		if err != nil {
			e.mu.Lock()
			e.errorCount += 1
			e.mu.Unlock()
		}

	}
	_ = <-syncChannel
}

func (e *Executer) startTasks(tasks []Task, maxErrorCount int, maxPacketExecution int) error {

	e.maxErrorCount = maxErrorCount
	syncChannel := make(chan interface{}, maxPacketExecution)
	for _, task := range tasks {
		if e.allowWork() {
			e.wg.Add(1)
			go e.startTask(task, syncChannel)
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
