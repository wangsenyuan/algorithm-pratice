Let's suppose you have an array a, a stack s (initially empty) and an array b (also initially empty).

You may perform the following operations until both a and s are empty:

Take the first element of a, push it into s and remove it from a (if a is not empty);
Take the top element from s, append it to the end of array b and remove it from s (if s is not empty).
You can perform these operations in arbitrary order.

If there exists a way to perform the operations such that array b is sorted in non-descending order in the end, then
array a is called stack-sortable.

For example, [3, 1, 2] is stack-sortable, because b will be sorted if we perform the following operations:

Remove 3 from a and push it into s;
Remove 1 from a and push it into s;
Remove 1 from s and append it to the end of b;
Remove 2 from a and push it into s;
Remove 2 from s and append it to the end of b;
Remove 3 from s and append it to the end of b.
After all these operations b =[1, 2, 3], so [3, 1, 2] is stack-sortable. [2, 3, 1] is not stack-sortable.

You are given k first elements of some permutation p of size n (recall that a permutation of size n is an array of size
n where each integer from 1 to n occurs exactly once). You have to restore the remaining n - k elements of this
permutation so it is stack-sortable. If there are multiple answers, choose the answer such that p is lexicographically
maximal (an array q is lexicographically greater than an array p iff there exists some integer k such that for every i<k
qi = pi, and qk>pk). You may not swap or change any of first k elements of the permutation.

Print the lexicographically maximal permutation p you can obtain.

If there exists no answer then output -1.

### ideas

1. 考虑a在什么情况下，才是sortable的
2. 如果 存在这样一个序列 i < j < k
3. 1, 2, 3 => ok
4. 1, 3, 2 => ok
5. 3, 1, 2 => ok
6. 3, 2, 1 => ok
7. 2, 3, 1 => bad
8. a[i] < a[j] > a[k], and a[i] > a[k]
9. 因为当a[i]进入stack后，因为还有比它小的数，它不能出栈
10. 但是a[j]进入的时候，也不能出栈，
11. 但是a[j]的前面有比它小的在栈中，且它在栈中，导致更小的a[i]无法出栈
12. 对于前k个，如果存在 a[i] < a[j] 且 最小的数x（包括j后面，已经未出现的）a[i] > x
13. 那么 =》 -1
14. 然后开始考虑如何构造后面的部分
15. 是不是应该从后面开始，尽量的使用最小的值，看能不能放到后面，以使得前面的值尽量的大？
16. 似乎还是有点困难
17. 对于x，如果现在未在后面填入的，包括前面已填入的，比它大的最小的值是y, 如果y和它之间的空位 > 未填入的比y更小的数
18. 那么这个地方就不能放入x，因为如果放入了x，必然会在这些空位中间填入某个比y大的数
19. x越小（比如1）那么y越小，但是中间的空位（是固定的），比y小的未填入的数就更少，那么就有可能不满足条件
20. 所以x需要一个足够大的数
21. 假设填好了x，那么下一个数理论上可以比x要小，这是因为首先这里消耗了一个比y小的数（x), 且空位也变得更少了
22. 当空位是w+1时，那么这个x就是第w+1个未选中的值，y = x + 1