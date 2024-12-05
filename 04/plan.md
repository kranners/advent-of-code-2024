## Day 4 plan

### Part 1

Since the string only ever starts with X, this is like a 7x7 grid search for any of:
```
S  S  S
 A A A
  MMM
SAMXMAS
  MMM
 A A A
S  S  S
```

To check if something is true, need to move the cursor in a direction, and see if the cursor hits another letter in the sequence
As in, to check for `XMAS`, start at the X, and move the cursor [1, 0] or +1x, +0y, then check for an M, repeat until the sequence is done.

To check all directions, we can run the same function but change the vector direction.

We don't ever need to edit the array since matches are allowed to overlap eachother.

So solution looks like:
1. Read in the file
2. Split file into a 2d array of chars
3. Iterate over all chars in the grid
4. If the char is not an X, continue
5. Otherwise, iterate over all the possible directions
6. For each direction, check if there's an XMAS in that direction
7. If so, add to the list

