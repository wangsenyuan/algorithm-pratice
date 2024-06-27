Sasha has two binary strings 𝑠
 and 𝑡
 of the same length 𝑛
, consisting of the characters 0 and 1.

There is also a computing machine that can perform two types of operations on binary strings 𝑎
 and 𝑏
 of the same length 𝑘
:

If 𝑎𝑖=𝑎𝑖+2=
 0, then you can assign 𝑏𝑖+1:=
 1 (1≤𝑖≤𝑘−2
).
If 𝑏𝑖=𝑏𝑖+2=
 1, then you can assign 𝑎𝑖+1:=
 1 (1≤𝑖≤𝑘−2
).
Sasha became interested in the following: if we consider the string 𝑎=𝑠𝑙𝑠𝑙+1…𝑠𝑟
 and the string 𝑏=𝑡𝑙𝑡𝑙+1…𝑡𝑟
, what is the maximum number of 1 characters in the string 𝑎
 that can be obtained using the computing machine. Since Sasha is very curious but lazy, it is up to you to answer this question for several pairs (𝑙𝑖,𝑟𝑖)
 that interest him.

### ideas
1. 有点乱
2. 对a/b都没有办法将1变成0，只能将0变成1
3. 且新出现的a[i+1] = 1， 是不能对b造成新的影响的
4. 但是操作的顺序是有关系的
5. 假设 a = 00000...000
6. 有办法把中间的都变成1吗？
7. 不管b一开始的状态咋样可以得到 b = ?11111?
8. 然后进一步可以得到a = 001111..1100
9. 如果 b[0] = 1, 那么a[1] = 1, 
10. 如果b[r-1] = 1, 那么 a[r-2] = 1
11. 如果a是有多段组成的 001100111类似这样子
12. 那么对于中间这些0的段，可以按照规则进行处理
13. 对于query (l, r)
14. 来说，如果a[l] = 1(未变化前), 貌似不用考虑这端的情况
15. 如果a[l] = 0 （无法变成0）也不用考虑
16. 只有到a[l] 从0变成了1
17. 这种情况发生在 b[l-1], b[l+1]将其变成了1
18. 因为l-1不在query范围内，所以它的作用就需要被忽略掉 