// Largest Palindrome Product

package main

import (
	// "fmt"
	"fmt"
	"strconv"
	// "os"
)

func main() {
	largest_no := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			prod := i * j
			if isPalindrome(prod) {
				if prod > largest_no {
					largest_no = prod
				}
				// msg := fmt.Sprinf("Largest palindrome from product of two 3-digit number: %d x %d = %d", i, j, prod)
				// fmt.Println(msg)
				// os.Exit(3)
			}
		}
	}
	fmt.Println(largest_no)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true

}
