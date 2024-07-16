Let's define the cost of a string 𝑠
 as the number of index pairs 𝑖
 and 𝑗
 (1≤𝑖<𝑗<|𝑠|
) such that 𝑠𝑖=𝑠𝑗
 and 𝑠𝑖+1=𝑠𝑗+1
.

You are given two positive integers 𝑛
 and 𝑘
. Among all strings with length 𝑛
 that contain only the first 𝑘
 characters of the Latin alphabet, find a string with minimum possible cost. If there are multiple such strings with minimum cost — find any of them.


 ### ideas
1. 把 s[i]s[i+1] 看成一对，那么对于字符集k，来说，共有 k * k 对
2. 如果 k * k >= n, 那么cost是0， else肯定会存在重复的
3. 就是一直放到，n % (k * k)
4. 1, 2, 3, 4, 1, 2, 3
5. 这个思路是对的，但是要构造第一个string还是比较麻烦的