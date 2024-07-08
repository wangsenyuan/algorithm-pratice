It is given a non-negative integer 𝑥
, the decimal representation of which contains 𝑛
 digits. You need to color each its digit in red or black, so that the number formed by the red digits is divisible by 𝐴
, and the number formed by the black digits is divisible by 𝐵
.

At least one digit must be colored in each of two colors. Consider, the count of digits colored in red is 𝑟
 and the count of digits colored in black is 𝑏
. Among all possible colorings of the given number 𝑥
, you need to output any such that the value of |𝑟−𝑏|
 is the minimum possible.

Note that the number 𝑥
 and the numbers formed by digits of each color, may contain leading zeros.

Example of painting a number for 𝐴=3
 and 𝐵=13
The figure above shows an example of painting the number 𝑥=02165
 of 𝑛=5
 digits for 𝐴=3
 and 𝐵=13
. The red digits form the number 015
, which is divisible by 3
, and the black ones — 26
, which is divisible by 13
. Note that the absolute value of the difference between the counts of red and black digits is 1
, it is impossible to achieve a smaller value.

Input
The first line contains one integer 𝑡
 (1≤𝑡≤10
) — the number of test cases. Then 𝑡
 test cases follow.

Each test case consists of two lines. The first line contains three integers 𝑛
, 𝐴
, 𝐵
 (2≤𝑛≤40
, 1≤𝐴,𝐵≤40
). The second line contains a non-negative integer 𝑥
 containing exactly 𝑛
 digits and probably containing leading zeroes.


 ### ideas
 1. r和b要足够接近
 2. A, B <= 40
 3. len(x) <= 40
 4. dp[ra][rb] = 最小差距 (到目前为止， 获取到对a求余 = ra, % b = rb时的最小差距)
 5. dp[(ra + num) % a][rb] = dp[ra][rb] + 1
 6. 还是得吧count放进去 