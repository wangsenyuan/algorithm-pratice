Jury initially had a sequence a
 of length n
, such that ai=i
.

The jury chose three integers i
, j
, k
, such that 1≤i<j<k≤n
, j−i>1
. After that, Jury reversed subsegments [i,j−1]
 and [j,k]
 of the sequence a
.

Reversing a subsegment [l,r]
 of the sequence a
 means reversing the order of elements al,al+1,…,ar
 in the sequence, i. e. al
 is swapped with ar
, al+1
 is swapped with ar−1
, etc.

You are given the number n
 and you should find i
, j
, k
 after asking some questions.

In one question you can choose two integers l
 and r
 (1≤l≤r≤n
) and ask the number of inversions on the subsegment [l,r]
 of the sequence a
. You will be given the number of pairs (i,j)
 such that l≤i<j≤r
, and ai>aj
.

Find the chosen numbers i
, j
, k
 after at most 40
 questions.

The numbers i
, j
, and k
 are fixed before the start of your program and do not depend on your queries.

### ideas
1. ask (1, n) => l * (l - 1) / 2 + r * (r - 1) / 2
2. 一个长度为n的，降序的permuation的inversion数量 = (n - 1) * n / 2
3. 上吗的那个l, r 能够唯一确定吗？不是很清楚。
4. 然后ask (mid, n) 如果 mid < i, 那么这个式子就是  = ask(1, n)的
5. 如果 (1, mid) == ask(1, n) => mid >= k
6. 这样子，可以找到i和k，
7. l + r = (k - i + 1)
8. 这样子似乎就能解出来了。这样子的j有两个位置，可以直接多查询一次。
9. 但是n是1e9，找到i(或者k)需要各30次，貌似会超出40次的限制
10. 先找到i， 然后k到i的距离假设为w
11. 当l = w - 2， r = 2 时， （w - 2) * (w - 3) / 2 + 1 >= value
12. 当l = r时，在给定的value时，w的值最大
13. 当 l = 2(w - 2)时，在给定的value时，w的值最小
14. 所以，这里就能得到k的上届和下届（猜想这个范围不会超过1000）
15. 然后在这个范围里面二分查找k