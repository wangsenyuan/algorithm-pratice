In the upcoming year, there will be many team olympiads, so the teachers of "T-generation" need to assemble a team of three pupils to participate in them. Any three pupils will show a worthy result in any team olympiad. But winning the olympiad is only half the battle; first, you need to get there...

Each pupil has an independence level, expressed as an integer. In "T-generation", there is exactly one student with each independence levels from 𝑙
 to 𝑟
, inclusive. For a team of three pupils with independence levels 𝑎
, 𝑏
, and 𝑐
, the value of their team independence is equal to (𝑎⊕𝑏)+(𝑏⊕𝑐)+(𝑎⊕𝑐)
, where ⊕
 denotes the bitwise XOR operation.

Your task is to choose any trio of students with the maximum possible team independence.

### ideas
1. a ^ b + b ^ c + a ^ c
2. 最高位有8种组合，除了(1, 1, 1), (0, 0, 0), 其他任何一种组合，结果都是贡献两个1
3. 其中一个可以是l（或者r）
4. 第二个数b = a + 1, 
5. 第三个数c = 从高位，在满足c == r 的情况下，尽量的和 a, b 不同就可