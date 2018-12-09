package main

import (
	"errors"
	"sync"
)

type MyWaitGroup struct {
	wg             sync.WaitGroup
	numberRoutines int
	maxRoutines    int
	mux            sync.Mutex
}

func (mvg *MyWaitGroup) CanNext() bool {
	mvg.mux.Lock()
	defer mvg.mux.Unlock()
	println("can next?", mvg.numberRoutines < mvg.maxRoutines)
	return mvg.numberRoutines < mvg.maxRoutines
}

func (mvg *MyWaitGroup) Done() {
	mvg.wg.Done()
	mvg.mux.Lock()
	mvg.numberRoutines--
	mvg.mux.Unlock()
}

func (mvg *MyWaitGroup) Next() error {
	mvg.mux.Lock()
	defer mvg.mux.Unlock()
	if mvg.numberRoutines < mvg.maxRoutines {
		mvg.wg.Add(1)
		mvg.numberRoutines++
		return nil
	}
	println("too much goroutines")
	return errors.New("too much goroutines")
}

func (mvg *MyWaitGroup) Wait() {
	mvg.wg.Wait()
}
