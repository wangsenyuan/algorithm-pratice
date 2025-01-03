You are given two integer arrays 𝑎
 and 𝑏
, both of size 𝑛
.

Let's define the cost of the subarray [𝑙,𝑟]
 as 𝑎𝑙+𝑎𝑙+1+⋯+𝑎𝑟−1+𝑎𝑟+𝑏𝑙+𝑏𝑟
. If 𝑙=𝑟
, then the cost of the subarray is 𝑎𝑙+2⋅𝑏𝑙
.

You have to perform queries of three types:

"1
 𝑝
 𝑥
" — assign 𝑎𝑝:=𝑥
;
"2
 𝑝
 𝑥
" — assign 𝑏𝑝:=𝑥
;
"3
 𝑙
 𝑟
" — find two non-empty non-overlapping subarrays within the segment [𝑙,𝑟]
 with the maximum total cost and print their total cost.


 ### ideas
 1. 假设选中了4个位置a < b < c < d
 2. p[d+1] + b[d] - (p[c] - b[c]) + p[b+1] + b[b] - (p[a] - b[a])
 3. p[d+1] + b[d] - (p[a] - b[a]) - (p[c] - b[c] - p[b+1] - b[b])
 4. p[d+1] + b[d] + p[b+1] + b[b] - (p[c] - b[c] + p[a] - b[a])
 5. 又一个限制是b要在a.c中间，d要在最后面
 6. 如果没有这个限制就很简单了
 7. 也可以看作是 p[d+1] + b[d] - (p[a] - b[a]) 这是以a...d的一个区间，然后在这个区间内
 8. 删除掉 p[c] - b[c] - p[b+1] - b[b] (b+1...c-1)这个区间（端点有点麻烦）
 9. 如果有个线段树，可以区间查询l...r之间的最大值，就好了
 10. 那么在这个线段树上维护什么信息呢？要么答案在左边，要么答案在右边
 11. 要么答案在两边。如果在两边，还有两种情况，这两个区间正好在两边，那么就是两边的最大值的和
 12. 要么其中一个区间在左边，另外一个区间跨了中点
 13. 貌似可以搞