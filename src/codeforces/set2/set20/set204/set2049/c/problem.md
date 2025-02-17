You are given integers 𝑛
, 𝑥
, and 𝑦
. There are 𝑛
 dragons sitting in a circle. The dragons are numbered 1,2,…,𝑛
. For each 𝑖
 (1≤𝑖≤𝑛
), dragon 𝑖
 is friends with dragon 𝑖−1
 and 𝑖+1
, where dragon 0
 is defined to be dragon 𝑛
 and dragon 𝑛+1
 is defined to be dragon 1
. Additionally, dragons 𝑥
 and 𝑦
 are friends with each other (if they are already friends, this changes nothing). Note that all friendships are mutual.

Output 𝑛
 non-negative integers 𝑎1,𝑎2,…,𝑎𝑛
 such that for each dragon 𝑖
 (1≤𝑖≤𝑛
), the following holds:

Let 𝑓1,𝑓2,…,𝑓𝑘
 be the friends of dragon 𝑖
. Then 𝑎𝑖=mex(𝑎𝑓1,𝑎𝑓2,…,𝑎𝑓𝑘)
.∗

### ideas
1. 先不考虑(x, y)的影响，每个dragon只和临近的有关系。 0, 1, 2, 3.... n - 1
2. a[1] = mex(a[0], a[2])
3. a[2] = mex(a[1], a[3])
4. ...
5. a[i] = mex(a[i-1], a[i+1])
6. a[0] = mex(a[n-1], a[1])
7. 如果n是偶数，比如4， 那么 0, 1, 0, 1 是一个有效的序列
8. 如果n是奇数, 比如5， 0, 1, 0, 1, 2 是一个有效的序列
9. 考虑(x, y)的影响
10. 如果(x, y)连起来后，不影响结果，那么就是ok的
11. 考虑n = 8, 0, 1, 0, 1, 0, 1, 0, 1 
12. 如果连起来是的两个1，那么其中一个必须变成2（就ok？）
13. 如果连接的是两个0， 增加一个为1，好像会有问题
14. 0, 1, 0, 1 (增加其中一个到2，是ok的)
15. 考虑奇数个 
16. 0, 1, 0, 1, 2 (如果连接的是(0, 0) => 2)
17. 如果连接的是1, 1 (似乎有问题，但是可以改变其中一个为2) 变成连接1， 2
18. 可以的