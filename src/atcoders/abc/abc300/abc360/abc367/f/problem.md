Given sequences of positive integers of length N:
A=(A1​,A2​,…,AN​)
B=(B1​,B2​,…,BN​)
You are given Q queries to process in order. The i-th query is as follows:
You are given positive integers li​,ri​,Li​,Ri​. Print "Yes" if it is possible to rearrange the subsequence (Ali​​,Ali​+1​,…,Ari​​) to match the subsequence (BLi​​,BLi​+1​,…,BRi​​), and "No" otherwise.

### ideas
1. 显然两个子串的长度要一致，频度要一致
2. 