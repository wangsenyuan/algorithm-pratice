You are given an array 𝑎
 consisting of 𝑛
 integers. The value of the 𝑖
-th element of the array is 𝑎𝑖
.

You are also given a set of 𝑚
 segments. The 𝑗
-th segment is [𝑙𝑗;𝑟𝑗]
, where 1≤𝑙𝑗≤𝑟𝑗≤𝑛
.

You can choose some subset of the given set of segments and decrease values on each of the chosen segments by one (independently). For example, if the initial array 𝑎=[0,0,0,0,0]
 and the given segments are [1;3]
 and [2;4]
 then you can choose both of them and the array will become 𝑏=[−1,−2,−2,−1,0]
.

You have to choose some subset of the given segments (each segment can be chosen at most once) in such a way that if you apply this subset of segments to the array 𝑎
 and obtain the array 𝑏
 then the value max𝑖=1𝑛𝑏𝑖−min𝑖=1𝑛𝑏𝑖
 will be maximum possible.

Note that you can choose the empty set.

If there are multiple answers, you can print any.

[problem link](https://codeforces.com/problemset/problem/1108/E1)


### ideas
1. n <= 300, m <= 300
2. 这个能二分吗？
3. 如果有个区间，它在所有的数字上进行作用，使用它/或者不使用它，对答案是没有影响的，可以忽略它（只有在其他的无法得到更优答案时，可以考虑使用它）
4. 现在考虑，所有的区间，都不能覆盖所有的元素
5. 那么考虑最优的答案x = max(b) - min(b)
6. 那么考虑最后使用的区间，这个区间肯定包含最小的那个数，但是不包含最大的那个数
7. 如果都包含/都不包含，那么它对结果没有影响
8. 如果反过来包含最大值，不包含最小值，显然因为只能减小的缘故，这样子只会得到更差的结果
9. 所以进一步推导出最大值：要么没有被覆盖到，要么它和最小值被最小值相同区域覆盖到？且那些相同的还是可以被忽略掉
10. 就分两类，一类是最大值没有被覆盖到，然后，找这个区域外，被覆盖到的最小值；要么找同一个覆盖区域中原有的最大值
11. 对于第二类，就是找那个没有任何操作时的x
12. 对于第二类，就是对于特定的位置i，它前面的所有的区域都使用，后面所有的区域也都是用，找最小值
13. 这里涉及到range update, range query的问题
14. 