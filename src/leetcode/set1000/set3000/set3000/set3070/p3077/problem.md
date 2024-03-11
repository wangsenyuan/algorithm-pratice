You are given a 0-indexed array of integers nums of length n, and a positive odd integer k.

The strength of x subarrays is defined as strength = sum[1] * x - sum[2] * (x - 1) + sum[3] * (x - 2) - sum[4] * (x -

3)
    + ... + sum[x] * 1 where sum[i] is the sum of the elements in the ith subarray. Formally, strength is sum of (-1)
      i+1 * sum[i] * (x - i + 1) over all i's such that 1 <= i <= x.

You need to select k disjoint
subarrays
from nums, such that their strength is maximum.

Return the maximum possible strength that can be obtained.

Note that the selected subarrays don't need to cover the entire array.

### ideas

1. 3 * a - 2 * b + c
2. a - b + c + a - b + a
3. 也不完全是prefix sum
4. dp?