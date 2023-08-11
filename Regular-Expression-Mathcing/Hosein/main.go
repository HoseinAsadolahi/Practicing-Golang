package main

import (
	"fmt"
)

func main() {
	s, p := "a", "ab*a"
	fmt.Print(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	return recursiveMatch(0, 0, len(s), len(p), s, p)
}

/*
recursiveMatch gets i, j, n, m. i is index of s and j is index of p
and n is len(s) and m is len(p)
if we hit a . character we go on recursively
if we hit the correct char we still go recursively
if we hit a char that do not math with the pattern we return false and all the recursive calls return false
*/
func recursiveMatch(i, j, n, m int, s, p string) bool {
	// if j and i hit the boundary together then it returns true
	if i >= n && j >= m {
		return true
	}

	// now they didn't hit the bound together
	// if j hit the boundary then it returns false
	if j >= m {
		return false
	}

	if j == m-1 {
		if i >= n {
			return false
		}
		if p[j] == '.' {
			return recursiveMatch(i+1, j+1, n, m, s, p)
		}
		if p[j] == s[i] {
			return recursiveMatch(i+1, j+1, n, m, s, p)
		}
		return false
	}
	if p[j+1] == '*' {
		if p[j] == '.' {
			if !recursiveMatch(i, j+2, n, m, s, p) {
				if i >= n {
					return false
				}
				return recursiveMatch(i+1, j, n, m, s, p)
			}
			return true
		} else {
			if !recursiveMatch(i, j+2, n, m, s, p) {
				if i >= n {
					return false
				}
				if p[j] == s[i] {
					return recursiveMatch(i+1, j, n, m, s, p)
				}
				return recursiveMatch(i, j+2, n, m, s, p)
			}
			return true
		}
	} else {
		if i >= n {
			return false
		}
		if p[j] == '.' {
			return recursiveMatch(i+1, j+1, n, m, s, p)
		}
		if p[j] == s[i] {
			return recursiveMatch(i+1, j+1, n, m, s, p)
		}
		return false
	}
}

// func isMatch(s string, p string) bool {
// 	var index int
// 	for i, v := range p {
// 		if index >= len(s) {
// 			return false
// 		}
// 		if i == len(p)-1 {
// 			if v != '*' {
// 				if byte(v) != s[index] {
// 					return false
// 				}
// 			}
// 		} else {
// 			if p[i+1] != '*' {
// 				if s[index] != byte(v) {
// 					return false
// 				}
// 				index++
// 			} else {
// 				for ; index < len(s); index++ {
// 					if s[index] != byte(v) {
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return index == len(s)
// }
