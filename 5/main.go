// Smallest Multiple
package main

import "fmt"

func main() {
	i := 1
	for {
		if isDivisible(i){
			fmt.Println("Smallest no divisible by 1-20: ", i)
			break
		}
		i++
	}
}

func isDivisible(n int) bool {
	for i := 11; i<=20; i++ {
		if n % i != 0 {
			return false
		}
	}
	return true
}

// ans = 232792560