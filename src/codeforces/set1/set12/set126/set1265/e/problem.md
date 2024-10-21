Creatnx has 𝑛
 mirrors, numbered from 1
 to 𝑛
. Every day, Creatnx asks exactly one mirror "Am I beautiful?". The 𝑖
-th mirror will tell Creatnx that he is beautiful with probability 𝑝𝑖100
 for all 1≤𝑖≤𝑛
.

Creatnx asks the mirrors one by one, starting from the 1
-st mirror. Every day, if he asks 𝑖
-th mirror, there are two possibilities:

The 𝑖
-th mirror tells Creatnx that he is beautiful. In this case, if 𝑖=𝑛
 Creatnx will stop and become happy, otherwise he will continue asking the 𝑖+1
-th mirror next day;
In the other case, Creatnx will feel upset. The next day, Creatnx will start asking from the 1
-st mirror again.
You need to calculate the expected number of days until Creatnx becomes happy.

This number should be found by modulo 998244353
. Formally, let 𝑀=998244353
. It can be shown that the answer can be expressed as an irreducible fraction 𝑝𝑞
, where 𝑝
 and 𝑞
 are integers and 𝑞≢0(mod𝑀)
. Output the integer equal to 𝑝⋅𝑞−1mod𝑀
. In other words, output such an integer 𝑥
 that 0≤𝑥<𝑀
 and 𝑥⋅𝑞≡𝑝(mod𝑀)
.

### ideas
1. 