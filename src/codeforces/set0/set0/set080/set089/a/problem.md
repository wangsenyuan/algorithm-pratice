It is nighttime and Joe the Elusive got into the country's main bank's safe. The safe has n cells positioned in a row, each of them contains some amount of diamonds. Let's make the problem more comfortable to work with and mark the cells with positive numbers from 1 to n from the left to the right.

Unfortunately, Joe didn't switch the last security system off. On the plus side, he knows the way it works.

Every minute the security system calculates the total amount of diamonds for each two adjacent cells (for the cells between whose numbers difference equals 1). As a result of this check we get an n - 1 sums. If at least one of the sums differs from the corresponding sum received during the previous check, then the security system is triggered.

Joe can move the diamonds from one cell to another between the security system's checks. He manages to move them no more than m times between two checks. One of the three following operations is regarded as moving a diamond: moving a diamond from any cell to any other one, moving a diamond from any cell to Joe's pocket, moving a diamond from Joe's pocket to any cell. Initially Joe's pocket is empty, and it can carry an unlimited amount of diamonds. It is considered that before all Joe's actions the system performs at least one check.

In the morning the bank employees will come, which is why Joe has to leave the bank before that moment. Joe has only k minutes left before morning, and on each of these k minutes he can perform no more than m operations. All that remains in Joe's pocket, is considered his loot.

Calculate the largest amount of diamonds Joe can carry with him. Don't forget that the security system shouldn't be triggered (even after Joe leaves the bank) and Joe should leave before morning.


### ideas
1. 如果 m < n - 1 是不是0?
2. 假设序列是 [a, b, c] 那么需要维持 (a + b, b + c)的条件不变
3. 如果从(a + b)中取出来了 x, 那么必须从其他地方增加x到a+b中去
4. a + b, b + c
5. 如果从a中取出x到b中, c中取出y个放到b中
6. a - x + b + x + y = a + b + y 不可以
7. a中的只能放置到b中 a - x + b + x
8. b + x + c - x 然后将c中取走x个如果它后面有的话，再放入后面的位置。
9. 否则的话，必须放入pocket
10. 这是一个可行的策略，但是是唯一的策略吗？如果是唯一的策略，那么n位偶数时 =》 答案为0， 因为没有机会将diamond装pocket
11. 如果是从左往右，那么第一个的必须迁入第二个，a - x + b + x 才能保证 第一个+第二个的和不变
12. 现在第二个变成了 b + x, 那么为了保证第二个+第三个和不变（b + x + c - y) 不变， y必须= x
13. 但是如果把y装入pocket会怎么样？
14. b + x + c - x 满足条件，但是 c - x + d + x，表示必须从后边迁入x到d中
15. 一次只能移动一个diamond
16. [a, b, c, d, e], min(k, a, c, e) ?