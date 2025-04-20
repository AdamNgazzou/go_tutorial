package main

// this is what multithreading fetching from the database at the same time (5 server requests in the same time)
// but it may get some unexpected results (miss data) when two processes writing to the same memory location at
// the same time may lead to corrupt data ()
// that's why  we use "mutex" (mutuelle exclusion ) using lock and unlock methods
import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

// create a wait group
var wg = sync.WaitGroup{}

// create mutex
var m = sync.RWMutex{}

func main() {

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// here is what slow version without running cuncurriently :  dbCall(i)
		//faster version using goroutines and sync:
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Printf("\nTotal excution time : %v", time.Since(t0))
	fmt.Printf("\n the results are %v", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\n the current results are : %v", results)
	m.RUnlock()

}
