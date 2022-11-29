package main

import "fmt"

func main() {
	integers := make([]int, 11)

	for i := 0; i <= 10; i++ {
		integers[i] = i
	}

	for _, v := range integers {
		if v%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
