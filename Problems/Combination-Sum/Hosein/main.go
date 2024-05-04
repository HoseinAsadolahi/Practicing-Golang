package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{4, 2, 8}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	out := make([][]int, 0)
	var bfs func (array []int, index, sum int)
	bfs = func (array []int, index, sum int) {
		if index >= len(candidates) {
			return
		}
		if sum > target {
			return
		}
		if sum == target {
			out = append(out, append([]int{}, array...))
			return
		}
		bfs(append(array, candidates[index]), index, sum+candidates[index])
		bfs(array, index+1, sum)
	}
	bfs(make([]int, 0), 0, 0)
	return out
}