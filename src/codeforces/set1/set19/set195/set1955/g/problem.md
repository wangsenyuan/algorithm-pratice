Not long ago, Egor learned about the Euclidean algorithm for finding the greatest common divisor of two numbers. The greatest common divisor of two numbers 𝑎
 and 𝑏
 is the largest number that divides both 𝑎
 and 𝑏
 without leaving a remainder. With this knowledge, Egor can solve a problem that he once couldn't.

Vasily has a grid with 𝑛
 rows and 𝑚
 columns, and the integer 𝑎𝑖𝑗
 is located at the intersection of the 𝑖
-th row and the 𝑗
-th column. Egor wants to go from the top left corner (at the intersection of the first row and the first column) to the bottom right corner (at the intersection of the last row and the last column) and find the greatest common divisor of all the numbers along the path. He is only allowed to move down and to the right. Egor has written down several paths and obtained different GCD values. He became interested in finding the maximum possible GCD.

Unfortunately, Egor is tired of calculating GCDs, so he asks for your help in finding the maximum GCD of the integers along the path from the top left corner to the bottom right corner of the grid.


### ideas
1. dp[i][j] 表示到(i, j)的最大gcd, 是一个集合，不是单个值，那么到下一个位置，可以算出新的gcd的集合
2. 这样子，需要考虑这个集合的大小，这个集合的大小似乎是logn
3. 首先这个集合都是 a[i][j]的因子，且相互之间不能整除。=> 每个因子至少有一个其他因子没有的质因子
4. 但是似乎还是有点问题，因为要保证这个集合是独立的条件，必须有一个重整的过程，而这个重整似乎是 o(lgn) * o(lgn)的
5. gcd(first, second) 然后检查每个因子就可以了。