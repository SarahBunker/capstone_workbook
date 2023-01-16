package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	m2 := merge([][]int{{1, 4}, {4, 5}})         // {{1,5}})
	m3 := merge([][]int{{1, 3}, {4, 5}})         // {{1,3}, {4,5}})
	m4 := merge([][]int{{1, 7}, {4, 5}})         // {{1,7}}  can enclose
	m5 := merge([][]int{{1, 7}, {4, 5}, {5, 6}}) // {{1,7}}  can enclose can overlap more then one
	m6 := merge([][]int{{1, 4}, {-4, 0}})        // {{1,5}})

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println(m5)
	fmt.Println(m6)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	results := [][]int{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	firstI := intervals[0]
	endIndex := len(intervals) - 1
	for i := 1; i <= endIndex; i++ {
		nextI := intervals[i]
		if firstI[1] >= nextI[0] {
			if firstI[1] < nextI[1] {
				firstI[1] = nextI[1]
			}

			if i == endIndex {
				results = append(results, firstI)
			}

			continue
		}

		results = append(results, firstI)
		firstI = nextI

		if i == endIndex {
			results = append(results, nextI)
		}

	}
	return results
}
