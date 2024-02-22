Chaneka writes down an array 𝑎
of 𝑛
positive integer elements. Initially, all elements are not circled. In one operation, Chaneka can circle an element. It
is possible to circle the same element more than once.

After doing all operations, Chaneka makes a sequence 𝑟
consisting of all uncircled elements of 𝑎
following the order of their indices.

Chaneka also makes another sequence 𝑝
such that its length is equal to the number of operations performed and 𝑝𝑖
is the index of the element that is circled in the 𝑖
-th operation.

Chaneka wants to do several operations such that sequence 𝑟
is equal to sequence 𝑝
. Help her achieve this, or report if it is impossible! Note that if there are multiple solutions, you can print any of
them.

### thoughts

1. 都不知道题目的意思～～～
2. p是有index组成的，顺序没关系，重要的是大小和次数
3. 假设下标1在p中出现了2次，那么在r中，也需要有两个1
4. 对于i, 要么i被放入了p, 要么a[i]被放入了r
5. 如果a[i]被放入了r，那么 j = a[i]被放入p一次
6. 那么a[j]就不能被放入r