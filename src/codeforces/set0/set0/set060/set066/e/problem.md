Little Vasya's uncle is a postman. The post offices are located on one circular road. Besides, each post office has its own gas station located next to it. Petya's uncle works as follows: in the morning he should leave the house and go to some post office. In the office he receives a portion of letters and a car. Then he must drive in the given car exactly one round along the circular road and return to the starting post office (the uncle can drive along the circle in any direction, counterclockwise or clockwise). Besides, since the car belongs to the city post, it should also be fuelled with gasoline only at the Post Office stations.

The total number of stations equals to n. One can fuel the car at the i-th station with no more than ai liters of gasoline. Besides, one can fuel the car no more than once at each station. Also, the distance between the 1-st and the 2-nd station is b1 kilometers, the distance between the 2-nd and the 3-rd one is b2 kilometers, ..., between the (n - 1)-th and the n-th ones the distance is bn - 1 kilometers and between the n-th and the 1-st one the distance is bn kilometers. Petya's uncle's high-tech car uses only one liter of gasoline per kilometer. It is known that the stations are located so that the sum of all ai is equal to the sum of all bi. The i-th gas station and i-th post office are very close, so the distance between them is 0 kilometers.

Thus, it becomes clear that if we start from some post offices, then it is not always possible to drive one round along a circular road. The uncle faces the following problem: to what stations can he go in the morning to be able to ride exactly one circle along the circular road and visit all the post offices that are on it?

Petya, who used to attend programming classes, has volunteered to help his uncle, but his knowledge turned out to be not enough, so he asks you to help him write the program that will solve the posed problem.

### ideas
1. 假设从1号开始，到达第i个station，那么它们之间的距离 x = sum(b1 + b2... + b[i-1])
2. 添加的油 y = sum(a1 + a2 + ... a[i-1])
3. 如果 x > y 那么从1号，就不能到达位置i
4. 也就是找到 sum(b[1...i-1]) > sum(a[1....i-1]) 的位置，如果存在这样的位置，那么1号就是不ok的
5. 但是这个是不能二分的
6. s1[i] = b[1] + b[2] + .. + b[i-1]
7. s2[i] = a[1] + a[2] + .. + a[i-1]
8. 也就是说要在1...n-1处，找到 s1[i] - s2[i]的最大值
9. 假设 i是1能到达的最远处，不能到达i+1
10. 那么从1移动到2的时候，b[1] <= a[1] 肯定是成立的（否则它都到不了位置2）
11. 那么从2，似乎也肯定到不了位置i+1（因为从第一个gas处加的油，本来还剩一点的，现在没有了）
12. 也就是说，从1.。。i都不能到达i+1.加下来从i+1开始判断就可以了
13. 但是反过来却不成立，也就是说，即使从1开始能够到达位置i，也不能证明从2也能到达位置i
14. s1[2..i] = b[2] + b[3] + .. b[i-1]
15. s2[2..i] = a[2] + a[3] + .. a[i-1]
16. b[2] - a[2] + b[3] - a[3] ... b[i-1] - a[i-1] > 0 的位置
17. 肯定是在某个 b[i-1] - a[i-1] 最大的地方出现？
18. 不一定，考虑它的前一个位置，a[i-2] 远大于b[i-2], 那么及时 b[i-1] - a[i-1]这里算出来是符合的
19. 也不能保证i-2前面是ok的
20. 不过，假设1能够转一圈，那么从1移动到2的时候，b[i] <= a[i], 如果这个差值 a[i] - b[i]比2...1的最高的差值，还要大，那么2..肯定转不了一圈
21. 