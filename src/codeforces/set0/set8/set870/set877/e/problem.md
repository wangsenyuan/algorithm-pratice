Danil decided to earn some money, so he had found a part-time job. The interview have went well, so now he is a light
switcher.

Danil works in a rooted tree (undirected connected acyclic graph) with n vertices, vertex 1 is the root of the tree.
There is a room in each vertex, light can be switched on or off in each room. Danil's duties include switching light in
all rooms of the subtree of the vertex. It means that if light is switched on in some room of the subtree, he should
switch it off. Otherwise, he should switch it on.

Unfortunately (or fortunately), Danil is very lazy. He knows that his boss is not going to personally check the work.
Instead, he will send Danil tasks using Workforces personal messages.

There are two types of tasks:

pow v describes a task to switch lights in the subtree of vertex v.
get v describes a task to count the number of rooms in the subtree of v, in which the light is turned on. Danil should
send the answer to his boss using Workforces messages.
A subtree of vertex v is a set of vertices for which the shortest path from them to the root passes through v. In
particular, the vertex v is in the subtree of v.

Danil is not going to perform his duties. He asks you to write a program, which answers the boss instead of him.

Input
The first line contains a single integer n (1 ≤ n ≤ 200 000) — the number of vertices in the tree.

The second line contains n - 1 space-separated integers p2, p3, ..., pn (1 ≤ pi<i), where pi is the ancestor of vertex
i.

The third line contains n space-separated integers t1, t2, ..., tn (0 ≤ ti ≤ 1), where ti is 1, if the light is turned
on in vertex i and 0 otherwise.

The fourth line contains a single integer q (1 ≤ q ≤ 200 000) — the number of tasks.

The next q lines are get v or pow v (1 ≤ v ≤ n) — the tasks described above.

Output
For each task get v print the number of rooms in the subtree of v, in which the light is turned on.

### ideas

1. 使用eular tour + range lazy update/query 应该是可以的
2. 每个节点表示一个子树的范围，存醋两个信息on/off, 在switch的时候交换这两个信息
3. 用lazy update处理
4. 有没有其他的方案?
5. 不需要这么复杂， get u的时候， 只要记录有多少个parent被处理过，以及总次数就可以了
6. 所以用那个 heavy-light分解的方案更好
7. 不对， 在子树内部的操作也会影响结果，还是要用前面的lazy update