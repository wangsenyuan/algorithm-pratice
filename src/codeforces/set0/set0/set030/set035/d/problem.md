Once upon a time DravDe, an outstanding person famous for his professional achievements (as you must remember, he works in a warehouse storing Ogudar-Olok, a magical but non-alcoholic drink) came home after a hard day. That day he had to drink 9875 boxes of the drink and, having come home, he went to bed at once.

DravDe dreamt about managing a successful farm. He dreamt that every day one animal came to him and asked him to let it settle there. However, DravDe, being unimaginably kind, could send the animal away and it went, rejected. There were exactly n days in DravDe’s dream and the animal that came on the i-th day, ate exactly ci tons of food daily starting from day i. But if one day the animal could not get the food it needed, it got really sad. At the very beginning of the dream there were exactly X tons of food on the farm.

DravDe woke up terrified...

When he retold the dream to you, he couldn’t remember how many animals were on the farm by the end of the n-th day any more, but he did remember that nobody got sad (as it was a happy farm) and that there was the maximum possible amount of the animals. That’s the number he wants you to find out.

It should be noticed that the animals arrived in the morning and DravDe only started to feed them in the afternoon, so that if an animal willing to join them is rejected, it can’t eat any farm food. But if the animal does join the farm, it eats daily from that day to the n-th.

### ideas
1. dp[i][j] = 在suf[i:]中，保留j个时，的最小花费