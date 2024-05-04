Recenlty Luba got a credit card and started to use it. Let's consider n consecutive days Luba uses the card.

She starts with 0 money on her account.

In the evening of i-th day a transaction ai occurs. If ai>0, then ai bourles are deposited to Luba's account. If ai<0,
then ai bourles are withdrawn. And if ai = 0, then the amount of money on Luba's account is checked.

In the morning of any of n days Luba can go to the bank and deposit any positive integer amount of burles to her
account. But there is a limitation: the amount of money on the account can never exceed d.

It can happen that the amount of money goes greater than d by some transaction in the evening. In this case answer will
be «-1».

Luba must not exceed this limit, and also she wants that every day her account is checked (the days when ai = 0) the
amount of money on her account is non-negative. It takes a lot of time to go to the bank, so Luba wants to know the
minimum number of days she needs to deposit some money to her account (if it is possible to meet all the requirements).
Help her!

### ideas

1. 假设在第i天结束后，有金额x (x < 0, 表示欠款， x >= 0 表示存款)
2. 如果 a[i+1] = 0, 说明这一天要check， 所以必须存入一笔钱以保证 x >= 0，但是这个存入的钱目前还是不知道的
3. 如果 a[i+1] != 0, x += a[i+1]
4. 在存入钱的时候，要保证 x <= d
5. 那么是不是只要在check的那天存入钱，就好了？
6. 因为反正可以存入任意的金额