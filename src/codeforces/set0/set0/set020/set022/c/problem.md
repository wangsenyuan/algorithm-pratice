Bob got a job as a system administrator in X corporation. His first task was to connect n servers with the help of m two-way direct connection so that it becomes possible to transmit data from one server to any other server via these connections. Each direct connection has to link two different servers, each pair of servers should have at most one direct connection. Y corporation, a business rival of X corporation, made Bob an offer that he couldn't refuse: Bob was asked to connect the servers in such a way, that when server with index v fails, the transmission of data between some other two servers becomes impossible, i.e. the system stops being connected. Help Bob connect the servers.


### ideas
1. 如果v失效了，存在x，y无法联通的情况
2. 那么只需要把(x, y)放在两个component中，每个component是个完全图
3. 在完全图中，需要的线的个数 S = cnt * (cnt - 1) / 2
4. 所以n越大越好，就可以尽量的消耗线
5. 这里cnt = n - 1 也就是说如果 S >= M - 1 (有一条线用来连接v和额外的节点)
6. 那么就存在答案，否则不存在