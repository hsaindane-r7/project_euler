// Even Fibonacci Numbers

package main


import "fmt"

func main() {
	sum := 0
	a, b := 0, 1
	for {
		c := a + b
		if c > 4_000_000 {
			break
		}
		if c % 2 == 0 {
			sum += c
		}
		fmt.Println(c)
		a = b
		b = c
	}
	fmt.Println("Sum of Even numbers in fibonacci nos before 4 mil: ", sum)
	
}
