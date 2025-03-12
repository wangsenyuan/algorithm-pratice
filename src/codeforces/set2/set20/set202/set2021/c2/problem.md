0≤𝑞≤2⋅105
. You can make hacks only if all the versions of the problem are solved.

A team consisting of 𝑛
 members, numbered from 1
 to 𝑛
, is set to present a slide show at a large meeting. The slide show contains 𝑚
 slides.

There is an array 𝑎
 of length 𝑛
. Initially, the members are standing in a line in the order of 𝑎1,𝑎2,…,𝑎𝑛
 from front to back. The slide show will be presented in order from slide 1
 to slide 𝑚
. Each section will be presented by the member at the front of the line. After each slide is presented, you can move the member at the front of the line to any position in the lineup (without changing the order of the rest of the members). For example, suppose the line of members is [3,1,2,4]
. After member 3
 presents the current slide, you can change the line of members into either [3,1,2,4]
, [1,3,2,4]
, [1,2,3,4]
 or [1,2,4,3]
.

There is also an array 𝑏
 of length 𝑚
. The slide show is considered good if it is possible to make member 𝑏𝑖
 present slide 𝑖
 for all 𝑖
 from 1
 to 𝑚
 under these constraints.

However, your annoying boss wants to make 𝑞
 updates to the array 𝑏
. In the 𝑖
-th update, he will choose a slide 𝑠𝑖
 and a member 𝑡𝑖
 and set 𝑏𝑠𝑖:=𝑡𝑖
. Note that these updates are persistent, that is changes made to the array 𝑏
 will apply when processing future updates.

For each of the 𝑞+1
 states of array 𝑏
, the initial state and after each of the 𝑞
 updates, determine if the slideshow is good.



### ideas
1. 脑子有点打结
2. 先考虑没有modification的情况
3. 就是 b[1]表示第一个slide要展示的对象，b[2]表示第二个要展示的对象，依次类推
4. 假设到i为止，都可以满足；现在考虑i，它要展示给b[i]
5. 如果b[i]是一个新的用户。那么它必须是目前头部的用户（已经出现过的，进入一个未决的状态）
6. 如果b[i]是一个老用户，那么就没有问题，不管i-1前面是如何处理，总可以找到一个方案，让b[i]出现在这里
7. 如果不是上面两种状态，那么b就不是一个good的序列
8. 假设当前的的序列b可以到达的最远的是第i个slice，到了i的时候， b[i] 不是第一个未出现的member
9. 现在考虑各种情况， 如果修改的 b[si] = t, si > i, 那么只需要记录下来就可以了（还没有处理到）
10. 如果 si < i, 修改了前面的某个观看对象是新的人
11.   又分两种情况，一种是新对象是在si前出现过的（没有问题）
12.   t是在si处第一次出现，这个就有点麻烦了。一个就是它是新的i，还有一种可能性是，i不需要变化
13. 这里不变的东西，貌似是序列a
14. 那如何从序列a的视角看待呢？
15. 就是进入活跃（未决）状态的member越多越好。且i进入活跃状态后，i-1肯定也在活跃
16. 假设目前的状态是 (i, j) 表示前i个member进入了活跃状态，且看完了前j个slides
17. 现在考虑让某个新的member x (x > i) 要看到slide y, 且y <= j
18. 假设f(j)表示slide j能够激活的最多的人数，如果x出现在前f(j)中，那么这个结果是没有变的，(x仍然是激活的)
19. 如果x在i的后面，那么肯定是无法到达的。也就是退回到状态(f(y), y)了