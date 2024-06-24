Rudolf has prepared a set of 𝑛
 problems with complexities 𝑎1<𝑎2<𝑎3<⋯<𝑎𝑛
. He is not entirely satisfied with the balance, so he wants to add at most one problem to fix it.

For this, Rudolf came up with 𝑚
 models of problems and 𝑘
 functions. The complexity of the 𝑖
-th model is 𝑑𝑖
, and the complexity of the 𝑗
-th function is 𝑓𝑗
. To create a problem, he selects values 𝑖
 and 𝑗
 (1≤𝑖≤𝑚
, 1≤𝑗≤𝑘
) and by combining the 𝑖
-th model with the 𝑗
-th function, he obtains a new problem with complexity 𝑑𝑖+𝑓𝑗
 (a new element is inserted into the array 𝑎
).

To determine the imbalance of the set, Rudolf sorts the complexities of the problems in ascending order and finds the largest value of 𝑎𝑖−𝑎𝑖−1
 (𝑖>1
).

What is the minimum value of imbalance that Rudolf can achieve by adding at most one problem, created according to the described rules?

### ideas
1. just find the two largest difference, and check if can decrease the largest one by using some d and f