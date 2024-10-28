User ainta has a permutation p1, p2, ..., pn. As the New Year is coming, he wants to make his permutation as pretty as possible.

Permutation a1, a2, ..., an is prettier than permutation b1, b2, ..., bn, if and only if there exists an integer k (1 ≤ k ≤ n) where a1 = b1, a2 = b2, ..., ak - 1 = bk - 1 and ak < bk all holds.

As known, permutation p is so sensitive that it could be only modified by swapping two distinct elements. But swapping two elements is harder than you think. Given an n × n binary matrix A, user ainta can swap the values of pi and pj (1 ≤ i, j ≤ n, i ≠ j) if and only if Ai, j = 1.

Given the permutation p and the matrix A, user ainta wants to know the prettiest permutation that he can obtain.

### ideas
1. 显然，p越靠近1,2,3...n越prettier
2. 