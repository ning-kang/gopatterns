package main

import (
	"fmt"
	"sync"
)

// sync.Once is another approach
var lock = &sync.Mutex{}

// The example single type
type single struct{}

// the single instance
var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// This is to ensure that if more than one goroutine bypasses the first check,
		// only one goroutine can create the singleton instance.
		// Otherwise, all goroutines will create their own instances of the singleton struct.
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
