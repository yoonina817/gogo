package main

import "fmt"

func main() {
	// Channel
	// Close : 채널 닫기

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 7777
		}
	}()

	val1, ok := <- ch
	fmt.Println("Example 1: ", val1, ok)

	val2, ok2 := <- ch
	fmt.Println("Example 1: ", val2, ok2)

	val3, ok3 := <- ch
	fmt.Println("Example 1: ", val3, ok3)

	close(ch) // 채널 닫기

	val4, ok4 := <- ch
	fmt.Println("Example 1: ", val4, ok4)

}