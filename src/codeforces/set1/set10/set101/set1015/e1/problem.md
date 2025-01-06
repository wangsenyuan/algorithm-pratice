A star is a figure of the following type: an asterisk character '*' in the center of the figure and four rays (to the left, right, top, bottom) of the same positive length. The size of a star is the length of its rays. The size of a star must be a positive number (i.e. rays of length 0
 are not allowed).

Let's consider empty cells are denoted by '.', then the following figures are stars:

The leftmost figure is a star of size 1
, the middle figure is a star of size 2
 and the rightmost figure is a star of size 3
.
You are given a rectangular grid of size 𝑛×𝑚
 consisting only of asterisks '*' and periods (dots) '.'. Rows are numbered from 1
 to 𝑛
, columns are numbered from 1
 to 𝑚
. Your task is to draw this grid using any number of stars or find out that it is impossible. Stars can intersect, overlap or even coincide with each other. The number of stars in the output can't exceed 𝑛⋅𝑚
. Each star should be completely inside the grid. You can use stars of same and arbitrary sizes.

In this problem, you do not need to minimize the number of stars. Just find any way to draw the given grid with at most 𝑛⋅𝑚
 stars.

 ### ideas
 1. 可以在任何地方，如果可以draw，就draw即可
 2. 最后判断那些star是否被cover住了