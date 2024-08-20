Given an integer 
N
N, consider all arrays 
A
A of size 
N
N such that:

All the elements are non-negative and distinct.
All prefix sums are odd. Formally, for all 
i
i such that 
1
≤
i
≤
N
1≤i≤N, 
∑
j
=
1
i
A
i
∑ 
j=1
i
​
 A 
i
​
  is odd.
Among all possible arrays 
A
A, output the smallest possible sum of the elements of the array.

Note: Since the Input/Output may be large, it is preferred to use fast I/O.

### ideas
1. 所有的前缀和必须是奇数， 说明第一个数是奇数，而其他的都是偶数
2. 1,0,2,4...
3. 1 + (0 + 2 * (n - 2)) * (n - 1) / 2
4. 1 + (n - 2) * (n - 1)