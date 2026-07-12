// Largest Prime factor (Sieve of Eratosthenes)

package main

import "fmt"

func main() {
	// n := 600851475143
	// // n := 100
	// nos := make([]bool, n+1)
	// for i := range nos {
	// 	nos[i] = true
	// }
	// nos[0], nos[1] = false, false

	// for i := 2; i*i < len(nos); i++ {
	// 	if nos[i] == true {
	// 		for j := i*2; j < len(nos); j += i {
	// 			// fmt.Println("Flipping multiple of", i)
	// 			nos[j] = false
	// 		}
	// 	}
	// 	if i % 10 == 0{
	// 		fmt.Println("Progress: On index: ", i)
	// 	}
	// }

	// for i := len(nos) - 1; i >= 0; i-- {
	// 	if nos[i] == true {
	// 		// fmt.Println("Prime num: ", i)
	// 		if n % i == 0{
	// 			fmt.Println("Largest prime factor of 600851475143: ", i)
	// 			break
	// 		}
	// 	}
		
	// }

//-------------------- Not the most optimal soln (requires 600 GB ram ). `slice` is taking 600851475143 bytes ~ 600GB

	n := 600851475143

	fct := 2

	for fct * fct <= n {
		if n % fct == 0 {
			n /= fct
		} else {
			fct ++
		}
	}

	fmt.Println("Largest prime factor: ", n)
}