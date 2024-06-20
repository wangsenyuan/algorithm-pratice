Anton got bored during the hike and wanted to solve something. He asked Kirill if he had any new problems, and of course, Kirill had one.

You are given a permutation 𝑝
 of size 𝑛
, and a number 𝑥
 that needs to be found. A permutation of length 𝑛
 is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in arbitrary order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears twice in the array), and [1,3,4]
 is also not a permutation (𝑛=3
 but there is 4
 in the array).

You decided that you are a cool programmer, so you will use an advanced algorithm for the search — binary search. However, you forgot that for binary search, the array must be sorted.

You did not give up and decided to apply this algorithm anyway, and in order to get the correct answer, you can perform the following operation no more than 2
 times before running the algorithm: choose the indices 𝑖
, 𝑗
 (1≤𝑖,𝑗≤𝑛
) and swap the elements at positions 𝑖
 and 𝑗
.

After that, the binary search is performed. At the beginning of the algorithm, two variables 𝑙=1
 and 𝑟=𝑛+1
 are declared. Then the following loop is executed:

If 𝑟−𝑙=1
, end the loop
𝑚=⌊𝑟+𝑙2⌋
If 𝑝𝑚≤𝑥
, assign 𝑙=𝑚
, otherwise 𝑟=𝑚
.
The goal is to rearrange the numbers in the permutation before the algorithm so that after the algorithm is executed, 𝑝𝑙
 is equal to 𝑥
. It can be shown that 2
 operations are always sufficient.

#### ideas
1. simulate? 
2. 如果 x在前半段，那么x <= p[mid]
3. 但是如果此时 x > p[mid]怎么办？必须要交换。但交换什么呢？
4. 那直接把x交换到 mid?
5. 好像不大行。因为这时候 l = mid, r = n + 1
6. 先不考虑交换，假设pos[x]在mid的左边，那这时，我们需要 p[mid] > x
7. 如果 p[mid] 已经大于x了，good 继续。否则的话，需要交换一个比x大的数到mid位置来（记一次交换)
8. 如果 pos[x]在mid的右边(包括mid), 那么需要 p[mid] <= x
9. 如果p[mid] > x， 那么需要交换一个比x小的数到mid这里来
10. 如果x = n, 且它在最左边？会发生什么？
11. 6， 5， 4， 3， 2， 1 （x = 6)
12. (l, r = 1, 7) mid = 4, 
13. p[4] = 3 < 6, 所以要往右移动
14. 不管它，按照规则一直二分，如果p[l] == x => good, 如果不是good的，
15. 那这里分两种情况，一种是，x出现在了路径上
16. x没有出现在路径上，
17. 第二种情况，可以把x和l交换即可
18. 如果x出现在了路径上，就比较麻烦了，得先把它交换出去，再交换回来
19. 假设 p[i] = x, 且i出现在了路径上
20. 那么此时满足 p[i] <= x, 所以此时把不在路径上的，比x小的数交换进来？
21. 但如果不存在这样的数呢？也就是比x小的数，都在路径上面
22. 第一步，把p[n]和x进行交换
23. 那么然后按照步骤进行访问，如果l不是落在n上，再把p[n]和p[l]交换