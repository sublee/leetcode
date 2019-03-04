package leetcode

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

			// skip if pattern is shorter than input
			if j == len(p) {
				continue
			}

			beforeStar := j+1 != len(p) && p[j+1] == '*'

			if beforeStar {
				current = append(current, j+2)
			}

			if i != len(s) && (p[j] == '.' || s[i] == p[j]) {
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
