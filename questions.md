# Golang test

#### 1. Nil channel

Fix the following code to get result sum = 10.

````go
package main
func main() {
	var resultChan chan int
	go func() {
		resultChan <- 1 + 2
	}()
	go func() {
		resultChan <- 3 + 4
	}()
	sum := 0
	sum += <-resultChan
	sum += <-resultChan
	println("sum:", sum)
}
````

#### 2. App can exit with active goroutines

Fix the following code to get "sum: 10" on stdout

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
````

#### 3. Loop variable in goroutine

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
