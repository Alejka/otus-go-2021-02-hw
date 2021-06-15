package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 1 {
		return ErrErrorsLimitExceeded
	}
	if n < 1 || len(tasks) < 1 {
		return nil
	}

	ch := make(chan Task, len(tasks))

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)

	worker := func(ch <-chan Task) {
		defer wg.Done()
		var stop bool
		for t := range ch {
			mu.Lock()
			stop = m < 1
			mu.Unlock()
			if stop {
				return
			}
			if err := t(); err != nil {
				mu.Lock()
				m--
				mu.Unlock()
			}
		}
	}
	for i := 0; i < n; i++ {
		go worker(ch)
	}

	for _, t := range tasks {
		ch <- t
	}
	close(ch)

	wg.Wait()

	if m < 1 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
