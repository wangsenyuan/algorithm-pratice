Pay attention to the non-standard memory limit in this problem.

In order to cut off efficient solutions from inefficient ones in this problem, the time limit is rather strict. Prefer to use compiled statically typed languages (e.g. C++). If you use Python, then submit solutions on PyPy. Try to write an efficient solution.

The array 𝑎=[𝑎1,𝑎2,…,𝑎𝑛]
 (1≤𝑎𝑖≤𝑛
) is given. Its element 𝑎𝑖
 is called special if there exists a pair of indices 𝑙
 and 𝑟
 (1≤𝑙<𝑟≤𝑛
) such that 𝑎𝑖=𝑎𝑙+𝑎𝑙+1+…+𝑎𝑟
. In other words, an element is called special if it can be represented as the sum of two or more consecutive elements of an array (no matter if they are special or not).

Print the number of special elements of the given array 𝑎
.

For example, if 𝑛=9
 and 𝑎=[3,1,4,1,5,9,2,6,5]
, then the answer is 5
:

𝑎3=4
 is a special element, since 𝑎3=4=𝑎1+𝑎2=3+1
;
𝑎5=5
 is a special element, since 𝑎5=5=𝑎2+𝑎3=1+4
;
𝑎6=9
 is a special element, since 𝑎6=9=𝑎1+𝑎2+𝑎3+𝑎4=3+1+4+1
;
𝑎8=6
 is a special element, since 𝑎8=6=𝑎2+𝑎3+𝑎4=1+4+1
;
𝑎9=5
 is a special element, since 𝑎9=5=𝑎2+𝑎3=1+4
.
Please note that some of the elements of the array 𝑎
 may be equal — if several elements are equal and special, then all of them should be counted in the answer.

 ### ideas
 1. 假设x是否special，如果是，所有等于x都是special的
 2. 所以就是检查是否存在sum = x的子数组
 3. 