package main

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		isSwap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwap = true
			}
		}

		if !isSwap {
			break
		}
	}
}

// func main() {
//     numbers := []int{64, 34, 25, 12, 22, 11, 90}
//     fmt.Println("Before sorting:", numbers)
//     BubbleSort(numbers)
//     fmt.Println("After sorting:", numbers)
// }

// T(n)= O(n^2)
// S(n)= O(1)
