Kamilka has a flock of 𝑛
 sheep, the 𝑖
-th of which has a beauty level of 𝑎𝑖
. All 𝑎𝑖
 are distinct. Morning has come, which means they need to be fed. Kamilka can choose a non-negative integer 𝑑
 and give each sheep 𝑑
 bunches of grass. After that, the beauty level of each sheep increases by 𝑑
.

In the evening, Kamilka must choose exactly two sheep and take them to the mountains. If the beauty levels of these two sheep are 𝑥
 and 𝑦
 (after they have been fed), then Kamilka's pleasure from the walk is equal to gcd(𝑥,𝑦)
, where gcd(𝑥,𝑦)
 denotes the greatest common divisor (GCD) of integers 𝑥
 and 𝑦
.

The task is to find the maximum possible pleasure that Kamilka can get from the walk.

### ideas
1. gcd(a[i] + d, a[j] + d) 最大
2. gcd(a, b) = gcd(a - b, b)
3. gcd(a[i] - a[j], a[j] + d)
4. 