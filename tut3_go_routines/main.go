package main

import (
	"fmt"
)

func square(nums <-chan int, result chan<- int) {
	for n := range nums {
		result <- n * n
	}
}

func main() {

	nums := make(chan int)
	results := make(chan int)

	for i := 0; i < 4; i++ {
		go square(nums, results)
	}

	num_list := [...]int{1, 2, 3, 4, 1, 2}
	go func() {
		for _, val := range num_list {
			nums <- val
		}
		close(nums)

	}()

	for i := 0; i < len(num_list); i++ {
		fmt.Println(<-results)
	}

}
