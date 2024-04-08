Ivan is playing a strange game.

He has a matrix a with n rows and m columns. Each element of the matrix is equal to either 0 or 1. Rows and columns are
1-indexed. Ivan can replace any number of ones in this matrix with zeroes. After that, his score in the game will be
calculated as follows:

Initially Ivan's score is 0;
In each column, Ivan will find the topmost 1 (that is, if the current column is j, then he will find minimum i such that
ai, j = 1). If there are no 1's in the column, this column is skipped;
Ivan will look at the next min(k, n - i + 1) elements in this column (starting from the element he found) and count the
number of 1's among these elements. This number will be added to his score.
Of course, Ivan wants to maximize his score in this strange game. Also he doesn't want to change many elements, so he
will replace the minimum possible number of ones with zeroes. Help him to determine the maximum possible score he can
get and the minimum possible number of replacements required to achieve that score.

### ideas

1. 先考虑能够获得最大分数
2. 对于每一列，需要找到该列最大的k子序列中的1的数量（从某个1开始）
3. 然后删除那个1上面的1