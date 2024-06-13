You have been given a matrix 𝑎
 of size 𝑛
 by 𝑚
, containing a permutation of integers from 1
 to 𝑛⋅𝑚
.

A permutation of 𝑛
 integers is an array containing all numbers from 1
 to 𝑛
 exactly once. For example, the arrays [1]
, [2,1,3]
, [5,4,3,2,1]
 are permutations, while the arrays [1,1]
, [100]
, [1,2,4,5]
 are not.

A matrix contains a permutation if, when all its elements are written out, the resulting array is a permutation. Matrices [[1,2],[3,4]]
, [[1]]
, [[1,5,3],[2,6,4]]
 contain permutations, while matrices [[2]]
, [[1,1],[2,2]]
, [[1,2],[100,200]]
 do not.

You can perform one of the following two actions in one operation:

choose columns 𝑐
 and 𝑑
 (1≤𝑐,𝑑≤𝑚
, 𝑐≠𝑑
) and swap these columns;
choose rows 𝑐
 and 𝑑
 (1≤𝑐,𝑑≤𝑛
, 𝑐≠𝑑
) and swap these rows.
You can perform any number of operations.

You are given the original matrix 𝑎
 and the matrix 𝑏
. Your task is to determine whether it is possible to transform matrix 𝑎
 into matrix 𝑏
 using the given operations.

 ### ideas
 1. 如果是一维的情况，答案当然是yes，那些不在需要位置上的数字，会组成一些component，在一个component中进行shift，就可以得到目标结果
 2. 现在是在二维的情况，要怎么处理呢？
 3. 假设a[x][y], a[u][v] 组成了一组，也就是交换它们后才能得到结果
 4. 要交换它们，必须同时交换(x, u) 和 (y, v)
 5. 此时什么东西应该保持不变性呢？
 6. 那些交换的行，必须在一个component中，也就是要组成一个环
 7. 那岂不是只要检查有没有那种只出不进，或者从进不出的就可以了？
 8. 同一行的两个数字，一个要到i行，一个要到j行，会怎样？
 9. 因为列的变化，不会改变数字在同一行的这样一个事实。通过行变化，又始终在同一行
 10. 所以只需要检查是否可以退化到一维就可以了 