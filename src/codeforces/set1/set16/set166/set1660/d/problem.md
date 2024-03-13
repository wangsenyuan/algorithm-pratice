You are given an array 𝑎
consisting of 𝑛
integers. For each 𝑖
(1≤𝑖≤𝑛
) the following inequality is true: −2≤𝑎𝑖≤2
.

You can remove any number (possibly 0
) of elements from the beginning of the array and any number (possibly 0
) of elements from the end of the array. You are allowed to delete the whole array.

You need to answer the question: how many elements should be removed from the beginning of the array, and how many
elements should be removed from the end of the array, so that the result will be an array whose product (multiplication)
of elements is maximal. If there is more than one way to get an array with the maximum product of elements on it, you
are allowed to output any of them.

The product of elements of an empty array (array of length 0
) should be assumed to be 1
.

### ideas

1. 要把全部的0删除掉，因为包含0，结果就为0，空串的结果为1
2. 删除0有三种方式，一类是前部全部删除，一类是后面全部删除，一类是取中间的部分（各段）
3. 所以就是取不含0的每段，计算它们能够得到的最大值