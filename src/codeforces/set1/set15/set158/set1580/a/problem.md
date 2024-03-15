CQXYM found a rectangle 𝐴
of size 𝑛×𝑚
. There are 𝑛
rows and 𝑚
columns of blocks. Each block of the rectangle is an obsidian block or empty. CQXYM can change an obsidian block to an
empty block or an empty block to an obsidian block in one operation.

A rectangle 𝑀
size of 𝑎×𝑏
is called a portal if and only if it satisfies the following conditions:

𝑎≥5,𝑏≥4
.
For all 1<𝑥<𝑎
, blocks 𝑀𝑥,1
and 𝑀𝑥,𝑏
are obsidian blocks.
For all 1<𝑥<𝑏
, blocks 𝑀1,𝑥
and 𝑀𝑎,𝑥
are obsidian blocks.
For all 1<𝑥<𝑎,1<𝑦<𝑏
, block 𝑀𝑥,𝑦
is an empty block.
𝑀1,1,𝑀1,𝑏,𝑀𝑎,1,𝑀𝑎,𝑏
can be any type.
Note that the there must be 𝑎
rows and 𝑏
columns, not 𝑏
rows and 𝑎
columns.
Note that corners can be any type

CQXYM wants to know the minimum number of operations he needs to make at least one sub-rectangle a portal.

### ideas

1. 如果迭代a, b, i, j 复杂性是 4 * 5 * 5 * 5 * 1e8
2. 显然是不可以的，至少要减少一个维度（或者2个）
3. 能否2分呢？
4. 可能会存在这样一种情况，就是一个很大的空白区域，然后四边也满足条件（或者做很少的改动）
5. 这个时候的a/b是非常大的
6. 所以还是没有抓住问题的本质～
7. 给定上界和下界(i, j), 然后扫过列c, 在确定c的时候，就确定了这一列需要的操作数，假设是x
8. 那么剩余的操作数就是best - x，那此时有办法了，就是要知道上、下边界变成1，中间变成0的操作数