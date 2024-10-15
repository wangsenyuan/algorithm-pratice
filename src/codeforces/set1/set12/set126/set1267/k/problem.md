Karl is developing a key storage service. Each user has a positive integer key.

Karl knows that storing keys in plain text is bad practice. So, instead of storing a key, he decided to store a fingerprint of a key. However, using some existing fingerprint algorithm looked too boring to him, so he invented his own one.

Karl's fingerprint is calculated by the following process: divide the given integer by 2, then divide the result by 3, then divide the result by 4, and so on, until we get a result that equals zero (we are speaking about integer division each time). The fingerprint is defined as the multiset of the remainders of these divisions.

For example, this is how Karl's fingerprint algorithm is applied to the key 11: 11 divided by 2 has remainder 1 and result 5, then 5 divided by 3 has remainder 2 and result 1, and 1 divided by 4 has remainder 1 and result 0. Thus, the key 11 produces the sequence of remainders [1,2,1]
 and has the fingerprint multiset {1,1,2}
.

Ksenia wants to prove that Karl's fingerprint algorithm is not very good. For example, she found that both keys 178800 and 123456 produce the fingerprint of {0,0,0,0,2,3,3,4}
. Thus, users are at risk of fingerprint collision with some commonly used and easy to guess keys like 123456.

Ksenia wants to make her words more persuasive. She wants to calculate the number of other keys that have the same fingerprint as the keys in the given list of some commonly used keys. Your task is to help her.

### ideas
1. 完全没有头绪
2. 给定k的fingerprint， 可以计算出来。但是如何计算剩余的呢？