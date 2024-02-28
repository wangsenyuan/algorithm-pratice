Some girl has stolen Barney's heart, and Barney wants to find her. He starts looking for in the root of the tree and (
since he is Barney Stinson not a random guy), he uses a random DFS to search in the cities. A pseudo code of this
algorithm is as follows:

```c
    
    let starting_time be an array of length n
    current_time = 0
    dfs(v):
        current_time = current_time + 1
        starting_time[v] = current_time
        shuffle children[v] randomly (each permutation with equal possibility)
        // children[v] is vector of children cities of city v
        for u in children[v]:
            dfs(u)
```

As told before, Barney will start his journey in the root of the tree (equivalent to call dfs(1)).

Now Barney needs to pack a backpack and so he wants to know more about his upcoming journey: for every city i, Barney
wants to know the expected value of starting_time[i]. He's a friend of Jon Snow and knows nothing, that's why he asked
for your help.

### thoughts

1. 期望值 = 到达这里的所有的starting的总和/到达这里所有的路线的计数
2. dp[u] = 到达u的期望值, 那么对于u的children， dp[v] = dp[u] + (sz[u] - sz[v] - 1) / 2