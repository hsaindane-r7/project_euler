//Sum Square Difference

package main

import "fmt"


func main() {
	var sum, square_sum int
	for i:= 1; i<=100; i++ {
		sum += i
		square_sum += (i * i)
	}
	sum = sum * sum
	// ans := fmt.Sprintf("Sum: %d and Square sum: %d", (sum), square_sum)
	// fmt.Println("Ans: ", ans)
	fmt.Println("Difference between sum of squares of first 100 natural nos and square of the sum: ",sum - square_sum)
}