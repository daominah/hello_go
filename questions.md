# Golang test

#### 1. Nil channel

Fix the following code to get result sum = 21.

````go
package main
func main() {
	var resultChan chan int
	go func() {
		resultChan <- 1 + 2 + 3
	}()
	go func() {
		resultChan <- 4 + 5 + 6
	}()
	sum := 0
	sum += <-resultChan
	sum += <-resultChan
	println("sum:", sum)
}
````

#### 2. App can exit with running goroutines

Fix the following code to get result on stdout

````go
package main
func main() {
	go func() {
		sum := 0
		for i := 0; i < 5; i++ {
			sum += i
		}
		println("sum:", sum)
	}()
}
// result: empty stdout
````

#### 3. Concurrent write

What is the result if you run the following code? Fix it.

````go
package main

func main() {
	m := make(map[int]bool)
	done := make(chan bool)
	go func() {
		for i := 0; i < 1000; i++ {
			m[i] = true
		}
		done <- true
	}()
	for k := 3000; k < 4000; k++ {
		m[k] = true
	}
	<-done
	println(len(m))
}
````

#### 4. Loop variable in goroutine

Fix the following code to get result sum = 10.

````go
package main
import "sync"
func main() {
	sum := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			sum += i
		}()
	}
	wg.Wait()
	println("sum:", sum)
}
````
