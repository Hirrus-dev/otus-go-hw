package hw05parallelexecution

import (
	"errors"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return errors.New("должна быть указана как минимум одна горутина")
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	if n > len(tasks) {
		n = len(tasks)
	}
	isErrorsLimitExceeded := false
	maxErrCount := 0
	tasksChannel := make(chan Task)
	resultsChannel := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < len(tasks); i++ {
			if maxErrCount >= m && m > 0 {
				isErrorsLimitExceeded = true
				break
			}
			tasksChannel <- tasks[i]
			time.Sleep(time.Millisecond)
		}
		close(tasksChannel) // закрываем канал по завершении
		wg.Done()
	}()
	for i := 0; i < n; i++ {
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
