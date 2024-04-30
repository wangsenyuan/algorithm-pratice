Bash likes playing with arrays. He has an array a1, a2, ... an of n integers. He likes to guess the greatest common
divisor (gcd) of different segments of the array. Of course, sometimes the guess is not correct. However, Bash will be
satisfied if his guess is almost correct.

Suppose he guesses that the gcd of the elements in the range [l, r] of a is x. He considers the guess to be almost
correct if he can change at most one element in the segment such that the gcd of the segment is x after making the
change. Note that when he guesses, he doesn't actually change the array — he just wonders if the gcd of the segment can
be made x. Apart from this, he also sometimes makes changes to the array itself.

Since he can't figure it out himself, Bash wants you to tell him which of his guesses are almost correct. Formally, you
have to process q queries of one of the following forms:

1 l r x — Bash guesses that the gcd of the range [l, r] is x. Report if this guess is almost correct.
2 i y — Bash sets ai to y.
Note: The array is 1-indexed.

### ideas

1. 对于seg [l...r] 如果 gcd(l...r) = x， 就是almost good的
2. 如果不相等，假设要改变中间的a[i], 不包括a[i]是的gcd是w, 改变a[i]
3. 要满足 gcd(w, ?) = x
4. 最稳妥的做法就是要a[i] 变成x
5. 那么就有 w % x = 0,
6. 显然没法迭代任何一个i，所以还需要深挖
7. gcd另外一种计算方式是，将每个数质因数分解后，看每个质因数质数的最小值
8. 所以对于x来说，它的质因数分解后，每个质因数的指数，在这个区间内，必须至少是 ln - 1
9. 如果都是ln，那么 ok, 如果是 其中是某一部分是 ln - 1那么就需要找出，那个不满足条件的i
10. 然后就可以回到上面的部分
11. 上面的不大对， 很难维护质因数的指数数组
12. 用线段树， tr[l...r] = 这个区间内gcd
13. 如果 gcd % x == 0, 把这个区间的长度加上去就可以了
14. 当gcd % x != 0 时，必须要谨慎处理，如果它的两个children也不满足，直接返回0即可
15. 否则至少要一个满足