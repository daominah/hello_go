package main

import "fmt"

func worker(workerId int, inChan chan *int, outChan chan int) {
	sum := 0
	for {
		data := <-inChan
		// a receive from a closed channel returns the zero value immediately
		if data == nil {
			break
		}
		fmt.Printf("worker %v received data: %v\n", workerId, *data)
		sum += *data
	}
	fmt.Printf("worker %v is about to return: %v\n", workerId, sum)
	outChan <- sum
}

func main() {
	nWorkers := 4
	inChan := make(chan *int)
	outChan := make(chan int)
	for i := 0; i < nWorkers; i++ {
		go worker(i, inChan, outChan)
	}
	for i := 0; i < 100; i++ {
		clonedValue := i
		inChan <- &clonedValue
	}
	close(inChan)
	sum := 0
	for i := 0; i < nWorkers; i++ {
		r := <-outChan
		sum += r
	}
	fmt.Println("sum", sum)
}
