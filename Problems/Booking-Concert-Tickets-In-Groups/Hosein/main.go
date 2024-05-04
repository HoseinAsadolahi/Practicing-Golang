package main

import "fmt"

func main() {
	bms := Constructor(2, 6)
	fmt.Println(bms.Scatter(2, 1))
	fmt.Println(bms.Scatter(8, 0))
}

type BookMyShow struct {
	m                int
	EmptySeats       []int
	tmp       []int
}

func Constructor(n int, m int) BookMyShow {
	bms := BookMyShow{
		EmptySeats:       make([]int, n),
		tmp:       make([]int, n),
		m:                m,
	}
	for i := 0; i < n; i++ {
		bms.EmptySeats[i] = m
	}
	return bms
}

func (bms *BookMyShow) Gather(k int, maxRow int) []int {
	for i := 0; i <= maxRow; i++ {
		if bms.EmptySeats[i] >= k {
			bms.EmptySeats[i] -= k
			bms.tmp[i] -= k
			return []int{i, bms.m - bms.EmptySeats[i] - k}
		}
	}
	return []int{}
}

func (bms *BookMyShow) Scatter(k int, maxRow int) bool {
	j := 0
	for i := 0; i <= maxRow; i++ {
		if bms.tmp[i] >= 0 {
			if bms.tmp[i] >= k {
				bms.tmp[i] -= k
				for k = 0; k <= i; k++ {
					bms.EmptySeats[k] = bms.tmp[k]
				}
				return true
			} else {
				k -= bms.EmptySeats[i]
				bms.EmptySeats[i] = 0
			}
		}
		j++
	}
	for k = 0; k <= j; k++ {
		bms.EmptySeats[k] = bms.tmp[k]
	}
	return false
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */