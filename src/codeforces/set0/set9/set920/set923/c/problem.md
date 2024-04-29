Alice has a very important message M consisting of some non-negative integers that she wants to keep secret from Eve.
Alice knows that the only theoretically secure cipher is one-time pad. Alice generates a random key K of the length
equal to the message's length. Alice computes the bitwise xor of each element of the message and the key (, where
denotes the bitwise XOR operation) and stores this encrypted message A. Alice is smart. Be like Alice.

For example, Alice may have wanted to store a message M =(0, 15, 9, 18). She generated a key K =(16, 7, 6, 3). The
encrypted message is thus A =(16, 8, 15, 17).

Alice realised that she cannot store the key with the encrypted message. Alice sent her key K to Bob and deleted her own
copy. Alice is smart. Really, be like Alice.

Bob realised that the encrypted message is only secure as long as the key is secret. Bob thus randomly permuted the key
before storing it. Bob thinks that this way, even if Eve gets both the encrypted message and the key, she will not be
able to read the message. Bob is not smart. Don't be like Bob.

In the above example, Bob may have, for instance, selected a permutation (3, 4, 1, 2) and stored the permuted key P =(6,
3, 16, 7).

One year has passed and Alice wants to decrypt her message. Only now Bob has realised that this is impossible. As he has
permuted the key randomly, the message is lost forever. Did we mention that Bob isn't smart?

Bob wants to salvage at least some information from the message. Since he is not so smart, he asks for your help. You
know the encrypted message A and the permuted key P. What is the lexicographically smallest message that could have
resulted in the given encrypted text?

More precisely, for given A and P, find the lexicographically smallest message O, for which there exists a permutation π
such that for every i.

Note that the sequence S is lexicographically smaller than the sequence T, if there is an index i such that Si<Ti and
for all j<i the condition Sj = Tj holds.

### ideas

1. 稍微有点绕，就是要找到这样一个序列，O，O[i] ^ pi(P[i]) = A[i]
2. 且使得O是一个最小的序列
3. 如果有让O[0] = 0 的， 那么就要选择这样的P[i]
4. 首先这里有两个问题，
5. 第一个问题是，有可能不存在这样的P[i], 也不能1，2，3.。。的尝试？
6. 第二个问题是，出现了多个，比如 P[i], P[j], ...
7. 第一个问题，应该不存在，因为所有的P，都可以和A[0] xor，并有一个确定的值，只需要在其中选择最小的即可
8. 第二个也不是问题，选择其中一个即可
9. 所以，剩下的问题，就是如何快速的找到那个最小的，就是要和A[0]尽量的接近的值