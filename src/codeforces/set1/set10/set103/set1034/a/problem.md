Mr. F has 𝑛
positive integers, 𝑎1,𝑎2,…,𝑎𝑛
.

He thinks the greatest common divisor of these integers is too small. So he wants to enlarge it by removing some of the
integers.

But this problem is too simple for him, so he does not want to do it by himself. If you help him, he will give you some
scores in reward.

Your task is to calculate the minimum number of integers you need to remove so that the greatest common divisor of the
remaining integers is bigger than that of all integers.

Input
The first line contains an integer 𝑛
(2≤𝑛≤3⋅105
) — the number of integers Mr. F has.

The second line contains 𝑛
integers, 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤1.5⋅107
).

Output
Print an integer — the minimum number of integers you need to remove so that the greatest common divisor of the
remaining integers is bigger than that of all integers.

You should not remove all of the integers.

If there is no solution, print «-1» (without quotes).

### ideas

1. 找出所有数字的lpf ?
2. g = 所有数字的公共质因数的乘积
3. 假设有个质因数x，它出现的最小的次数是a，第二小的次数是b，那么只要把那些出现次数是a的数字删除调
4. 新的g中至少包含b个x