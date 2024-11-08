Recently, on the course of algorithms and data structures, Valeriy learned how to use a deque. He built a deque filled with n
 elements. The i
-th element is ai
 (i
 = 1,2,…,n
). He gradually takes the first two leftmost elements from the deque (let's call them A
 and B
, respectively), and then does the following: if A>B
, he writes A
 to the beginning and writes B
 to the end of the deque, otherwise, he writes to the beginning B
, and A
 writes to the end of the deque. We call this sequence of actions an operation.

For example, if deque was [2,3,4,5,1]
, on the operation he will write B=3
 to the beginning and A=2
 to the end, so he will get [3,4,5,1,2]
.

The teacher of the course, seeing Valeriy, who was passionate about his work, approached him and gave him q
 queries. Each query consists of the singular number mj
 (j=1,2,…,q)
. It is required for each query to answer which two elements he will pull out on the mj
-th operation.

Note that the queries are independent and for each query the numbers A
 and B
 should be printed in the order in which they will be pulled out of the deque.

Deque is a data structure representing a list of elements where insertion of new elements or deletion of existing elements can be made from both sides.

### ideas
1. 如果m = 1, or n - 1, 都很好解决
2. 每次都往最远的地方跳
3. (0, 0) => (n - 1, 0)
4.  