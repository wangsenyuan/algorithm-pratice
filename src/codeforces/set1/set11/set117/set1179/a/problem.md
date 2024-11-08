Recently, on the course of algorithms and data structures, Valeriy learned how to use a deque. He built a deque filled with 𝑛
 elements. The 𝑖
-th element is 𝑎𝑖
 (𝑖
 = 1,2,…,𝑛
). He gradually takes the first two leftmost elements from the deque (let's call them 𝐴
 and 𝐵
, respectively), and then does the following: if 𝐴>𝐵
, he writes 𝐴
 to the beginning and writes 𝐵
 to the end of the deque, otherwise, he writes to the beginning 𝐵
, and 𝐴
 writes to the end of the deque. We call this sequence of actions an operation.

For example, if deque was [2,3,4,5,1]
, on the operation he will write 𝐵=3
 to the beginning and 𝐴=2
 to the end, so he will get [3,4,5,1,2]
.

The teacher of the course, seeing Valeriy, who was passionate about his work, approached him and gave him 𝑞
 queries. Each query consists of the singular number 𝑚𝑗
 (𝑗=1,2,…,𝑞)
. It is required for each query to answer which two elements he will pull out on the 𝑚𝑗
-th operation.

Note that the queries are independent and for each query the numbers 𝐴
 and 𝐵
 should be printed in the order in which they will be pulled out of the deque.

Deque is a data structure representing a list of elements where insertion of new elements or deletion of existing elements can be made from both sides.


### ideas
1. 每次操作，相当于把前两个数中的，最小值，移动到了尾巴上
2. 所以在a移动到尾巴前， 在它后面的部分，比它小的，都移动到了尾巴上，然后是a
3. 然后是比b小的数
4. 所以，感觉经过足够的次数后， a就变成sorted了。然后就一直在那里循环
5. 这个次数是多少呢？
6. 5,4,3,2,1
7. 5,3,2,1,4
8. 5,2,1,4,3
9. 5,1,4,3,2
10. 5,4,3,2,1
11. 所以上面那个结论不对。
12. 假设x是其中的最大值，那么一旦遇到x，那么剩下的，都不会再变了
13. 所以在遇到x前，可以模拟
14. 遇到x后，就可以计算了
15. 