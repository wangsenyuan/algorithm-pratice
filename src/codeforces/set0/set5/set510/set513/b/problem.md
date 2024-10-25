You are given a permutation p of numbers 1, 2, ...,  n . Let f ( p ) denote the following sum:


Find the lexicographically m -th permutation of length n that has the maximum possible value f ( p ) .


### ideas
1. 先考虑 最大的f(p)是多少？
2. 对于给定的permuation, f(p)等于所有子串最小值的和
3. 那么考虑n的贡献,显然它只能贡献1次 = n
4. 考虑n-1, 它只有和 n连接在一起的时候，才能贡献 2次 （n - 1, n)
5. 考虑n-2, 它只有和n-1（n）靠近的时候，才能贡献3 次 (n - 2, n - 1, n) 或者 (n - 2, n, n - 1) 
6. 但是它不能到n 和 n - 1 的中间， 会破坏n-1的贡献
7. 所以f(p) = [1, 2, ... n]的结果 = 1 * n + 2 * (n - 1) + ... + n * 1
8. 但是 [1....n] 是第一个这样的，1..n-2,n, n-1 也是 第二个
9. 1...n-2,n,n-1,n-3 也是满足条件的 第三个
10. 考虑 i...n （它们是连续的，怎么排列不去考虑，但它们可以提供最大的贡献）
11. 现在处理i-1, 如果它放到后面（这个序列的后面，会增加位置）排在前面，不改变位置
12. 增加了多少位置呢？
13. cnt(n) = pow(2, n - 1) 
14. cnt(1) = 1
15. cnt(2) = 2  (1, 2) (2, 1)
16. cnt(3) = 考虑把 (2, 3)排好的情况下，吧 1放在两头 2 * cnt(2)
17. ..
18. cnt(i) = 2 * cnt(i-1)
19. m = 只需要排最后的几个即可
20. 然后就是把如果 m(i) = 1, 把小的排在右边（这样子会增加当前的位置）