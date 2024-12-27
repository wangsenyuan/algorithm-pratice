Ksyusha decided to start a game development company. To stand out among competitors and achieve success, she decided to write her own game engine. The engine must support a set initially consisting of 𝑛
 distinct integers 𝑎1,𝑎2,…,𝑎𝑛
.

The set will undergo 𝑚
 operations sequentially. The operations can be of the following types:

Insert element 𝑥
 into the set;
Remove element 𝑥
 from the set;
Report the 𝑘
-load of the set.
The 𝑘
-load of the set is defined as the minimum positive integer 𝑑
 such that the integers 𝑑,𝑑+1,…,𝑑+(𝑘−1)
 do not appear in this set. For example, the 3
-load of the set {3,4,6,11}
 is 7
, since the integers 7,8,9
 are absent from the set, and no smaller value fits.

Ksyusha is busy with management tasks, so you will have to write the engine. Implement efficient support for the described operations.

### ideas
1. 维护一个tree，区间l...r, 已经这个区间内最大的连续空洞的大小k
2. 对于一个查询k, 如果它的左子树的k>= k ， go left
3. 但是问题出现在左子树 < k, 
4. 这里有两种情况，一种左边确实没有这样的区间，go right
5. 还有一种情况，就是左边是有的，但是它是跨mid的，所以，还必须知道左区间的最大值，和右区间的最小值
6. 查询没有问题，但是插入/删除呢？应该也可以维护