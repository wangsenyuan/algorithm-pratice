You are given an integer array 𝑎1,𝑎2,…,𝑎𝑛
.

The array 𝑏
is called to be a subsequence of 𝑎
if it is possible to remove some elements from 𝑎
to get 𝑏
.

Array 𝑏1,𝑏2,…,𝑏𝑘
is called to be good if it is not empty and for every 𝑖
(1≤𝑖≤𝑘
) 𝑏𝑖
is divisible by 𝑖
.

Find the number of good subsequences in 𝑎
modulo 109+7
.

Two subsequences are considered different if index sets of numbers included in them are different. That is, the values
​of the elements ​do not matter in the comparison of subsequences. In particular, the array 𝑎
has exactly 2𝑛−1
different subsequences (excluding an empty subsequence).

Input
The first line contains an integer 𝑛
(1≤𝑛≤100000
) — the length of the array 𝑎
.

The next line contains integers 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤106
).

Output
Print exactly one integer — the number of good subsequences taken modulo 109+7
.

### ideas

1. 先不考虑时间复杂性， dp[i][j] 表示a[i]是b[j]的情况
2. a[i] % j = 0
3. a[i] <= 1000000，n <= 100000
4. 那么a[i]的因数x ( <= i)