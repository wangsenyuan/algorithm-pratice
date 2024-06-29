Nikita loves mountains and has finally decided to visit the Berlyand mountain range! The range was so beautiful that Nikita decided to capture it on a map. The map is a table of 𝑛
 rows and 𝑚
 columns, with each cell containing a non-negative integer representing the height of the mountain.

He also noticed that mountains come in two types:

With snowy caps.
Without snowy caps.
Nikita is a very pragmatic person. He wants the sum of the heights of the mountains with snowy caps to be equal to the sum of the heights of the mountains without them. He has arranged with the mayor of Berlyand, Polikarp Polikarpovich, to allow him to transform the landscape.

Nikita can perform transformations on submatrices of size 𝑘×𝑘
 as follows: he can add an integer constant 𝑐
 to the heights of the mountains within this area, but the type of the mountain remains unchanged. Nikita can choose the constant 𝑐
 independently for each transformation. Note that 𝑐
 can be negative.

Before making the transformations, Nikita asks you to find out if it is possible to achieve equality of the sums, or if it is impossible. It doesn't matter at what cost, even if the mountains turn into canyons and have negative heights.


Each test consists of several test cases. The first line contains an integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. This is followed by a description of test cases.

The first line of each test case contains three integers 𝑛,𝑚,𝑘
 (1≤𝑛,𝑚≤500,1≤𝑘≤𝑚𝑖𝑛(𝑛,𝑚)
).

The next 𝑛
 lines of each test case contain 𝑚
 integers 𝑎𝑖𝑗
 (0≤𝑎𝑖𝑗≤109
) — the initial heights of the mountains.

The next 𝑛
 binary strings of length 𝑚
 for each test case determine the type of mountain, '0
' — with snowy caps, '1
' — without them.

It is guaranteed that the sum of 𝑛⋅𝑚
 for all test cases does not exceed 250000
.


### ideas
1. 有点懵呐
2. sum(height of snowy mountain) = sum(not snowy)
3. 每次修改的都是一个k*k的正方形的区域
4. a = gcd(k, n), b = gcd(k, m)
5. 是不是每次修改的其实是不重叠的 (a, b)大小的长方形？
6. 考虑一个长条，k = 3， 那每个格子是否能够独立的变化？
7. 比如像让每个有不同的delta, want d[0], d[1], d[2], ... 
8. 使用k=3， 可以
9. (0, d[0]), 此时 x[1] += d[0], x[2] += d[0], x[3]。。。不变
10. (1, d[1] - d[0]), 此时 x[2] = x[2] + d[0] - d[0] + d[1], x[3] = x[3] + d[1] - d[0], x[4] = x[4] + d[1] - d[0]
11. (2, d[2] - d[1]), x[3] = x[3] + d[1] - d[0] + d[2] - d[1] = x[3] + d[2] - d[0], x[4] = x[4] + d[1] - d[0] + d[2] - d[1] = x[4] + d[2] - d[0]
12. 假设长度 = 4（所以，可以在0，1，2 处开始）
13. 所以，可以发现是最后的k-1个是根据前面的增长搞出来的，无法改变它们的值
14. 在上面的例子里面 x[0] + d[0], x[1] + d[1], x[2] + d[2], x[3] + d[3], x[4] + d[3] + d[2] - d[0], x[5] + d[3]
15. d[0] + d[1] + d[2] + d[3] + d[3] + d[2] - d[0] + d[3]
16. = d[1] + 2 * d[2] + 3 * d[3]
17. 大概知道了，考虑任何一个 k * k 的区域，这个区域内，snowy 和 非snowy 的高度差是 x, 如果 (total sum - x) % (k * k) = 0, 那么就是ok的
18. 否则一定是非ok的？