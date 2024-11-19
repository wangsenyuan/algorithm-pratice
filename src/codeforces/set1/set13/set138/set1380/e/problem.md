You have a set of 𝑛
 discs, the 𝑖
-th disc has radius 𝑖
. Initially, these discs are split among 𝑚
 towers: each tower contains at least one disc, and the discs in each tower are sorted in descending order of their radii from bottom to top.

You would like to assemble one tower containing all of those discs. To do so, you may choose two different towers 𝑖
 and 𝑗
 (each containing at least one disc), take several (possibly all) top discs from the tower 𝑖
 and put them on top of the tower 𝑗
 in the same order, as long as the top disc of tower 𝑗
 is bigger than each of the discs you move. You may perform this operation any number of times.

For example, if you have two towers containing discs [6,4,2,1]
 and [8,7,5,3]
 (in order from bottom to top), there are only two possible operations:

move disc 1
 from the first tower to the second tower, so the towers are [6,4,2]
 and [8,7,5,3,1]
;
move discs [2,1]
 from the first tower to the second tower, so the towers are [6,4]
 and [8,7,5,3,2,1]
.
Let the difficulty of some set of towers be the minimum number of operations required to assemble one tower containing all of the discs. For example, the difficulty of the set of towers [[3,1],[2]]
 is 2
: you may move the disc 1
 to the second tower, and then move both discs from the second tower to the first tower.

You are given 𝑚−1
 queries. Each query is denoted by two numbers 𝑎𝑖
 and 𝑏𝑖
, and means "merge the towers 𝑎𝑖
 and 𝑏𝑖
" (that is, take all discs from these two towers and assemble a new tower containing all of them in descending order of their radii from top to bottom). The resulting tower gets index 𝑎𝑖
.

For each 𝑘∈[0,𝑚−1]
, calculate the difficulty of the set of towers after the first 𝑘
 queries are performed.

 ### ideas
 1. 考虑状态变化，假设n，n-1...i，如果i-1和i不在同一个tower上，那么至少需要一次把i-1,从某个地方带到i这里来
 2. 那么也就是这样的i的个数？
 3. 怎么维护这个状态呢？
 4. 每个序列都维护几个节点，就是那些i的上面不是i-1的数字
 5. 这样merge的时候处理后，得到新的tower
 6. n - 1 - 连续的序列的值
 7. 这样似乎就简单了
 8. 计算有多少个位置它们是连续的
 9. 合并的时候，只需要处理边界
 10. 比如合并a, b, 其中有一段，[l, r] 
 11. 如果l-1在a中（假设b是少的那一段）如果，那么更新l-1为1即可
 12. 如果r+1在a中，更新r为1 （或者反过来）
 13. 剩下的问题是，如何检查l-1或者r+1在a中，以及如何维护这个区间[l, r]
 14. 这个区间在哪个集合中，可以单独维护
 15. 主要的问题在于，知道边界，以合并区间
 16. 倒是可以用AVL tree来处理，但需要合并的时候，直接把对应区间删除就可以了
 17. 