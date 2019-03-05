package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	i := 0
	j := len(heights) - 1

	area := 0

	for i < j {
		w := (j - i)
		h := min(heights[i], heights[j])
		area = max(area, w*h)

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return area
}

func assert(b bool) {
	if !b {
		panic("Assertion failure")
	}
}

func main() {
	assert(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	assert(maxArea([]int{2, 3, 4, 5, 18, 17, 6}) == 17)
	assert(maxArea([]int{1, 2, 1}) == 2)
	assert(maxArea([]int{9, 6, 14, 11, 2, 2, 4, 9, 3, 8}) == 72)
}
