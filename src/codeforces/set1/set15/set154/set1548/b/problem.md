British mathematician John Littlewood once said about Indian mathematician Srinivasa Ramanujan that "every positive integer was one of his personal friends."

It turns out that positive integers can also be friends with each other! You are given an array 𝑎
 of distinct positive integers.

Define a subarray 𝑎𝑖,𝑎𝑖+1,…,𝑎𝑗
 to be a friend group if and only if there exists an integer 𝑚≥2
 such that 𝑎𝑖mod𝑚=𝑎𝑖+1mod𝑚=…=𝑎𝑗mod𝑚
, where 𝑥mod𝑦
 denotes the remainder when 𝑥
 is divided by 𝑦
.

Your friend Gregor wants to know the size of the largest friend group in 𝑎
.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤2⋅104
).

Each test case begins with a line containing the integer 𝑛
 (1≤𝑛≤2⋅105
), the size of the array 𝑎
.

The next line contains 𝑛
 positive integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤1018
), representing the contents of the array 𝑎
. It is guaranteed that all the numbers in 𝑎
 are distinct.

It is guaranteed that the sum of 𝑛
 over all test cases is less than 2⋅105
.

### ideas
1. a % m = b % m => (a - b) % m = 0
2. 也就是相临两数差的因子。即使如此，这个差也会有1e18, 所以，因子的数量也有 1e9
3. (a, b, c) (a - b) % m = 0, (b - c) % m = 0
4. 那么m就是 (a - b, b - c) 的公因子
5. 换句话说，就是连续的差的gcd > 1