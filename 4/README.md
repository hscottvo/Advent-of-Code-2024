## Day 4 - Ceres Search
The problem can be found at [this link](https://adventofcode.com/2024/day/4)
### Overview
Ceres Search involves solving a word search, but the only word present in the search is multiple instances of `XMAS`. We are tasked with finding the number of `XMAS`'s in the entire puzzle.

The prompt provides the following example:
```
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
```

The solution then finds these instances of `XMAS`:
```
..XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
```
### Approach
While the problem can be viewed as a directed graph problem, the code can be much simpler(or more complicated, depending on how you look at it.).

Let each letter be a graph node. An edge `(u,v)` exists only when `v` is `u`'s surrounding neighbor (including the cardinal directions as well as diagonals), and `v` is the next letter in `XMAS` from `u`.
- If `u` has letter `A` and `v` has letter `X`, then `(u,v)` does not exist.
- For `(u,v)` to exist, `v` must be exactly 1 space away from `u` in the vertical and/or horizontal direction.

Looking through the entire puzzle, we find all instances of `X`. Starting from there, we traverse until we fail, or until we can find `XMAS` in one direction.

Whenever we visit `S` in this manner, in a straight line, we increment the counter. 

### Time and Space Complexity

#### Time Complexity
With word search of size `mxn`, this solution is in `O(mn)`. We visit each letter in the puzzle a constant amount of times. At most, a single `X` visits in 8 directions, for up to 3 nodes in each direction.

#### Space Complexity
The algorithm uses `O(1)` additional space. Any memory allocated for an instance of `X` can be deallocated immediately after use. 
