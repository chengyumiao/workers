package main

import (
	"fmt"
	"github.com/chengyumiao/workers"
	"runtime"
	"time"
)

// Create the concurrent go program
func ExampleCreate() {
	// The maximum number of concurrency is 5
	pools := workers.NewWorkersPool(5)

	for {
		pools.GetOneWorker() // get one worker
		go Runing(pools)
		time.Sleep(100*time.Microsecond)
	}

}

// create the concurrent go program with group wait.
func ExampleCreateWG() {

	// The maximum number of concurrency is 5
	pools := workers.NewWorkersPoolWithWG(5)

	array := [100]int{}

	//init array
	for i:=0; i<100; i++ {
		array[i] = i
	}

	for i:=0; i<100; i++ {
		pools.GetOneWorker()
		go convert(pools, &array[i])
	}

	fmt.Println(array)

}

func Runing(pool *workers.WorkersPool) {
	defer pool.ReleaseOneWorker() // release the worker
	time.Sleep(10*time.Second)
}

func convert(pool *workers.WorkersPool, c *int) {
	defer pool.ReleaseOneWorker()
	*c = (*c)*(*c)
}

func main() {

	// example1
	//go ExampleCreate()

	// example2
	go ExampleCreateWG()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(runtime.NumGoroutine())
	}

}