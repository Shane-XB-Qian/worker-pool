package workerpool_test

import (
	"fmt"
	"sync"

	workerpool "github.com/vardius/worker-pool"
)

func Example() {
	var wg sync.WaitGroup

	poolSize := 1
	jobsAmount := 3
	workersAmount := 2

	// create new pool
	pool := workerpool.New(poolSize)
	out := make(chan int, jobsAmount)

	pool.Start(workersAmount, func(i int) {
		defer wg.Done()
		out <- i
	})

	wg.Add(jobsAmount)

	for i := 0; i < jobsAmount; i++ {
		pool.Delegate(i)
	}

	go func() {
		// stop all workers after jobs are done
		wg.Wait()
		close(out)
		pool.Stop()
	}()

	sum := 0
	for n := range out {
		sum += n
	}

	fmt.Println(sum)
	// Output:
	// 3
}
