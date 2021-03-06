package main

import "fmt"

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
}

// O(n^2)
func countTriplets(arr []int) int {
	n := len(arr)
	for i := 1; i < n; i++ {
		arr[i] ^= arr[i-1]
	}
	ans := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			num := arr[j]
			if i != 0 {
				num ^= arr[i-1]
			}

			if num == 0 {
				ans += j - i
			}
		}
	}

	return ans
}

// O(n^3)
//func countTriplets(arr []int) int {
//	n := len(arr)
//	for i := 1; i < n; i++ {
//		arr[i] ^= arr[i-1]
//	}
//	ans := 0
//
//	for i := 0; i < n-1; i++ {
//		for j := i + 1; j < n; j++ {
//			for k := j; k < n; k++ {
//				a := arr[j-1]
//				if i != 0 {
//					a ^= arr[i-1]
//				}
//
//				b := arr[k] ^ arr[j-1]
//
//				if a == b {
//					ans++
//				}
//			}
//		}
//	}
//	return ans
//}
