// To execute Go code, please declare a func main() in a package "main"

/*
Given a string s, return true IF the s can be a palindrome after deleting at MOST one character from it.


Example 1: No deletion needed
Input: s = "aba"
Output: true

Example 2: Delete either b or c
Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.

Example 3: Cannot be palidrome even with deletion
Input: s = "abc"
Output: false

Example 4:
Input s:
gmlupuupuculmg
   s
				   e

 s
               e
"abbcbgba"


abbgcbgba
  s   e

while start <= end
  if str[start] == str[end]
    start++
    end--
  else
    check if is palindromic from start+1 to end
    check if is palindromic from end-1 to start


*/

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
	}
}

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return checkForPalindrome(s, start+1, end) || checkForPalindrome(s, start, end-1)
		}
	}
	return true
}

func checkForPalindrome(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
