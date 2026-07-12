// 10 001st Prime

package main

import "fmt"

func main() {
	n := 1000000
	arr := make([]bool, n+1)
	for i := range arr {
		arr[i] = true
	}

	arr[0], arr[1] = false, false
	cnt := 0

	for i:= 2; i*i < len(arr); i++ {
		if arr[i] == true {
			for j := i*2; j<len(arr); j = j+i {
				arr[j] = false
			}
		}
	}

	for i:= 2; i<=len(arr); i++ {
		if arr[i] == true {
			cnt ++
			if cnt == 10001 {
				fmt.Println("10,001 th prime number: ", i)
				break
			}
		}
	}


}