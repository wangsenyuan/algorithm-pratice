Randall is a software engineer at a company with N
 employees. Every year, the company re-evaluates its employees. At the end of every year, the company replaces its several worst-performing employees and replaces with the same number of new employees, so that the company keeps having N
 employees. Each person has a constant performance and can be represented by an integer (higher integer means better performance), and no two people have the same performance.

The performance of the initial employees are represented by an array of integers A=[A1,A2,…,AN]
 where Ai
 is the performance of the ith
 employee. Randall is employee 1
, so his performance is A1
. We will consider the first M
 years. At the end of the ith
 year, the company replaces its Ri
 worst-performing employees and replaces with Ri
 new employees. The performance of these new employees are represented by an array of integers Bi=[(Bi)1,(Bi)2,…,(Bi)Ri]
 where (Bi)j
 is the performance of the jth
 new employee.

He will consider Q
 scenarios. On the ith
 scenario, he will change the value of (BXi)Yi
 to Zi
. For each scenario, Randall is wondering whether he will still be in the company after M
 years. Note that the changes in each scenario are kept for the subsequent scenarios.

Input
Input begins with a line containing three integers: N
 M
 Q
 (2≤N≤100000
; 1≤M,Q≤100000
) representing the number of employees, the number of years to be considered, and the number of scenarios, respectively. The next line contains N
 integers: Ai
 (0≤Ai≤109
) representing the performance of the initial employees. The next M
 lines each contains several integers: Ri
 (Bi)1
, (Bi)2
, ⋯
, (Bi)Ri
 (1≤Ri<N
; 0≤(Bi)j≤109
) representing the number of employees replaced and the performance of the new employees, respectively. It is guaranteed that the sum of Ri
 does not exceed 106
. The next Q
 lines each contains three integers: Xi
 Yi
 Zi
 (1≤Xi≤M
; 1≤Yi≤R(Xi)
; 0≤Zi≤109
) representing a scenario. It is guaranteed that all integers in all Ai
, (Bi)j
, and Zi
 (combined together) are distinct.

 ### ideas
 1. 考虑每次replace，都是把最后的w个employee替换掉，也就是必须保证a[1]在后面的employee中
 2. 乱糟糟的一个题目
 3. 因为所有的数都是不一样的，所以a[1]在整体中的位置是确定的
 4. a[1]要能进入最后的阶段，a[1]不能被删除，那么比a[1]更大的数，也不能被删除
 5. 如果在某个阶段出现了一个比a[1]大的数，同样不能被删除
 6. 也就是说， a[1]的位置，肯定是递减的（比如一开始排在第一位，第二个月就可能排在第二位，.... 第n位）
 7. 最后必须保证至少在第n位
 8. 现在修改某个月的某个值，如果这个数变小了，那么可以计算它它a[1]的相对位置，如果他加入后，在a[1]的左边（变小没有关系）
 9. 不影响a[1]的位置，但是如果跑到后边去了，那么a[1]的后续位置都要-1, 那就有问题了
 10. range update + point query(最后一个位置即可)
 11. 不行～
 12. 