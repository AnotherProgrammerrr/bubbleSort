package main

import "fmt"

func bubbleSort(x []int) []int {
	n := len(x)
	for i := 0; i < n; i++ {
		for l := 0; l < n-i-1; l++ {
			if x[l] > x[l+1] {
				temp := x[l]
				x[l] = x[l+1]
				x[l+1] = temp
			}
		}
	}
	return x
}

func main() {
	x := []int{7, 5, 3, 9, 1, 0}
	fmt.Printf("%d", bubbleSort(x))
}