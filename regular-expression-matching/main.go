package main

func isMatch(s string, p string) bool {
	current := []int{0}
	next := []int{}

	for i := 0; ; i++ {
		dedup := make(map[int]bool)

		for k := 0; k < len(current); k++ {
			j := current[k]

			// skip already evaluated pattern
			if dedup[j] {
				continue
			}
			dedup[j] = true

			// fully-matched?
			if i == len(s) && j == len(p) {
				return true
			}

			// is pattern shorter than input?
			if j == len(p) {
				continue
			}

			beforeStar := j+1 != len(p) && p[j+1] == '*'

			if beforeStar {
				current = append(current, j+2)
			}

			// end of input?
			if i == len(s) {
				continue
			}

			if p[j] == '.' || s[i] == p[j] {
				if beforeStar {
					next = append(next, j, j+2)
				} else {
					next = append(next, j+1)
				}
			}
		}

		// any pattern not matched
		if len(next) == 0 {
			return false
		}

		current, next = next, current[:0]
	}
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
	assert(isMatch("aaba", "ab*a*c*a") == false)
	assert(isMatch("ab", ".*c") == false)
	assert(isMatch("aaa", "a*a") == true)
	assert(isMatch("aaa", "ab*a") == false)
	assert(isMatch("a", "ab*") == true)
	assert(isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*a*a*b") == true)
}
