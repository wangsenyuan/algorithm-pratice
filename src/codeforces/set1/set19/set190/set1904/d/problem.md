You are given two arrays 𝑎
and 𝑏
of length 𝑛
.

You can perform the following operation some (possibly zero) times:

choose 𝑙
and 𝑟
such that 1≤𝑙≤𝑟≤𝑛
.
let 𝑥=max(𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟)
.
for all 𝑙≤𝑖≤𝑟
, set 𝑎𝑖:=𝑥
.
Determine if you can make array 𝑎
equal to array 𝑏
.

### thoughts

1. 如果a[i] > b[i] => false
2. 对下标i进行，按照b[i]从小到大排序
3. 假设要将a[i] => b[i] 要么从左边设置，要么从右边设置，如果这两边距离内已经设置了某个更小的数了，那么就不可以