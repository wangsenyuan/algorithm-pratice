In the nation of AtCoder, there are 
N cities numbered 
1 to 
N, and 
M roads numbered 
1 to 
M.
Road 
i connects cities 
A 
i
​
  and 
B 
i
​
  bidirectionally and has a length of 
C 
i
​
 .

You are given 
Q queries to process in order. The queries are of the following two types.

1 i: Road 
i becomes closed.
2 x y: Print the shortest distance from city 
x to city 
y, using only roads that are not closed. If city 
y cannot be reached from city 
x, print -1 instead.
It is guaranteed that each test case contains at most 
300 queries of the first type.

### ideas
1. 假设操作1，都进行brute force 的操作 
2. 300 * n * m (显然不行)
3. 如果能降低复杂性到300 * n 是ok的
4. 这里的关键是，删除一条边以后，影响的应该是少数才对
5. 倒过来能处理吗？
6. 如果在某个时刻删除了边i, 那么在这个时刻后，边i是不可用的
7. 那么倒过来处理，相当于在这个时刻将边i加入了进来
8. 好像是可以的
9. 然后就是加入新边后怎么处理（这个可以用flood算法