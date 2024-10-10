Welcome! Everything is fine.

You have arrived in The Medium Place, the place between The Good Place and The Bad Place. You are assigned a task that will either make people happier or torture them for eternity.

You have a list of 𝑘
 pairs of people who have arrived in a new inhabited neighborhood. You need to assign each of the 2𝑘
 people into one of the 2𝑘
 houses. Each person will be the resident of exactly one house, and each house will have exactly one resident.

Of course, in the neighborhood, it is possible to visit friends. There are 2𝑘−1
 roads, each of which connects two houses. It takes some time to traverse a road. We will specify the amount of time it takes in the input. The neighborhood is designed in such a way that from anyone's house, there is exactly one sequence of distinct roads you can take to any other house. In other words, the graph with the houses as vertices and the roads as edges is a tree.

The truth is, these 𝑘
 pairs of people are actually soulmates. We index them from 1
 to 𝑘
. We denote by 𝑓(𝑖)
 the amount of time it takes for the 𝑖
-th pair of soulmates to go to each other's houses.

As we have said before, you will need to assign each of the 2𝑘
 people into one of the 2𝑘
 houses. You have two missions, one from the entities in The Good Place and one from the entities of The Bad Place. Here they are:

The first mission, from The Good Place, is to assign the people into the houses such that the sum of 𝑓(𝑖)
 over all pairs 𝑖
 is minimized. Let's define this minimized sum as 𝐺
. This makes sure that soulmates can easily and efficiently visit each other;
The second mission, from The Bad Place, is to assign the people into the houses such that the sum of 𝑓(𝑖)
 over all pairs 𝑖
 is maximized. Let's define this maximized sum as 𝐵
. This makes sure that soulmates will have a difficult time to visit each other.
What are the values of 𝐺
 and 𝐵
?

### ideas
1. B似乎很容易计算？就是对于每条边计算min(sz[u], n - sz[u]) * value[e] 加起来就可以了
2. G要怎么算呢？
3. 如果某个子树的size是奇数，肯定要有一对，要离开这条边
4. 否则，这条边其实可以不用到（当然也可以使用）
5. 也就是从2 * k - 1 条边中，选择k条边，将节点进行配对
6. dp[u][0] 表示将所有子树u进行配对后的最优解
7. dp[u][1] 表示，出了u，其他的子树节点都进行了配对的最优解
8. dp[u][0] = dp[v][0] + dp[c][1] + value(edge u, c)
9. 