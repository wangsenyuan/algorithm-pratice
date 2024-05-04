You are given an array a consisting of n integers, and additionally an integer m. You have to choose some sequence of
indices b1, b2, ..., bk (1 ≤ b1<b2<...<bk ≤ n) in such a way that the value of is maximized. Chosen sequence can be
empty.

Print the maximum possible value of .

### ideas

1. 难道不是全部选择吗？
2. 不是的， 比如 m = 4, a = 5 2 4 1
3. 如果全部选择 (5 + 2 + 4 + 1) % 4 = 0
4. 但是选择 (5, 2) 或者(1, 2) => 3
5. n <= 35
6. n/2 <= 18
7. 先求出前半部分的全部的，后半部分全部
8. 假设分别是x, y
9. 然后就是找出是否有超过max(x, y)的