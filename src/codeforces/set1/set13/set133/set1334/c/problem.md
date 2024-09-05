The first line contains one integer 𝑇
 (1≤𝑇≤150000
) — the number of test cases.

Then the test cases follow, each test case begins with a line containing one integer 𝑛
 (2≤𝑛≤300000
) — the number of monsters. Then 𝑛
 lines follow, each containing two integers 𝑎𝑖
 and 𝑏𝑖
 (1≤𝑎𝑖,𝑏𝑖≤1012
) — the parameters of the 𝑖
-th monster in the circle.

It is guaranteed that the total number of monsters in all test cases does not exceed 300000
.

### ideas
1. 显然要依次slay怪物（如果在中间先干掉一个怪物，有可能打断explosion，使得伤害减低）且可以使用交换论证，证明依次处理更优
2. 但是，选择哪个开始是有关系的
3. 如果从0，开始，那么可以计算出，需要使用多少次射击（使用贪心的方式）
4. 但是如果从1开始呢？
5. 假设dp是从当前开始射击的结果；（从0开始的贪心计算出来）
6. 然后移动到下一个位置；
7. 如果下一个位置，在当前的处理中，不是通过爆炸被消灭的，那么就比较好计算，假设在它开始时，上个节点对它造成的伤害时b[i], 
8. 那么这b[i]个伤害，必须通过一枪枪打出来；如果上一个节点是被爆炸伤害致死的（似乎也只需要额外b[i]的伤害即可）
9. 还有就是最后一个i-1对i的伤害