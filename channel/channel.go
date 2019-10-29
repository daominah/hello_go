package main

import "log"

func worker(workerId int, inChan chan *int, outChan chan int) {
	sum := 0
	for {
		data := <-inChan
		// a receive from a closed channel returns the zero value immediately
		if data == nil {
			break
		}
		//log.Printf("worker %v received data: %v\n", workerId, *data)
		sum += *data
	}
	log.Printf("worker %v is about to return: %v\n", workerId, sum)
	outChan <- sum
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	nWorkers := 4
	inChan := make(chan *int)
	outChan := make(chan int)
	for i := 0; i < nWorkers; i++ {
		go worker(i, inChan, outChan)
	}

	log.Println("starting the main")
	for i := 0; i < 10000000; i++ {
		clonedValue := i
		inChan <- &clonedValue
	}
	close(inChan)
	sum := 0
	for i := 0; i < nWorkers; i++ {
		r := <-outChan
		sum += r
	}
	log.Println("sum:", sum)
}
