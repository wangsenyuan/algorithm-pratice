There are 𝑛
trees in a park, numbered from 1
to 𝑛
. The initial height of the 𝑖
-th tree is ℎ𝑖
.

You want to water these trees, so they all grow to the same height.

The watering process goes as follows. You start watering trees at day 1
. During the 𝑗
-th day you can:

Choose a tree and water it. If the day is odd (e.g. 1,3,5,7,…
), then the height of the tree increases by 1
. If the day is even (e.g. 2,4,6,8,…
), then the height of the tree increases by 2
.
Or skip a day without watering any tree.
Note that you can't water more than one tree in a day.

Your task is to determine the minimum number of days required to water the trees so they grow to the same height.

You have to answer 𝑡
independent test cases.

### ideas

1. 显然如果能在x天完成，必然也可以在x+1, x+2天完成，以后后面可以简单的啥也不做
2. 如何判断能在x天完成呢？
3. 首先最高的tree，要们不长高，要么长高一米，不用长高两米，所以可以在这个前提下进行处理