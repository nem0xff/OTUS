package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(4)
}

const (
	MAX_TIME_EXECUTION   = 500  // Максимальное время выполнения создаваемых функций
	MAX_ERROR_COUNT      = 3    // Максимальное допустимое количество ошибок после которого останавливаем если еще что-то осталось
	NUMBER_FUNCS         = 1000 //количество функций
	MAX_PACKET_EXECUTION = 5    // Максимальное количество одновременно выполняемых заданий
)

func main() {
	var executer Executer

	pf := getFuncs(NUMBER_FUNCS)
	_ = executer.startTasks(pf[:], MAX_ERROR_COUNT, MAX_PACKET_EXECUTION)
}

// Создаем массив из функций с различным временем выполнения.
func getFuncs(n int) []Task {
	var result []Task
	for i := 0; i < n; i++ {
		z := i
		timeExec := time.Millisecond * time.Duration(rand.Intn(MAX_TIME_EXECUTION))
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
