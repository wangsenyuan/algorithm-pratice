Piegirl was asked to implement two table join operation for distributed database system, minimizing the network traffic.

Suppose she wants to join two tables, A and B. Each of them has certain number of rows which are distributed on
different number of partitions. Table A is distributed on the first cluster consisting of m partitions. Partition with
index i has ai rows from A. Similarly, second cluster containing table B has n partitions, i-th one having bi rows from
B.

In one network operation she can copy one row from any partition to any other partition. At the end, for each row from A
and each row from B there should be a partition that has both rows. Determine the minimal number of network operations
to achieve this.

### ideas

1. hard to understand
2. 可以得到最后只需要最多2个partition
3. 1个partition的情况，就是所有其他的都移动它上面
4. 如果是2个partition，会出现两种情况
5. x属于A，把B中的都移动过去， y属于B，把A中的都移动过去？似乎不大对
6. 如果A中还有一个额外的分区z，和B中的分区w，它们没有在一起的机会
7. 对于A中的每一个分区，a[i], 要么把a[i]都移动到某个有所有b[j]的分区上，要么把它们都移动过来