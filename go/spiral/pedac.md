Algorithm
  initialize an object to store the min and max index for rows and columns
  create methods for increase and decrease for rows and columns
    method iterates through matrix keeping either row or column constant
      adds element at current row and index to results slice
    after iterating through update the boundaries
  
  create a while loop that calls the four methods in order
    increase along row
      row is constant at minI
      iterate from minJ to max J
    increase along column
    decreaase along row
    decrease along column

    stop when min == max for row or column

  return results slice

