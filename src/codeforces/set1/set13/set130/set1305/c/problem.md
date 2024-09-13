To become the king of Codeforces, Kuroni has to solve the following problem.

He is given 𝑛
 numbers 𝑎1,𝑎2,…,𝑎𝑛
. Help Kuroni to calculate ∏1≤𝑖<𝑗≤𝑛|𝑎𝑖−𝑎𝑗|
. As result can be very big, output it modulo 𝑚
.

If you are not familiar with short notation, ∏1≤𝑖<𝑗≤𝑛|𝑎𝑖−𝑎𝑗|
 is equal to |𝑎1−𝑎2|⋅|𝑎1−𝑎3|⋅
 …
 ⋅|𝑎1−𝑎𝑛|⋅|𝑎2−𝑎3|⋅|𝑎2−𝑎4|⋅
 …
 ⋅|𝑎2−𝑎𝑛|⋅
 …
 ⋅|𝑎𝑛−1−𝑎𝑛|
. In other words, this is the product of |𝑎𝑖−𝑎𝑗|
 for all 1≤𝑖<𝑗≤𝑛
.

### ideas
1. a1, a2, a3
2. 考虑a3, (a3 - a2) * (a3 - a1)
3. 这里m比较小 也就是说，只要知道有多少个数 num % m = rem, 那么
4. 根据鸽洞原理，如果a的个数，超过m+1个，那么肯定有两个同余，那么结果就等于 0