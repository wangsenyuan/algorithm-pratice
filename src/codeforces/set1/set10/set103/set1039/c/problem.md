The Metropolis computer network consists of 𝑛
 servers, each has an encryption key in the range from 0
 to 2𝑘−1
 assigned to it. Let 𝑐𝑖
 be the encryption key assigned to the 𝑖
-th server. Additionally, 𝑚
 pairs of servers are directly connected via a data communication channel. Because of the encryption algorithms specifics, a data communication channel can only be considered safe if the two servers it connects have distinct encryption keys. The initial assignment of encryption keys is guaranteed to keep all data communication channels safe.

You have been informed that a new virus is actively spreading across the internet, and it is capable to change the encryption key of any server it infects. More specifically, the virus body contains some unknown number 𝑥
 in the same aforementioned range, and when server 𝑖
 is infected, its encryption key changes from 𝑐𝑖
 to 𝑐𝑖⊕𝑥
, where ⊕
 denotes the bitwise XOR operation.

Sadly, you know neither the number 𝑥
 nor which servers of Metropolis are going to be infected by the dangerous virus, so you have decided to count the number of such situations in which all data communication channels remain safe. Formally speaking, you need to find the number of pairs (𝐴,𝑥)
, where 𝐴
 is some (possibly empty) subset of the set of servers and 𝑥
 is some number in the range from 0
 to 2𝑘−1
, such that when all servers from the chosen subset 𝐴
 and none of the others are infected by a virus containing the number 𝑥
, all data communication channels remain safe. Since this number can be quite big, you are asked to find its remainder modulo 109+7
.

### ideas
1. 可以按位处理吗？好像必须要按位处理。但是怎么处理呢？
2. 考虑最低位，如果一条边连接了0-1的边, 
3. 如果选择server的时候，两边都包括，那么x[0]的值没有任何影响
4. 如果选择server的时候，只选择其中的一边，那么如果选择了0, 那么x[0] = 0是安全的， 如果选择x[0] = 1, 就是不安全的
5. 假设x[0]= 0时，为了安全，可以找出所有那些0-1边中0的那部分server
6. 有点乱～～～～
7. 如果不考虑是否安全，那么总数 = pow(2, sz) * (pow(2, k))
8. 任意的server组成的集合 * x所有可能的取值
9. 减去不安全的选择？
10. 比如一条边，u-v, 如果选择 x = c[v] ^ c[u] => 这条边就是安全的
11. 考虑另外一条边, a-b, c[a] != c[b], c[a] ^ x = c[b] =》 x = c[a] ^ c[b] 似乎也是有可能的
12. x = c[u] ^ c[v]， 做成一个map<x, cnt>
13. 当取值为x时，选择任意一条边，且c[u] ^ c[v] = x， 只选择u或者v时，是不安全的
14. x = c[u] ^ c[v]的这些边组成一些连通块（这些连通块肯定是二部图）
15. 只要不是选择0个或者全部，就肯定是不安全的？
16. 考虑一条线 0 - 1 - 0 - 1 - 0
17. 当x = 1时，貌似成立的