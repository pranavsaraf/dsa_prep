package main

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		smallestIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
	}
}

// func main() {
//     numbers := []int{64, 34, 25, 12, 22, 11, 90}
//     fmt.Println("Before sorting:", numbers)
//     SelectionSort(numbers)
//     fmt.Println("After sorting:", numbers)
// }
