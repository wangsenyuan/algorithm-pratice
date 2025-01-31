We have
N camels numbered
1,2,…,N. Snuke has decided to make them line up in a row.

The happiness of Camel
i will be
L
i
​
if it is among the
K
i
​
frontmost camels, and
R
i
​
otherwise.

Snuke wants to maximize the total happiness of the camels. Find the maximum possible total happiness of the camel.

Solve this problem for each of the
T test cases given.

### ideas

1. 那些 li <= ri 的不需要处理（排到最后面去）
2. 只考虑那些 li > ri 的
3. 这部分按照ki升序处理，然后在其中选择 li - ri 最大的那些排列起来