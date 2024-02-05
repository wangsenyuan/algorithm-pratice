You have an array 𝑎
of size 𝑛
, initially filled with zeros (𝑎1=𝑎2=…=𝑎𝑛=0
). You also have an array of integers 𝑐
of size 𝑛
.

Initially, you have 𝑘
coins. By paying 𝑐𝑖
coins, you can add 1
to all elements of the array 𝑎
from the first to the 𝑖
-th element (𝑎𝑗+=1
for all 1≤𝑗≤𝑖
). You can buy any 𝑐𝑖
any number of times. A purchase is only possible if 𝑘≥𝑐𝑖
, meaning that at any moment 𝑘≥0
must hold true.

Find the lexicographically largest array 𝑎
that can be obtained.

An array 𝑎
is lexicographically smaller than an array 𝑏
of the same length if and only if in the first position where 𝑎
and 𝑏
differ, the element in array 𝑎
is smaller than the corresponding element in 𝑏
.

### thoughts

1. 应该尽可能的增加a[0], 再尽可能的增加a[1], ....
2. 如果两个c[i] = c[j], i < j 那么应该选择j,因为它们的成本一样，但j的收益更大
3. 按照c升序排，相同时，按照下标倒序排
4. 因为c越小，增加的次数越多，a[0]升高就越多
5. ～～～ c可以被使用多次
6. 那一个可能性就是尽量的使用最小的，然后在不改变前面的结果时，替换成一个更长的
7. 好难呐～
8. 先不考虑细节，整体上考虑一下，
9. 假设最后使用了一些下标 i1, i2, i3, .., im 得到了最大的一个a
10. 其中每项使用了x1, x2, x3, ..., xm 次
11. 需要满足什么条件呢？
12. c[i1] * x1 + c[i2] * x2 + ... + c[im] * xm <= k
13. (k - c[i1] * x1) / c[i1] == x2
14. => k - c[i1] * x1 < c[i1] * x2
15. 比如 c = [3, 4], k = 8, 那么选择两个4更合理
16. got， 按照 k / c[i] 进行排序， 然后按照i,进行降序排
17. 前面不对。
18. 居然搞对了～