Vasya has an array of integers of length n.

Vasya performs the following operations on the array: on each step he finds the longest segment of consecutive equal
integers (the leftmost, if there are several such segments) and removes it. For example, if Vasya's array
is [13, 13, 7, 7, 7, 2, 2, 2], then after one operation it becomes [13, 13, 2, 2, 2].

Compute the number of operations Vasya should make until the array becomes empty, i.e. Vasya removes all elements from
it.

### ideas

1. 每个位置最多被删除一次，
2. 所以问题的关键是快速的找到最长（最小的）区间
3. 删除掉一个区间后，可能发生两种情况， 不变（前后不一样）或者合并
4. union find + priority queue