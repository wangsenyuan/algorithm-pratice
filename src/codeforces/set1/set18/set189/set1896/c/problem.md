You are given two arrays 𝑎
and 𝑏
of size 𝑛
. The beauty of the arrays 𝑎
and 𝑏
is the number of indices 𝑖
such that 𝑎𝑖>𝑏𝑖
.

You are also given an integer 𝑥
. Determine whether it is possible to rearrange the elements of 𝑏
such that the beauty of the arrays becomes 𝑥
. If it is possible, output one valid rearrangement of 𝑏
.

### thoughts

1. a和b都可以先排序
2. 要找到x个a[i] > b[i], 那么应该优先考虑较大的数
3. 所以，选择a右端的x个数和b中左端的x个数，看是否满足条件
4. 