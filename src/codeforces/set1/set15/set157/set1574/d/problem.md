Ivan is playing yet another roguelike computer game. He controls a single hero in the game. The hero has 𝑛
 equipment slots. There is a list of 𝑐𝑖
 items for the 𝑖
-th slot, the 𝑗
-th of them increases the hero strength by 𝑎𝑖,𝑗
. The items for each slot are pairwise distinct and are listed in the increasing order of their strength increase. So, 𝑎𝑖,1<𝑎𝑖,2<⋯<𝑎𝑖,𝑐𝑖
.

For each slot Ivan chooses exactly one item. Let the chosen item for the 𝑖
-th slot be the 𝑏𝑖
-th item in the corresponding list. The sequence of choices [𝑏1,𝑏2,…,𝑏𝑛]
 is called a build.

The strength of a build is the sum of the strength increases of the items in it. Some builds are banned from the game. There is a list of 𝑚
 pairwise distinct banned builds. It's guaranteed that there's at least one build that's not banned.

What is the build with the maximum strength that is not banned from the game? If there are multiple builds with maximum strength, print any of them.

Input
The first line contains a single integer 𝑛
 (1≤𝑛≤10
) — the number of equipment slots.

The 𝑖
-th of the next 𝑛
 lines contains the description of the items for the 𝑖
-th slot. First, one integer 𝑐𝑖
 (1≤𝑐𝑖≤2⋅105
) — the number of items for the 𝑖
-th slot. Then 𝑐𝑖
 integers 𝑎𝑖,1,𝑎𝑖,2,…,𝑎𝑖,𝑐𝑖
 (1≤𝑎𝑖,1<𝑎𝑖,2<⋯<𝑎𝑖,𝑐𝑖≤108
).

The sum of 𝑐𝑖
 doesn't exceed 2⋅105
.

The next line contains a single integer 𝑚
 (0≤𝑚≤105
) — the number of banned builds.

Each of the next 𝑚
 lines contains a description of a banned build — a sequence of 𝑛
 integers 𝑏1,𝑏2,…,𝑏𝑛
 (1≤𝑏𝑖≤𝑐𝑖
).

The builds are pairwise distinct, and there's at least one build that's not banned.

Output
Print the build with the maximum strength that is not banned from the game. If there are multiple builds with maximum strength, print any of them.


### ideas
1. 问题很直观，但却无从下手，显然无法所有的组合都尝试一遍，因为这个组合是天文数字
2. 而且大部分组合也是多余的。
3. 比如对于第一组来说，如果其他组的组合的数量超过了m，那么最后一个，肯定可以在某个组合中出现（但不一定是最优的组合中）
4. 按照ban的sum，从高到低排序，对于i, 如果b[j][i] < c[i], 且不存在ban[x], 在替换了 c[i]后，存在
5. 那么就可以用c[i]替换，产生一个有效的组合
6. 不大对，如果存在呢？那就要尝试下一个。。。。 无穷无尽了
7. 还有一个就是，从最右边开始尝试，名次使用最接近的字符去替换，然后检查，直到找到有效的
8. 这样子的复杂性 = m * x
9. 因为最多检查m次，x依赖于替换的效率？