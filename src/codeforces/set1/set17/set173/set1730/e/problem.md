You are given an array 𝑎1,𝑎2,…,𝑎𝑛
of positive integers.

Find the number of pairs of indices (𝑙,𝑟)
, where 1≤𝑙≤𝑟≤𝑛
, that pass the check. The check is performed in the following manner:

The minimum and maximum numbers are found among 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
.
The check is passed if the maximum number is divisible by the minimum number.

### thoughts

1. 对于某个下标i, 如果以它为最小值，它能够覆盖的区域为 [li, ri]
2. 如果知道 l[i] 表示左边的下标j, 且a[j] % a[i] = 0, 且 a[j] = max(a[i...j]), a[i] = min(a[i...j])

### solution

Let's introduce some new variables: 𝑙𝑔𝑒𝑖

- the position of the nearest left greater or equal than 𝑎𝑖
  (−1
  if there is none). 𝑟𝑔𝑖
- position of the nearest right greater than 𝑎𝑖
  (𝑛
  if there is none). 𝑙𝑙𝑖
- position of the nearest left lower than 𝑎𝑖
  (−1
  if there is none). 𝑟𝑙𝑖
- position of the nearest right lower than 𝑎𝑖
  (𝑛
  if there is none). All this can be calculated, for example, using a stack in 𝑂(𝑛)
  or using binary search and sparse table in 𝑂(𝑛log𝑛)
  Let's iterate over the position 𝑖
  of the leftmost maximum of the good segment. Then the 𝑖
  -th element will be the maximum on the segment [l, r] if 𝑙𝑔𝑒𝑖<𝑙≤𝑖
  and 𝑖≤𝑟<𝑟𝑔𝑖
  (1)
  . For the segment to pass the test, the minimum must be a divisor of the maximum. Let's iterate over this divisor 𝑑
  and find the number of segments where the maximum is 𝑎𝑖
  and the minimum is 𝑑
  . Consider positions of occurrence of 𝑑
  𝑗1
  and 𝑗2
  — the nearest left and right to 𝑖
  (they can be found using two pointers). Let's find the number of segments satisfying the condition (1)
  , in which the element 𝑗1
  is a minimum. To do this, similar conditions must be added to (1)
  : 𝑙𝑙𝑖<𝑙≤𝑗1
  and 𝑗1≤𝑟<𝑟𝑔𝑖
  . Intersecting these conditions, we obtain independent segments of admissible values of the left and right boundaries
  of the desired segment. Multiplying their lengths, we get the number of required segments. Similarly, the number of
  segments satisfying (1)
  , in which 𝑗2
  is the minimum, is found, but in order not to count 2 times one segment, one more condition must be added: 𝑗1<𝑙
  . The sum of these quantities over all 𝑖
  and divisors of 𝑎𝑖
  will give the answer to the problem.

To enumerate divisors, it is better to precompute the divisors of all numbers in 𝑂(𝐴log𝐴)
, where 𝐴
is the constraint on 𝑎𝑖
. So the whole solution runs in 𝑂(𝐴log𝐴+𝑛𝐷)
, where 𝐷
is the maximum number of divisors.