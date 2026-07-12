//Special Pythagorean Triplet

package main

import "fmt"


func main() {
	limit := 1000
	for a:=2; a<=limit; a++ {
		for b:=1; b<=a; b++ {
			a, b, c := euclid_triplet(a, b)
			if a+b+c == 1000{
				fmt.Printf("Digits: %d, %d, %d", a, b, c)
				fmt.Println("PROD: ", a*b*c)
				break
			}

		}
	}
}

func euclid_triplet(m, n int) (int, int, int) {
	a := m*m - n*n
	b := 2*m*n
	c := m*m + n*n

	return a, b, c

}