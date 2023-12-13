You are given a string 𝑠
, consisting of lowercase Latin letters.

In one operation, you can select several (one or more) positions in it such that no two selected positions are adjacent
to each other. Then you remove the letters on the selected positions from the string. The resulting parts are
concatenated without changing their order.

What is the smallest number of operations required to make all the letters in 𝑠
the same?

### solution

1. 考虑最后剩下的字符是x，那么x一个都不用删除，删除其中的任何一个都没有任何好处，还可能造成原来不靠近的变成靠近了
2. 其他的字符需要删除掉，但是每次都不能删除连续的字符
3. 两个x中间的，假设长度是m,
4. 第一次，可以删除 (m + 1) / 2, 剩余 m - (m+1) / 2
5. 然后继续删除 其中的一半
6. 不同区间的可以同一次操作内完成，所以找出最大的区间即可

