Diagonal Difference Problem
=========================

Given a square matrix, calculate the absolute difference between the sums of its two diagonals.

- The primary diagonal is the one that runs from the top-left to the bottom-right.
- The secondary diagonal is the one that runs from the top-right to the bottom-left.

For example, given the following 3x3 matrix:

11  2   4
 4  5   6
10  8 -12

The sum of the primary diagonal is:
11 + 5 + (-12) = 4

The sum of the secondary diagonal is:
4 + 5 + 10 = 19

The absolute difference is:
|4 - 19| = 15

Input:
- A single integer n, the number of rows and columns in the matrix.
- n lines, each containing n space-separated integers describing the matrix rows.

Output:
- The absolute diagonal difference as a single integer. 