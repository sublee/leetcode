package main

func isMatch(s string, p string) bool {
	return false
}

func assert(b bool) {
	if !b {
		panic("Assertion failure")
	}
}

func main() {
	assert(isMatch("aa", "a") == false)
	assert(isMatch("aa", "a*") == true)
	assert(isMatch("ab", ".*") == true)
	assert(isMatch("aab", "c*a*b") == true)
	assert(isMatch("mississippi", "mis*is*p*.") == false)
}
