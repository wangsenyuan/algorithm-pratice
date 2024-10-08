Two integers x and y are compatible, if the result of their bitwise "AND" equals zero, that is, a & b = 0. For example, numbers 90 (10110102) and 36 (1001002) are compatible, as 10110102 & 1001002 = 02, and numbers 3 (112) and 6 (1102) are not compatible, as 112 & 1102 = 102.

You are given an array of integers a1, a2, ..., an. Your task is to find the following for each array element: is this element compatible with some other element from the given array? If the answer to this question is positive, then you also should find any suitable element.

Input
The first line contains an integer n (1 ≤ n ≤ 106) — the number of elements in the given array. The second line contains n space-separated integers a1, a2, ..., an (1 ≤ ai ≤ 4·106) — the elements of the given array. The numbers in the array can coincide.


### solution

Consider some number x from the array. Inverse all bits in x and say it is number y. Consider an integer a[i] from array. It can be an answer to the number x if for every position of zero bit from y there is zero bit in a[i] in the same position. Other bits in a[i] we can change to ones.

Then we will use such dynamic programming z[mask] = {0, 1} which means if we can change some zero bits to ones in some integer from given array a and get mask mask. Initial states are all integers from array a (z[a[i]] = 1). To go from one state to another we consider every bit in mask and try to change it to one. The length of bit representation of all integers in a is less or equal than 22.

To answer the question YES or NO for some number x you need to get value [z(y)&(1«22) - 1] (inverse all bits in x and make the length of the number 22). If you want to know the exact answer what number a[i] you should choose you can save previous states for every state z[mask] or just save it in z[mask]. Complexity O(2K * K), where K – length of bit representation of numbers (K <  = 22).