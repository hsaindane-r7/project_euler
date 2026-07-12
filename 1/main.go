package main

import "fmt"

func main() {
	sum := 0
	for i:=1; i<=1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of all multiples of 3 or 5 from 1 - 1000 =  ", sum)
}