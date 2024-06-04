package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return errors.New("Должна быть указана как минимум одна горутина")
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	var wg sync.WaitGroup
	if n > len(tasks) {
		n = len(tasks)
	}
	isErrorsLimitExceeded := false
	maxErrCount := 0

	//var maxErrorsCount = m
	var workersCount = n
	tasksChannel := make(chan Task)
	resultsChannel := make(chan error)
	wg.Add(1)
	go func() {
		for i := 0; i < len(tasks); i++ {
			if maxErrCount >= m && m > 0 {
				isErrorsLimitExceeded = true
				break
			} else {
				tasksChannel <- tasks[i]
				time.Sleep(time.Millisecond)
			}

		}
		close(tasksChannel) // закрываем канал по завершении
		fmt.Println("===channel closed===")
		wg.Done()
	}()
	for i := 0; i < workersCount; i++ {
		//i := i
		wg.Add(1)
		go func() {

			for {
				task, open := <-tasksChannel

				if !open {
					break
				}
				resultsChannel <- task()
			}
			wg.Done()
		}()
	}
	go func() {

		for {
			result, open := <-resultsChannel
			if !open {
				break
			}
			if result != nil {
				maxErrCount++
			}
		}
	}()
	wg.Wait()
	close(resultsChannel)

	if isErrorsLimitExceeded {
		return ErrErrorsLimitExceeded
	}
	return nil
}
