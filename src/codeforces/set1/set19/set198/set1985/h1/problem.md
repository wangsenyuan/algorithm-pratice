Easy and hard versions are actually different problems, so read statements of both problems completely and carefully. The only difference between the two versions is the operation.

Alex has a grid with 𝑛
 rows and 𝑚
 columns consisting of '.' and '#' characters. A set of '#' cells forms a connected component if from any cell in this set, it is possible to reach any other cell in this set by only moving to another cell in the set that shares a common side. The size of a connected component is the number of cells in the set.

In one operation, Alex selects any row 𝑟
 (1≤𝑟≤𝑛
) or any column 𝑐
 (1≤𝑐≤𝑚
), then sets every cell in row 𝑟
 or column 𝑐
 to be '#'. Help Alex find the maximum possible size of the largest connected component of '#' cells that he can achieve after performing the operation at most once.

Input
The first line of the input contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of each test case contains two integers 𝑛
 and 𝑚
 (1≤𝑛⋅𝑚≤106
) — the number of rows and columns of the grid.

The next 𝑛
 lines each contain 𝑚
 characters. Each character is either '.' or '#'.

It is guaranteed that the sum of 𝑛⋅𝑚
 over all test cases does not exceed 106
.

### ideas
1. 直觉上，要brute force 每行，每行，看看变化后的component size
2. dp[i][j][0]表示从上到下时，dp[i][j]能够到到的component的size
3. 对于把第r行变成#后，需要检查它下上的(i,j)[0/1]的component的情况
4. got it