You are given a 0-1 rectangular matrix. What is the number of squares in it? A square is a solid square frame (border) with linewidth equal to 1. A square should be at least 2 × 2. We are only interested in two types of squares:

squares with each side parallel to a side of the matrix;
squares with each side parallel to a diagonal of the matrix.

For example the following matrix contains only one square of the first type: 
0000000 
0111100 
0100100 
0100100 
0111100

The following matrix contains only one square of the second type:
0000000
0010000
0101000
0010000
0000000
Regardless of type, a square must contain at least one 1 and can't touch (by side or corner) any foreign 1. Of course, the lengths of the sides of each square should be equal.

How many squares are in the given matrix?

### ideas
1. 考虑水平/垂直的情况，倾斜的情况，可以摆正后重新处理
2. 在水平/垂直的情况下，可以迭代 (i, j, l) 
3. l其实也可以不用， 它必须要一直延伸到最远的地方（否则就和1靠在一起了）
4. 然后要判断，它能不能再另外一个点连在一起（可以用sum判断）