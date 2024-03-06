https://codeforces.com/problemset/problem/1889/C1

### problem

Doremy lives in a rainy country consisting of 𝑛
cities numbered from 1
to 𝑛
.

The weather broadcast predicted the distribution of rain in the next 𝑚
days. In the 𝑖
-th day, it will rain in the cities in the interval [𝑙𝑖,𝑟𝑖]
. A city is called dry if it will never rain in that city in the next 𝑚
days.

It turns out that Doremy has a special power. She can choose 𝑘
days (in the easy version, 𝑘=2
), and during these days it will not rain. Doremy wants to calculate the maximum number of dry cities after using the
special power.

### ideas

1. 对于每个city，可以知道它下雨的日期
2. 然后找出那些下雨天数 <= 2的
3. 然后考虑dp[i] 表示消除它下雨的天数时，能够和它一起不下雨的最左边的下标
4. 也就是要左边的下雨天数，要包含在它的下雨天数中
5. fp[i]表示右边的
