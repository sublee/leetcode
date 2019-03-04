package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func test(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))
	assert.True(t, isMatch("aa", "a*"))
	assert.True(t, isMatch("ab", ".*"))
	assert.True(t, isMatch("aab", "c*a*b"))
	assert.False(t, isMatch("mississippi", "mis*is*p*."))
	assert.False(t, isMatch("aaba", "ab*a*c*a"))
	assert.False(t, isMatch("ab", ".*c"))
	assert.True(t, isMatch("aaa", "a*a"))
	assert.False(t, isMatch("aaa", "ab*a"))
	assert.True(t, isMatch("a", "ab*"))
	assert.True(t, isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*a*a*b"))
}
