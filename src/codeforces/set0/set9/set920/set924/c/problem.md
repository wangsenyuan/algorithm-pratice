Arkady decides to observe a river for n consecutive days. The river's water level on each day is equal to some real
value.

Arkady goes to the riverside each day and makes a mark on the side of the channel at the height of the water level, but
if it coincides with a mark made before, no new mark is created. The water does not wash the marks away. Arkady writes
down the number of marks strictly above the water level each day, on the i-th day this value is equal to mi.

Define di as the number of marks strictly under the water level on the i-th day. You are to find out the minimum
possible sum of di over all days. There are no marks on the channel before the first day.

### ideas

1. 理解一下题目，m[i]表示在第i天，在水位上面的mark
2. d[i]表示在第i天，在水位下面的mark，求最小的d的和
3. 有一个不变的点，就是前面水位的条数是不会变的
4. 假设到目前为止出现了x个mark，当前情况下，如果m[i] <= x, 那么可以不产生新的mark
5. 如果 m[i] == x + 1, 那么就产生一个新的mark (也有可能是前面缺了一个mark)
6. 如果 m[i] > x + 1， 那么前面至少要补 m[i] - x - 1 个mark
7. x = 多少呢？= max(m[i]) - min(m[i]) + 1
8. 