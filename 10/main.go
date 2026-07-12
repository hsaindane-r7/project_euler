// Summation of Primes

package main 

import "fmt"

func main() {
	n := 2_000_000
	arr := make([]bool, n+1)
	for i := range arr {
		arr[i] = true
	}
	arr[0], arr[1] = false, false
	for i := 2; i*i<len(arr); i++ {
		if arr[i] == true {
			for j := i*2; j < len(arr); j += i {
				arr[j] = false
			}
		}
	}
	sum := 0
	for i := range arr {
		if arr[i] {
			sum += i
		}
	}
	fmt.Printf("Sum of %d prime elements: %d", len(arr)-1, sum)
}

/*
% go run main.go
Sum of 2000000 prime elements: 142913828922% 
*/