Can the greatest common divisor and bitwise operations have anything in common? It is time to answer this question.

Suppose you are given a positive integer 𝑎
. You want to choose some integer 𝑏
 from 1
 to 𝑎−1
 inclusive in such a way that the greatest common divisor (GCD) of integers 𝑎⊕𝑏
 and 𝑎&𝑏
 is as large as possible. In other words, you'd like to compute the following function:

𝑓(𝑎)=max0<𝑏<𝑎𝑔𝑐𝑑(𝑎⊕𝑏,𝑎&𝑏).

Here ⊕
 denotes the bitwise XOR operation, and &
 denotes the bitwise AND operation.

The greatest common divisor of two integers 𝑥
 and 𝑦
 is the largest integer 𝑔
 such that both 𝑥
 and 𝑦
 are divided by 𝑔
 without remainder.

You are given 𝑞
 integers 𝑎1,𝑎2,…,𝑎𝑞
. For each of these integers compute the largest possible value of the greatest common divisor (when 𝑏
 is chosen optimally).

 ### ideas
 1. gcd(a ^ b, a & b) 最大
 2. a & b = 0 （似乎也没问题)
 3. a ^ b = 1 << x - 1 似乎是最大的？