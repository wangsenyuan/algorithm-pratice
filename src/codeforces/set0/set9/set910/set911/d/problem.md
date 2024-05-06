A permutation of size n is an array of size n such that each integer from 1 to n occurs exactly once in this array. An
inversion in a permutation p is a pair of indices (i, j) such that i>j and ai<aj. For example, a
permutation [4, 1, 3, 2] contains 4 inversions: (2, 1), (3, 1), (4, 1), (4, 3).

You are given a permutation a of size n and m queries to it. Each query is represented by two indices l and r denoting
that you have to reverse the segment [l, r] of the permutation. For example, if a =[1, 2, 3, 4] and a query l = 2, r = 4
is applied, then the resulting permutation is [1, 4, 3, 2].

After each query you have to determine whether the number of inversions is odd or even.

### ideas

1. reverse [l...r] 后，对它前面的，已经后面的，已经前、中、后之间的相互关系没有影响
2. 只有 l...r中间的inversion 发生了变化，
3. 假设 [l..r] 的长度 ln = r - l + 1
4. 那么总的 sum = (0 + ln - 1) * ln / 2
5. 如果原来这个区间内，inversion的个数是x, 那么翻转后的inversion = sum - old inversion
6. 但是整体考虑似乎也不大对，因为后面有可能包含这个区间的一部分
7. 如果 sum是偶数，似乎是不改变这个区间的奇偶性的，
8. 如果是奇数，就肯定改变奇偶性
9. 求的是整个区间的奇偶性