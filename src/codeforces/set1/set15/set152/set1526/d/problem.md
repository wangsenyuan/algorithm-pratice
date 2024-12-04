After rejecting 10100
 data structure problems, Errorgorn is very angry at Anton and decided to kill him.

Anton's DNA can be represented as a string 𝑎
 which only contains the characters "ANTON" (there are only 4
 distinct characters).

Errorgorn can change Anton's DNA into string 𝑏
 which must be a permutation of 𝑎
. However, Anton's body can defend against this attack. In 1
 second, his body can swap 2
 adjacent characters of his DNA to transform it back to 𝑎
. Anton's body is smart and will use the minimum number of moves.

To maximize the chance of Anton dying, Errorgorn wants to change Anton's DNA the string that maximizes the time for Anton's body to revert his DNA. But since Errorgorn is busy making more data structure problems, he needs your help to find the best string 𝐵
. Can you help him?



### ideas
1. 首先要知道如何计算两个字符串a, b的最少交换次数
2. ANTON vs NNOTA
3. A移动到第一位需要4秒  NTON vs NNOT
4. 0 => TON NOT
5. 2 => ON NO
6. 1 => N, N
7. 最小的次数，就是当前字符，（右边）最靠近的位置进行交换
8. 直接reverse是不行的。
9. 如果把第一个字符放置在j后，那么所有同样的字符都应该放置在j的后面（否则的话，后面同一个字符放置在j的前面，还不如就让第一个字符放置在那个位置）
10. pos[i] = min(i + diff, 当前最靠右的，且有count(a[i])个空位的地方)
11. pos[1] = min(1 + 4, n) = 5 ????A
12. pos[2] = min(2 + 2, 3) = 3 ??NNA
13.                            OTNNA
14.  ANTON vs OTNNA
15.  4 => NTON OTNN
16.  2 => TON OTN
17.  1 => ON ON
18.  NAAN
19.  pos[1] = ??NN 
20.  OAANTTON 