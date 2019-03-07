package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	ns := []int{}
	for x != 0 {
		n := x % 10
		ns = append(ns, n)
		x /= 10
	}

	for i := 0; i < len(ns)/2; i++ {
		if ns[i] != ns[len(ns)-i-1] {
			return false
		}
	}
	return true
}

func assert(b bool) {
	if !b {
		panic("Assertion failure")
	}
}

func main() {
	assert(isPalindrome(121) == true)
	assert(isPalindrome(1221) == true)
	assert(isPalindrome(12321) == true)
	assert(isPalindrome(-121) == false)
}
