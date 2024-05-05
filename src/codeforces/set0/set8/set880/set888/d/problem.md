A permutation p of size n is an array such that every integer from 1 to n occurs exactly once in this array.

Let's call a permutation an almost identity permutation iff there exist at least n - k indices i (1 ≤ i ≤ n) such that
pi = i.

Your task is to count the number of almost identity permutations for given numbers n and k.

### ideas

1. 换句话说，就是最多有k个位置 p[i] != i
2. 不可能有1个位置不一样，但是可以有2，3，4个位置不一样