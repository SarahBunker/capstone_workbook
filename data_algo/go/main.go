package main

import (
	"fmt"
)

func main() {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(data))
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	newMatrix := Matrix{data: matrix}
	newMatrix.initializeBounds()
	fmt.Println(newMatrix)
	for !newMatrix.theEnd() {
		arr := newMatrix.increaseOnRow()
		result = append(result, arr...)
		// newMatrix.increaseOnCol()
		arr = newMatrix.decreaseOnRow()
		result = append(result, arr...)
		// newMatrix.decreaseOnCol()
		fmt.Println("Hello")
		fmt.Println(result)
		break
	}
	return []int{}
}

// fix breaking from for loop
// fix adding the other functions

type Matrix struct {
	data [][]int
	minI int
	maxI int
	minJ int
	maxJ int
}

func (m *Matrix) initializeBounds() {
	m.maxI = len(m.data) - 1
	m.maxJ = len(m.data[0]) - 1
}

// type Indices struct {
// 	minI int
// 	maxI int
// 	minJ int
// 	maxJ int
// }

func (m *Matrix) theEnd() bool {
	return m.minI == m.maxI || m.minJ == m.maxJ
}

func (m *Matrix) increaseOnRow() []int {
	rowi := m.minI
	result := []int{}
	for j := m.minJ; j <= m.maxJ; j++ {
		result = append(result, m.data[rowi][j])
	}
	return result
}

func (m *Matrix) decreaseOnRow() []int {
	rowi := m.maxI
	result := []int{}
	for j := m.maxJ; j <= m.minJ; j-- {
		result = append(result, m.data[rowi][j])
	}
	return result
}
