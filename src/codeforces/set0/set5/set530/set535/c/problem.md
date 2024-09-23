Each Karafs has a positive integer height. Tavas has an infinite 1-based sequence of Karafses. The height of the i-th Karafs is si = A + (i - 1) × B.

For a given m, let's define an m-bite operation as decreasing the height of at most m distinct not eaten Karafses by 1. Karafs is considered as eaten when its height becomes zero.

Now SaDDas asks you n queries. In each query he gives you numbers l, t and m and you should find the largest number r such that l ≤ r and sequence sl, sl + 1, ..., sr can be eaten by performing m-bite no more than t times or print -1 if there is no such number r.


### ideas
1. 对于sl = A + (l - 1) * B
2. 要在t秒内，能够吃完的，最长的[l...r]序列，且每次只能最多吃掉m个（不需要连续）
3. 给定一个序列，如何最快的把它们都吃完呢？
4. sum = sl + ... + sr = (r - l * 1) * A + B * (l + r - 1) * (r - l + 1) / 2
5. ceil(sum / m) <= t ?
6. 似乎不大对。 假设A = 1, B = 1, 且l = 1, r = 10, m = 10
7.  sum = 10 +  (9) * 10 / 2 = 55
8.  x = 55 / 10 = 6
9.  但是，答案却是10。 数学规律的这种，好麻烦
10. 如果这个序列足够的长，远超过m，那么确实可以达到这个最优解（只要长度超过 2 * m)就是可以的
11. 每次可以选择最高的m个，然后必然会达到某个阶段，产生两个(m)高的，然后是3个, 如果有2 * m个数，那么就有办法最终产生2 * m 个相等的
12. 但是但长度 <= m, 那么答案 = s[r]. 
13. 考虑上面的那个列子，如果 m = 9
14. 第一次，处理[2...10], 得到的是 1, 1, 2, .... 9
15. 接下来，还是需要9次，
16. 所以这个r是越大越好的？