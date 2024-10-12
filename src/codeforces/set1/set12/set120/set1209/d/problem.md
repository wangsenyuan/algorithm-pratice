The legendary Farmer John is throwing a huge party, and animals from all over the world are hanging out at his house. His guests are hungry, so he instructs his cow Bessie to bring out the snacks! Moo!

There are 𝑛
 snacks flavors, numbered with integers 1,2,…,𝑛
. Bessie has 𝑛
 snacks, one snack of each flavor. Every guest has exactly two favorite flavors. The procedure for eating snacks will go as follows:

First, Bessie will line up the guests in some way.
Then in this order, guests will approach the snacks one by one.
Each guest in their turn will eat all remaining snacks of their favorite flavor. In case no favorite flavors are present when a guest goes up, they become very sad.
Help Bessie to minimize the number of sad guests by lining the guests in an optimal way.

### ideas
1. 把snack作为节点, 每个人作为边，组成一个图
2. 每个节点被访问一次，每条边也最多被访问一次
3. 希望被访问到的边，最多
4. 是不是按照节点的deg来处理就可以了？
5. deg越小的，越早处理掉越好；
6. 题目理解错了，第一个人到达的时候，会把两个都吃点