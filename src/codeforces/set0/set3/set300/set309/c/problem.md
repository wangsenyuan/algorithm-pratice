You get to work and turn on the computer. You start coding and give little thought to the RAM role in the whole process.
In this problem your task is to solve one of the problems you encounter in your computer routine.

We'll consider the RAM as a sequence of cells that can contain data. Some cells already contain some data, some are
empty. The empty cells form the so-called memory clusters. Thus, a memory cluster is a sequence of some consecutive
empty memory cells.

You have exactly n memory clusters, the i-th cluster consists of ai cells. You need to find memory for m arrays in your
program. The j-th array takes 2bj consecutive memory cells. There possibly isn't enough memory for all m arrays, so your
task is to determine what maximum number of arrays can be located in the available memory clusters. Of course, the
arrays cannot be divided between the memory clusters. Also, no cell can belong to two arrays.

Input
The first line of the input contains two integers n and m (1 ≤ n, m ≤ 106). The next line contains n integers a1,
a2, ..., an (1 ≤ ai ≤ 109). The next line contains m integers b1, b2, ..., bm (1 ≤ 2bi ≤ 109).

### thoughts

1. 一个memory cluster可以分配多个array，要求它们的sum <= a[x]
2. bi <= 29,
3. 这里是要满足最多的b的数量，另外一个就是如果b1 = 18, b2 = 19, 那么满足一个b2， 可以满足两个b1
4. 所以，贪心的结果是（对p排序后），肯定满足前i个最小的数组
5. 对a也进行排序，
6. 二分检查是否可以对前i个b进行分配
7. 但怎么检查呢？没想法～
8. 是不是把最大的拿来分配更合理，然后把剩余的放进去？直觉上感觉可行，但怎么证明呢？
9. 最大的a[n]在某个最优的答案里它必然包括了最大的数组
10. 假设没有，如果最大的数组分配在a[i]中，如果a[n]有足够的空间，把最大的数组移过来不改变结果
11. 假设a[n]没有足够的空间，但是必然存在一对数组的sum <= 最大的数组，不失一般性，假设只有一个，
12. 通过交换它们，总是能成功，且不会改变结果
