Given an integer 𝑟
, find the number of lattice points that have a Euclidean distance from (0,0)
 greater than or equal to 𝑟
 but strictly less than 𝑟+1
.

A lattice point is a point with integer coordinates. The Euclidean distance from (0,0)
 to the point (𝑥,𝑦)
 is 𝑥2+𝑦2‾‾‾‾‾‾‾√
.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases.

The only line of each test case contains a single integer 𝑟
 (1≤𝑟≤105
).

The sum of 𝑟
 over all test cases does not exceed 105
.

### ideas
1. 首先肯定不可能去每个点的计算
2. 只考虑一个区间先
3. 半径r时，它在x轴的交点时 (0, r), 在y轴时 (r, 0)
4. 我知道了，迭代y坐标，从0到r-1，然后计算最大的x坐标