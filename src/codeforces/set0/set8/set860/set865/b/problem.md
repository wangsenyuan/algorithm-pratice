It's another Start[c]up finals, and that means there is pizza to order for the onsite contestants. There are only 2
types of pizza (obviously not, but let's just pretend for the sake of the problem), and all pizzas contain exactly S
slices.

It is known that the i-th contestant will eat si slices of pizza, and gain ai happiness for each slice of type 1 pizza
they eat, and bi happiness for each slice of type 2 pizza they eat. We can order any number of type 1 and type 2 pizzas,
but we want to buy the minimum possible number of pizzas for all of the contestants to be able to eat their required
number of slices. Given that restriction, what is the maximum possible total happiness that can be achieved?

Input
The first line of input will contain integers N and S (1 ≤ N ≤ 105, 1 ≤ S ≤ 105), the number of contestants and the
number of slices per pizza, respectively. N lines follow.

The i-th such line contains integers si, ai, and bi (1 ≤ si ≤ 105, 1 ≤ ai ≤ 105, 1 ≤ bi ≤ 105), the number of slices the
i-th contestant will eat, the happiness they will gain from each type 1 slice they eat, and the happiness they will gain
from each type 2 slice they eat, respectively.

Output
Print the maximum total happiness that can be achieved.

### ideas

1. 假设总数是sum = s[1] + s[2] + ... + s[n]
2. pizza的数量 x = (sum + S - 1) / S
3. x是定的，然后考虑其中类型1的有u个， 那么 类型2的有v个
4. 然后按照 s[i] * (a[i] - b[i])排序，前面的用类型1，后面的用类型2
5. u * sum (s[i] * a[i]) + (x - u) * sum (s[j] * b[j])
6. 直觉上，应该中间某个点，达到最大值