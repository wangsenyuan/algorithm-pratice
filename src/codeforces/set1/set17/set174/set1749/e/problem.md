Monocarp is playing Minecraft and wants to build a wall of cacti. He wants to build it on a field of sand of the size of
𝑛×𝑚
cells. Initially, there are cacti in some cells of the field. Note that, in Minecraft, cacti cannot grow on cells
adjacent to each other by side — and the initial field meets this restriction. Monocarp can plant new cacti (they must
also fulfil the aforementioned condition). He can't chop down any of the cacti that are already growing on the field —
he doesn't have an axe, and the cacti are too prickly for his hands.

Monocarp believes that the wall is complete if there is no path from the top row of the field to the bottom row, such
that:

each two consecutive cells in the path are adjacent by side;
no cell belonging to the path contains a cactus.
Your task is to plant the minimum number of cacti to build a wall (or to report that this is impossible).

#### thoughts

1. 首先这肯定是一个关于图的问题
2. 怎么阻止从top到bottom的路径呢？
3. 从横向看，每一列都需要至少一个格子；且这些格子，通过corner连接后，要隔断上下
4. 如果把左右两边看作两个点，中间是路径，就是找最短距离
5. 如何避免和已有的仙人掌相临？
6. 