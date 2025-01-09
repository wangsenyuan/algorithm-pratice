The nation of Panel holds an annual show called The Number Games, where each district in the nation will be represented by one contestant.

The nation has 𝑛
 districts numbered from 1
 to 𝑛
, each district has exactly one path connecting it to every other district. The number of fans of a contestant from district 𝑖
 is equal to 2𝑖
.

This year, the president decided to reduce the costs. He wants to remove 𝑘
 contestants from the games. However, the districts of the removed contestants will be furious and will not allow anyone to cross through their districts.

The president wants to ensure that all remaining contestants are from districts that can be reached from one another. He also wishes to maximize the total number of fans of the participating contestants.

Which contestants should the president remove?

### ideas
1. n肯定要能留下来，单独留下n，都是最好的结果
2. 所以以n为root
3. 如果节点u,没有保留，那么它的整个子树也不需要保留
4. 如果能保留n-1, 也应该尽量的保留
5. 如果 n...n-1中间的节点的数量（包括两端， <= n - k) 保留它
6. 然后是 n - 2, ...
7. 但问题是要找到u到目前active集合的最短距离才行，就是目前它的祖先节点的最低位置
8. 使用light-heavy decomposition是可以的，因为祖先节点肯定在一条路径上
9. 有没有其他的结构可以用呢？