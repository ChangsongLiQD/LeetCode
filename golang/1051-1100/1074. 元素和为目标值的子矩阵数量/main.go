package main

import "fmt"

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0))
}

func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{0: 1}
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			ans += mp[pre-k]
		}
		mp[pre]++
	}
	return
}

func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	for i := range matrix { // 枚举上边界
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				sum[c] += v // 更新每列的元素和
			}
			ans += subarraySum(sum, target)
		}
	}
	return
}

func numSubmatrixSumTargetBestMemory(matrix [][]int, target int) int {
	n := len(matrix)
	m := len(matrix[0])
	ans := 0

	for x2 := 0; x2 < n; x2++ {
		for y2 := 0; y2 < m; y2++ {
			topLeft, topRight, left := 0, 0, 0
			if y2 > 0 {
				left = matrix[x2][y2-1]
			}
			if x2 > 0 {
				topRight = matrix[x2-1][y2]
			}
			if x2 > 0 && y2 > 0 {
				topLeft = matrix[x2-1][y2-1]
			}
			value := matrix[x2][y2]
			matrix[x2][y2] += topRight + left - topLeft

			for x1 := 0; x1 <= x2; x1++ {
				for y1 := 0; y1 <= y2; y1++ {
					if x2 == x1 && y2 == y1 {
						if value == target {
							ans++
						}
					} else {
						sumTop, sumLeft, sumBoth := 0, 0, 0
						if x1 > 0 {
							sumTop = matrix[x1-1][y2]
						}
						if y1 > 0 {
							sumLeft = matrix[x2][y1-1]
						}
						if x1 > 0 && y1 > 0 {
							sumBoth = matrix[x1-1][y1-1]
						}

						if matrix[x2][y2]-sumTop-sumLeft+sumBoth == target {
							ans++
						}
					}
				}
			}
		}
	}

	return ans
}
