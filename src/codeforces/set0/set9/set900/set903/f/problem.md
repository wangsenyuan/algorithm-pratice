You are given a matrix f with 4 rows and n columns. Each element of the matrix is either an asterisk (*) or a dot (.).

You may perform the following operation arbitrary number of times: choose a square submatrix of f with size k × k (where 1 ≤ k ≤ 4) and replace each element of the chosen submatrix with a dot. Choosing a submatrix of size k × k costs ak coins.

What is the minimum number of coins you have to pay to replace all asterisks with dots?


### ideas
1. 记录前3列的状态，共有（2 ^^ 12)，+当前列状态，转移到新的 三列状态
2. 然后根据当前的状态，使用最多1个4，4个3，4个2，16个1，转移到新的状态
3. 