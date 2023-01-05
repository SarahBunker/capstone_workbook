Merge Linked List
  pointer to dummy to hold head of list
  pointer to current starts at dummy
  Two pointers to hold each list
  compare values at two pointers
    have current point to the node with the smaller value
    then update the head (L1 or L2 ) to point to the next node in the list
      L1 = l1.next or l2 = l2.next
  moving pointer adds current node to list

  end condition one of the lists points to nil
    have current point to L2
    return the node that dummy points to

  Use dummy node to point to head
  or could have a variable head that points to node

Reverse Linked List
  pointer to current
  pointer to next
  pointer to nextNext
  set next to point to current
  update current to next
  update next to nextnext
  find next, next and set to nextnext
  
  stop when reach end of linked list

Reverse Linked List II
  
  iterate through linked list to find the node with the value given as left
  store pointer to node before left as start
  store pointer to left << previous head
    using code from previous problem
    reverse the linked list until you get to the node that points to the node with value right
  store pointer to right
  store pointer to node after right as end
  assign node before left to point to right
  assign previous node to point to end

Remove Duplicates from Sorted List II
  Use Hash to store values of nodes and count of values
  anchor points to last unique value
  runner iterates through linked list
  nextVal points to the node with the nextVal

  runner goes to new node
    adds value of node to hash
    if value of node is different then value of previous node
      if count is 1
        set anchor to point to next val
      move anchor to point to nextval
      move next val to point to runner
  start with dummy head?
  pointers to previous node
    