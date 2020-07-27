package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	// cpu prof
	file, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = pprof.StartCPUProfile(file)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// mem prof
	defer func() {
		file2, err := os.Create("mem.out")
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()
		runtime.GC()
		err = pprof.WriteHeapProfile(file2)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// example computation
	fakeSum := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			fakeSum += i
			time.Sleep(1 * time.Second)
			wg.Add(-1)
		}(i)
	}
	wg.Wait()
	fmt.Println("fakeSum: ", fakeSum)
}
