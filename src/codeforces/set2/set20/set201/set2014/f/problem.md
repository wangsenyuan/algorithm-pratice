The folk hero Robin Hood has been troubling Sheriff of Nottingham greatly. Sheriff knows that Robin Hood is about to attack his camps and he wants to be prepared.

Sheriff of Nottingham built the camps with strategy in mind and thus there are exactly 𝑛
 camps numbered from 1
 to 𝑛
 and 𝑛−1
 trails, each connecting two camps. Any camp can be reached from any other camp. Each camp 𝑖
 has initially 𝑎𝑖
 gold.

As it is now, all camps would be destroyed by Robin. Sheriff can strengthen a camp by subtracting exactly 𝑐
 gold from each of its neighboring camps and use it to build better defenses for that camp. Strengthening a camp doesn't change its gold, only its neighbors' gold. A camp can have negative gold.

After Robin Hood's attack, all camps that have been strengthened survive the attack, all others are destroyed.

What's the maximum gold Sheriff can keep in his surviving camps after Robin Hood's attack if he strengthens his camps optimally?

Camp 𝑎
 is neighboring camp 𝑏
 if and only if there exists a trail connecting 𝑎
 and 𝑏
. Only strengthened camps count towards the answer, as others are destroyed.

### ideas
1. 看起来所有的都可以被保护起来，但是这样子，需要花费太多，如果c比较少，也可能是值得的
2. dp[u] 表示在子树u中被保护后获得的最大值
3. 如果要计算u的状态，是需要知道v的状态的，（v是u的直接子节点）
4. 因为如果v被保护了，那么a[u] - c, 也就是说如果u被保护起来后，得到的金额也是 dp[v] + a[u] - c
5. dp[u][1] 表示父节点被保护时的状态，
6. dp[u][0] 表示父节点被放弃时的状态
7. dp[u][0] = max(sum of dp[v][0], a[u] + sum of dp[v][1])
8. 因为父节点未被保护，所以它对u没有影响。 如果u也没有保护（dp[v][0]中的0），
9. dp[u][1] = max(sum of dp[v][0]) u不被保护
10.           sum of dp[v][1] + a[u] - c 如果u被保护时，（p也被保护，所以要把p的收益给减去c）