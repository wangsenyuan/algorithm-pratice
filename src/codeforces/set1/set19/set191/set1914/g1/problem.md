There are 2𝑛
 light bulbs arranged in a row. Each light bulb has a color from 1
 to 𝑛
 (exactly two light bulbs for each color).

Initially, all light bulbs are turned off. You choose a set of light bulbs 𝑆
 that you initially turn on. After that, you can perform the following operations in any order any number of times:

choose two light bulbs 𝑖
 and 𝑗
 of the same color, exactly one of which is on, and turn on the second one;
choose three light bulbs 𝑖,𝑗,𝑘
, such that both light bulbs 𝑖
 and 𝑘
 are on and have the same color, and the light bulb 𝑗
 is between them (𝑖<𝑗<𝑘
), and turn on the light bulb 𝑗
.
You want to choose a set of light bulbs 𝑆
 that you initially turn on in such a way that by performing the described operations, you can ensure that all light bulbs are turned on.

Calculate two numbers:

the minimum size of the set 𝑆
 that you initially turn on;
the number of sets 𝑆
 of minimum size (taken modulo 998244353
).
Input
The first line of the input contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Then follow the descriptions of the test cases.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤1000
) — the number of pairs of light bulbs.

The second line of each test case contains 2𝑛
 integers 𝑐1,𝑐2,…,𝑐2𝑛
 (1≤𝑐𝑖≤𝑛
), where 𝑐𝑖
 is the color of the 𝑖
-th light bulb. For each color from 1
 to 𝑛
, exactly two light bulbs have this color.

Additional constraint on the input: the sum of values of 𝑛2
 over all test cases does not exceed 106
.

### ideas
1. 先考虑如何计算最小的set的size
2. 首先，两头的bulb，如果它们的颜色一致，那至少要一个
3. 如果它们的颜色不一样，那两个都必须点亮
4. 如果两头的颜色一致，貌似这个1，就是最小的size
5. 因为，其他的可以按照步骤2，一个个点亮
6. 但如果不一致，首先它们两个必须被点亮。但是接下来就麻烦了
7. 分两种种情况 
8. 1...1 ? ??? n... n
9. 1.。.n...1....n 
10. 第一种情况是中间部分是一个递归更小的问题，两头是颜色一致的情况
11. 第二种情况就比较复杂了, 选择1，可以搞定n，进而搞定第二个n，首先可以确定的是它的set的size是1
12. 但是方案数，却有很多，比如任何一个包含1、n的进而任何一个包含x(假设x的)
13. 所以变成一个树（图？）
14. 似乎是个图 (1...2...1...2) 它们是相互关联的
15. got了。就是变成图以后，计算有多少个component
16. 如果 a包含b，那么就有一条b -> a的有向边，然后看根节点有几个
17. 如果出现了环。选择环中的任意一个都可以