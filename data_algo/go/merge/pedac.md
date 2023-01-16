input nested array
output nested array
  - any overlapping intervals are joined

Examples
merge([][]int{{1,3}, {2,6}, {8,10}, {15,18}})
merge([][]int{{1,4}, {4,5}})  // {{1,5}})
merge([][]int{{1,3}, {4,5}})  // {{1,3}, {4,5}})

?
merge([][]int{{1,7}, {4,5}}) // {{1,7}}
sorted?
can only merge two adjacent intervals?
Edge case
  end interval


Sort
have first interval
Iterate through arrays
  - iterate starting with second interval
  - use length of array
  - stop when index is greater then one minus length of array

  next interval is interval at current index
  check if end of first interval is greater then the start of the second interval
    if it is not
      add first interval to results array
      set first interval to current interval
      if current index is last index then add second interval to results array
    if it is
      encloses second interval?

      set end of first interval to the end of second interval

      last interval?

Agorithm
if the length is one just return the array
Sort the arrays
firstI is first interval in array of arrays
iterate through the other arrays
  - start at second interval
  - end at last interval

  nextI is interval at current index

  is end of firstI greater then the start of the next interval
  Yes:
    is end of firstI greater then end of nextI?
    Yes: - firstI includes all of nextI
      firstI stays the same
    No: - the two intervals overlap
      firstI end is the nextI end
    
    if index at end of array add firstI to results array
  
  No: - no overlap
    add firstI to results array
    set firstI to nextI
    
    if index at end of array add nextI to results array
return results array


Big O
  runtime NlogN from sorting
  memory N same as original
