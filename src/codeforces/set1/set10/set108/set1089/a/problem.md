Alice is a big fan of volleyball and especially of the very strong "Team A".

Volleyball match consists of up to five sets. During each set teams score one point for winning a ball. The first four sets are played until one of the teams scores at least 25 points and the fifth set is played until one of the teams scores at least 15 points. Moreover, if one of the teams scores 25 (or 15 in the fifth set) points while the other team scores 24 (or 14 in the fifth set), the set is played until the absolute difference between teams' points becomes two. The match ends when one of the teams wins three sets. The match score is the number of sets won by each team.

Alice found a book containing all the results of all matches played by "Team A". The book is old, and some parts of the book became unreadable. Alice can not read the information on how many sets each of the teams won, she can not read the information on how many points each of the teams scored in each set, she even does not know the number of sets played in a match. The only information she has is the total number of points scored by each of the teams in all the sets during a single match.

Alice wonders what is the best match score "Team A" could achieve in each of the matches. The bigger is the difference between the number of sets won by "Team A" and their opponent, the better is the match score. Find the best match score or conclude that no match could end like that. If there is a solution, then find any possible score for each set that results in the best match score.

### ideas
1. dp[i][x][y] 表示在前i个sets中，A得分/B得分分别为x/y时的最优解
2. dp[i][x + 25][y + ? <= 23] = dp[i-1][x][y] + 1
3. dp[i][x + ? <= 23][y + 25] = dp[i-1][x][y] - 1
4. 