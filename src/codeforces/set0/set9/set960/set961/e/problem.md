One day Polycarp decided to rewatch his absolute favourite episode of well-known TV series "Tufurama". He was pretty surprised when he got results only for season 7 episode 3 with his search query of "Watch Tufurama season 3 episode 7 online full hd free". This got Polycarp confused — what if he decides to rewatch the entire series someday and won't be able to find the right episodes to watch? Polycarp now wants to count the number of times he will be forced to search for an episode using some different method.

TV series have n seasons (numbered 1 through n), the i-th season has ai episodes (numbered 1 through ai). Polycarp thinks that if for some pair of integers x and y (x < y) exist both season x episode y and season y episode x then one of these search queries will include the wrong results. Help Polycarp to calculate the number of such pairs!



### ideas
1. 1 a1 它的贡献 = min(n - 1, a1)
2. 2 a2,它的贡献 = min(n - 2, a2 - 1)