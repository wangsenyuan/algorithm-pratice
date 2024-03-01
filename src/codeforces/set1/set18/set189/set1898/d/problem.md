Kirill has two integer arrays 𝑎1,𝑎2,…,𝑎𝑛
and 𝑏1,𝑏2,…,𝑏𝑛
of length 𝑛
. He defines the absolute beauty of the array 𝑏
as
∑𝑖=1𝑛|𝑎𝑖−𝑏𝑖|.
Here, |𝑥|
denotes the absolute value of 𝑥
.

Kirill can perform the following operation at most once:

select two indices 𝑖
and 𝑗
(1≤𝑖<𝑗≤𝑛
) and swap the values of 𝑏𝑖
and 𝑏𝑗
.
Help him find the maximum possible absolute beauty of the array 𝑏
after performing at most one swap.

### thoughts

1. 假设swap了(i, j), 分情况讨论
2. 假设在swap前的结果位ans
3. a[i] > b[i], a[j] > b[j], a[i] > b[j], a[j] > a[i]
3. 交换后 tmp = ans - (a[j] - b[j]) - (a[i] - b[i]) + abs(a[i] - b[j]) + abs(a[j] - b[i])
4.           = ans  - a[i] + b[j] - a[i] + b[i] + a[i] - b[j] + a[j] - b[i] = ans
5. 如果 a[i] < b[j], a[j] < a[i]

### solutionIf 𝑎𝑖>𝑏𝑖

, swap them. Imagine the pairs (𝑎𝑖,𝑏𝑖)
as intervals. How can we visualize the problem?

The pair (𝑎𝑖,𝑏𝑖)
represents some interval, and |𝑎𝑖−𝑏𝑖|
is its length. Let us try to maximize the sum of the intervals' lengths. We present three cases of what a swap does to
two arbitrary intervals.

Notice how the sum of the lengths increases only in the first case. We want the distance between the first interval's
right endpoint and the second interval's left endpoint to be as large as possible. So we choose integers 𝑖
and 𝑗
such that 𝑖≠𝑗
and 𝑎𝑗−𝑏𝑖
is maximized. We add twice the value, if it is positive, to the original absolute beauty. If the value is 0
or negative, we simply do nothing.

To quickly find the maximum of 𝑎𝑗−𝑏𝑖
over all 𝑖
and 𝑗
, we find the maximum of 𝑎1,𝑎2,…𝑎𝑛
and the minimum of 𝑏1,𝑏2,…𝑏𝑛
. Subtracting the two extremum values produces the desired result.
