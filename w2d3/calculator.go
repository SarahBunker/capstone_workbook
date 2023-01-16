/*
current op = ''
current num = ""
stack empty slice
firstNum, start := nextNum(string, 0)
add firstNum to stack
iterate through string from index start
	case current rune is a space
		continue
	case current rune is an operator
		nextNum, index := nextNum(string, index + 1)
		operations(stackPtr, opRune, nextNum)

return sum of values in slice

helper functions

nextNum
inputs: string,  index where num could start
curNum = ""
iterate from index to len of string
switch curVal
  case ""
	  continue
	case integer
		add to curNum
	case operator
	  break
return curNum converted to a number and current index
--------------------
operations
inputs: ptr to stack, operator rune, curVal int
switch operator
	case '+'
		// add positive value to stack
	case '-'
		curVal *= -1
	case '*'
		preVal := pop last value from stack
		curVal *= preVal
	case '*'
		preVal := pop last value from stack
		curVal /= preVal  <<< make sure to round down
add curVal to stack
return
-------------------

*/

package main

import (
	"fmt"
	"strconv"
)

func nextNum(s string, index int) (int, int) {
	curNum := ""
	for ; index < len(s); index++ {
		curVal := s[index]
		switch curVal {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			curNum += string(curVal)
		case '+', '-', '*', '/':
			num, err := strconv.Atoi(curNum)
			if num == 0 {
				fmt.Println(err)
			}
			return num, index
		}
	}
	num, err := strconv.Atoi(curNum)
	if num == 0 {
		fmt.Println(err)
	}
	return num, index
}

func calculate(s string) int {
	curNum, index := nextNum(s, 0)
	stack := []int{curNum}
	for index < len(s) {
		curOp := s[index]
		curNum, index = nextNum(s, index+1)
		switch curOp {
		case '+':
			stack = append(stack, curNum)
		case '-':
			stack = append(stack, curNum*-1)
		case '*':
			preNum := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curNum *= preNum
			stack = append(stack, curNum)
		case '/':
			preNum := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curNum = preNum / curNum
			stack = append(stack, curNum)
		}
	}
	fmt.Println(stack)
	return sum(stack)
}

func sum(slice []int) (total int) {
	for _, v := range slice {
		total += v
	}
	return
}

func main() {
	str := "3+2*2"
	str = " 3/2 "
	str = " 3+5 / 2 "
	str = " 3*4 + 6 - 2 + 3* 4 * 6"
	output := calculate(str)
	fmt.Println(output)
}
