You are given a permutation p of length n. Remove one element from permutation to make the number of records the maximum
possible.

We remind that in a sequence of numbers a1, a2, ..., ak the element ai is a record if for every integer j (1 ≤ j<i) the
following holds: aj<ai.

### ideas

1. 先理解一下这个题目，就是要存在一个最长的连续增长的前缀
2. 那岂不是删掉第一个 a[i] > a[i+1]的就可以了？
3. 好像不是的
4. 假设删掉i后，它的影响是什么，如果 a[i] > 它前面所有的数, 那么删掉i， 减少了一个
5. 如果后面存在这样的a[j] > a[i]前面的数，但是小于 a[i], 那么就是增加了这些数
6. 反过来假设对于j来说，i是它的阻挡，那么 i的贡献 + 1