Bob and Alice are often participating in various programming competitions. Like many competitive programmers, Alice and Bob have good and bad days. They noticed, that their lucky and unlucky days are repeating with some period. For example, for Alice days [𝑙𝑎;𝑟𝑎]
 are lucky, then there are some unlucky days: [𝑟𝑎+1;𝑙𝑎+𝑡𝑎−1]
, and then there are lucky days again: [𝑙𝑎+𝑡𝑎;𝑟𝑎+𝑡𝑎]
 and so on. In other words, the day is lucky for Alice if it lies in the segment [𝑙𝑎+𝑘𝑡𝑎;𝑟𝑎+𝑘𝑡𝑎]
 for some non-negative integer 𝑘
.

The Bob's lucky day have similar structure, however the parameters of his sequence are different: 𝑙𝑏
, 𝑟𝑏
, 𝑡𝑏
. So a day is a lucky for Bob if it lies in a segment [𝑙𝑏+𝑘𝑡𝑏;𝑟𝑏+𝑘𝑡𝑏]
, for some non-negative integer 𝑘
.

Alice and Bob want to participate in team competitions together and so they want to find out what is the largest possible number of consecutive days, which are lucky for both Alice and Bob.


### ideas
1. [l + k * t, r + k * t], when d in the range it is 
2. d >= l + k * t => k = (d - l) / t
3. 这里考虑2种情况， 就是一种是， 第一个区间能够完全嵌入第二个区间
4. 第二个是它们有重叠
5. 完全嵌入，要满足 (r1 - l1 + 1) <= (r2 - l2 + 1)
6. 且开始的时间肯定是l1 + ? * t1
7. l1 + k1 * t1 >= l2 + k2 * t2
8. r1 + k1 * t2 <= r2 + k2 * t2
9. k2 * t2 <= l1 + k1 * t1 - l2
10. k2 * t2 >= r1 + k1 * t1 - r2
11. r1 + k1 * t1 - r2 <= l1 + k1 * t1 - l2
12. r1 - l1 <= r2 - l2 (回到原点了)
13. 是不是说，只要满足这个条件，就一定能得到这个最小区间？
14. 先不考虑l, 只考虑t，两个t，不断的成倍增长，那么经过多久，它们可以相等呢？
15. L = lcm(t1, t2)
16. 那么经过 k1 = L / t1, k2 = L / t2
17. 那么它们就可以相同
18. 但是此时，它们只是相当于回到了原点，它们的距离 = l2 - l1
19. 这里我们的目的是为了减少 l2-l1 (假设它们相等了，那么答案就是更短的那个区间)
20. 哪似乎，就可以不断的尝试了(不行， 因为 L 有可能 = t1 * t2)
21. 什么情况下，这两个会很接近呢？
22. d1 = l1 + k1 * t1
23. d2 = l2 + k2 * t2
24. 目的是 d1 和 d2 接近
25. 考虑在没有l的情况下， k1 * t1, 和 k2 * t2 的差是怎么变化的
26. let d = k1 * t1 = dx + k2 * t2 
27. d - dx < t2 
28. dx = (k1 * t1) % (k2 * t2)
29. k2 = k1 * t1 / t2
30. let g = gcd(t1, t2) 
31. 貌似它们可以按照这个差额不断的变化, 假设，当前的差额是 l = abs(l1 - l2)， 那么 l % g = 最左边的那个差额

### solution

All possible shifts of Alice's and Bobs' pattern periods are the multiples of gcd(𝑡𝑎,𝑡𝑏)
. You want to use such a shift that the start of their lucky days is as close as possible. Also if they can't coincide precisely, you need to try shifting in such way that Alice's lucky days start before Bob's, and the opposite.