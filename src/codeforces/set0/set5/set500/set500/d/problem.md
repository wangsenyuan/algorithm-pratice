New Year is coming in Tree World! In this world, as the name implies, there are n cities connected by n - 1 roads, and for any two distinct cities there always exists a path between them. The cities are numbered by integers from 1 to n, and the roads are numbered by integers from 1 to n - 1. Let's define d(u, v) as total length of roads on the path between city u and city v.

As an annual event, people in Tree World repairs exactly one road per year. As a result, the length of one road decreases. It is already known that in the i-th year, the length of the ri-th road is going to become wi, which is shorter than its length before. Assume that the current year is year 1.

Three Santas are planning to give presents annually to all the children in Tree World. In order to do that, they need some preparation, so they are going to choose three distinct cities c1, c2, c3 and make exactly one warehouse in each city. The k-th (1 ≤ k ≤ 3) Santa will take charge of the warehouse in city ck.

It is really boring for the three Santas to keep a warehouse alone. So, they decided to build an only-for-Santa network! The cost needed to build this network equals to d(c1, c2) + d(c2, c3) + d(c3, c1) dollars. Santas are too busy to find the best place, so they decided to choose c1, c2, c3 randomly uniformly over all triples of distinct numbers from 1 to n. Santas would like to know the expected value of the cost needed to build the network.

However, as mentioned, each year, the length of exactly one road decreases. So, the Santas want to calculate the expected after each length change. Help them to calculate the value.

### ideas
1. 先不考虑变化，那么这个expect值，应该等于每条边的期望的贡献值的和？
2. 对于边i，有多种情况，一种是3个点都在它的一边，一个点在它的右边，两个点在左边(or reversed)
3. 假设它一边是 x个点，另外一边是 n - x个点，那么不经过它的情况 = c(x, 3) + c(n-x, 3)
4. 经过它的 = C(x, 2) * C(n-x, 1) + C(x, 1) * C(n-x, 2)
5. 加起来后再除以C(n,3)即可
6. 那么一条边变短的情况下，就是看它变小的贡献即可？