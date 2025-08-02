package main 

import "fmt"

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		current := arr[i]
		prev := i - 1

		for prev >= 0 && arr[prev] > current {
			arr[prev+1] = arr[prev]
			prev--
		}
		arr[prev+1] = current
	}
}

func main() {	
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Before sorting:", numbers)
	InsertionSort(numbers)
	fmt.Println("After sorting:", numbers)
}