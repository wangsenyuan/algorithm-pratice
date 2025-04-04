As meticulous Gerald sets the table, Alexander finished another post on Codeforces and begins to respond to New Year greetings from friends. Alexander has n friends, and each of them sends to Alexander exactly one e-card. Let us number his friends by numbers from 1 to n in the order in which they send the cards. Let's introduce the same numbering for the cards, that is, according to the numbering the i-th friend sent to Alexander a card number i.

Alexander also sends cards to friends, but he doesn't look for the new cards on the Net. He simply uses the cards previously sent to him (sometimes, however, he does need to add some crucial details). Initially Alexander doesn't have any cards. Alexander always follows the two rules:

He will never send to a firend a card that this friend has sent to him.
Among the other cards available to him at the moment, Alexander always chooses one that Alexander himself likes most.
Alexander plans to send to each friend exactly one card. Of course, Alexander can send the same card multiple times.

Alexander and each his friend has the list of preferences, which is a permutation of integers from 1 to n. The first number in the list is the number of the favorite card, the second number shows the second favorite, and so on, the last number shows the least favorite card.

Your task is to find a schedule of sending cards for Alexander. Determine at which moments of time Alexander must send cards to his friends, to please each of them as much as possible. In other words, so that as a result of applying two Alexander's rules, each friend receives the card that is preferred for him as much as possible.

Note that Alexander doesn't choose freely what card to send, but he always strictly follows the two rules.


### ideas
1. 朋友依次送card给A，在时刻i，A拥有1，2，..., i 这些cards
2. 在任何一个时刻，都应该根据A目前最喜欢的（在前i个中），然后去更新所有人的答案
3. 模拟就可以了