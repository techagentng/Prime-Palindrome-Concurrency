package main

import (
	"fmt"
	"techagentng/pal/utils"
	"sync"
)

func main() {
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	results := make(chan int)
	var wg sync.WaitGroup

	// Start Goroutines to find prime palindromic numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		num := 2 
		for count < N {
			if utils.IsPrime(num) && utils.IsPalindrome(num) {
				results <- num
				count++
			}
			num++
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

// calc sum
	sum := 0
	for num := range results {
		sum += num
	}
	fmt.Printf("Sum of the first %d prime palindromic numbers: %d\n", N, sum)
}