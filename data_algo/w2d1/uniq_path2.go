// https://leetcode.com/problems/unique-paths-ii/

/*
if m is less then 0 or if n is less then 0 or if obstacle grid at m and n are equal to 1
	return zero

if at start m == 0 and n == 0
	return 1

check if value is in cache return cache value

otherwise calculate it
cache it
return it

always
- check cache
- pass cache
- return value of cache

most problems use __ for cache
  map - go
	object - js
	hash - ruby
*/