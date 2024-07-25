package main

// // ======= Parallel execution with goroutines =======
// import "fmt"

// func main() {
// 	ch := make(chan int)

// 	go compute(ch, 1)
// 	go compute(ch, 2)

// 	result1 := <-ch
// 	result2 := <-ch

// 	fmt.Println("Result 1:", result1)
// 	fmt.Println("Result 2:", result2)
// }

// func compute(ch chan int, num int) {
// 	result := num * 2
// 	ch <- result
// }

// // ======= Coordination and Synchronization with Channels =======
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan string)

// 	go producer(ch)
// 	go consumer(ch)

// 	time.Sleep(time.Second)
// }

// func producer(ch chan string) {
// 	ch <- "Data"
// 	fmt.Println("Sent: Data")
// 	close(ch)
// }

// func consumer(ch chan string) {
// 	for data := range ch {
// 		fmt.Println("Received:", data)
// 	}
// }

// ======= Select Statement for Multichannel Communication  =======
import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "World"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
