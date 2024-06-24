Elsie is learning how to paint. She has a canvas of 𝑛
 cells numbered from 1
 to 𝑛
 and can paint any (potentially empty) subset of cells.


Elsie has a 2D array 𝑎
 which she will use to evaluate paintings. Let the maximal contiguous intervals of painted cells in a painting be [𝑙1,𝑟1],[𝑙2,𝑟2],…,[𝑙𝑥,𝑟𝑥]
. The beauty of the painting is the sum of 𝑎𝑙𝑖,𝑟𝑖
 over all 1≤𝑖≤𝑥
. In the image above, the maximal contiguous intervals of painted cells are [2,4],[6,6],[8,9]
 and the beauty of the painting is 𝑎2,4+𝑎6,6+𝑎8,9
.

There are 2𝑛
 ways to paint the strip. Help Elsie find the 𝑘
 largest possible values of the beauty of a painting she can obtain, among all these ways. Note that these 𝑘
 values do not necessarily have to be distinct. It is guaranteed that there are at least 𝑘
 different ways to paint the canvas.

 ### ideas
 1. 显然最大的时候，所有的正值的和 + 所有的0
 2. 如果正值的个数 + 0 的个数超过 (2 << m) >= k, 可以只考虑这些最大的部分
 3. 似乎不管正值，可以找出最大的m个数， 2 << m >= k 的，然后brute force生成？
 4. 然后再排序即可
 5. 完蛋了， a[i][j] 表示从i到j连续序列的值，
 6. 所以，比如连续选了两个cell， 不能简单的a[i] + a[j]
 7. 按照题目要求，只要它们是连续的，就要计算 a[i...j] 不能是 a[i] + a[j]
 8. dp[i][j] 表示到i为止前j个元素的值
 9. 如果选取了一段连续的 l...r, 那么更新 dp[r][j] = (dp[l-2][j] + cost(l...r)) ++ cost(l....r)
 10. 没想法～
1.  

### solution

Let 𝚍𝚙[𝑖]
 store a list of all min(𝑘,2𝑖)
 highest beauties of a painting in non-increasing order if you only paint cells 1,2,…,𝑖
.

Let's define merging two lists as creating a new list that stores the 𝑘
 highest values from the two lists.

First, let's look at a slow method of transitioning. We iterate over the left endpoint 𝑙
 such that 𝑙…𝑖
 is a maximal painted interval. At each 𝑙
, we merge 𝚍𝚙[𝑙−2]
, with 𝑎𝑙,𝑖
 added to each value, into 𝚍𝚙[𝑖]
. We also merge 𝚍𝚙[𝑖−1]
 into 𝚍𝚙[𝑖]
 to handle the case in which we do not paint cell 𝑖
. If implemented well, transitioning takes (𝑛𝑘)
 time leading to an (𝑛2𝑘)
 solution which is too slow.

We can speed this up. This merging process boils down to finding the 𝑘
 largest elements from 𝑛
 non-increasing lists in ((𝑛+𝑘)log𝑛)
 time. We can use a priority queue that stores (𝚎𝚕𝚎𝚖𝚎𝚗𝚝,𝚒𝚗𝚍𝚎𝚡
) for each list. We can do the following 𝑘
 times: add the top of the priority queue to our answer, advance its index, and update the priority queue accordingly. This allows us to transition in ((𝑛+𝑘)log𝑛)
 time which leads to an ((𝑛2+𝑛𝑘)log𝑛)
 solution.

