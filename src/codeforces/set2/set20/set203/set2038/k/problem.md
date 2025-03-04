You have an 𝑛×𝑛
 grid and two integers 𝑎
 and 𝑏
. Both the rows and the columns are numbered from 1
 to 𝑛
. Let's denote the cell at the intersection of the 𝑖
-th row and the 𝑗
-th column as (𝑖,𝑗)
.

You are standing in the cell (1,1)
 and want to move into the cell (𝑛,𝑛)
.

Suppose you are in the cell (𝑖,𝑗)
; in one step, you can move either into the cell (𝑖,𝑗+1)
 or into the cell (𝑖+1,𝑗)
 if the corresponding cells exist.

Let's define the cost of the cell (𝑖,𝑗)
 as 𝑐(𝑖,𝑗)=gcd(𝑖,𝑎)+gcd(𝑗,𝑏)
 (here, gcd(𝑥,𝑦)
 denotes the greatest common divisor of 𝑥
 and 𝑦
). The cost of the route from (1,1)
 to (𝑛,𝑛)
 is the sum of costs of the visited cells (including the starting cell and the finishing cell).


 ### ideas
 1. n 很大， 所以n*n的肯定是不行的
 2. gcd(i, a) + gcd(j, b) 
 3. 如果选择第一列，再选择最后一行
 4. n + gcd(1, a) + gcd(2, a) + ... + gcd(n - 1, a) + n * gcd(n, a) + gcd(2, b) + ... + gcd(n, b)
 5. 如果选择第一行，再选择最后一列
 6. n + gcd(1, b) + gcd(2, b) + .... + gcd(n - 1, b) + n * gcd(n, b) + gcd(2, a) + ... + gcd(n, a)
 7. (gcd(1, b) + gcd(n, a)) vs (gcd(1, a) + gcd(n, b))  
 8. (1, 1) => (i1, j1) => (i2, j2) ... => (n, n)
 9. 如果 gcd(n, a) = 1, 或者 gcd(n, b) = 1, 似乎就是显然的，gcd(n, a) = 1, 那么就先第一列，再第n行
 10. 如果这个不成立, 那是不是先找到n1 gcd(n1, a) = 1
 11. 成了贪心算法了？
 12. 差一点点。😄