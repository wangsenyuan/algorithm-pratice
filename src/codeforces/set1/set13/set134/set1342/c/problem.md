You are given two integers 𝑎
 and 𝑏
, and 𝑞
 queries. The 𝑖
-th query consists of two numbers 𝑙𝑖
 and 𝑟𝑖
, and the answer to it is the number of integers 𝑥
 such that 𝑙𝑖≤𝑥≤𝑟𝑖
, and ((𝑥mod𝑎)mod𝑏)≠((𝑥mod𝑏)mod𝑎)
. Calculate the answer for each query.

Recall that 𝑦mod𝑧
 is the remainder of the division of 𝑦
 by 𝑧
. For example, 5mod3=2
, 7mod8=7
, 9mod4=1
, 9mod9=0
.

### ideas
1. a, b比较小， 不超过 200
2. 如果 a = b, 那么就不存在这样的x，答案是0
3. 不妨设a < b
4. x % a = m, x % b = n
5. m % b != n % a
6. 试试计算 m % b = n % a
7. m % b => m = w * b + r
8. n % a => n = v * a + r
9. 如果x是a的倍数，但不是b的倍数，有可能是成立的，因为左边肯定是0，但右边如果 x % b 不是a的倍数，就不是0
10. 如果 a 既不是a的倍数，也不是b的倍数，不是一定成立呢？
11. 