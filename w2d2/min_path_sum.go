package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
memo a nested array to store the min path sum from each location

start in bottom right
min path is current value

next level
min path is current value plus the smaller value between down and right
store min path in memo

stop when you reach top right
*/
