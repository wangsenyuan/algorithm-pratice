https://codeforces.com/problemset/problem/1304/C

Gildong owns a bulgogi restaurant. The restaurant has a lot of customers, so many of them like to make a reservation
before visiting it.

Gildong tries so hard to satisfy the customers that he even memorized all customers' preferred temperature ranges!
Looking through the reservation list, he wants to satisfy all customers by controlling the temperature of the
restaurant.

The restaurant has an air conditioner that has 3 states: off, heating, and cooling. When it's off, the restaurant's
temperature remains the same. When it's heating, the temperature increases by 1 in one minute. Lastly, when it's
cooling, the temperature decreases by 1 in one minute. Gildong can change the state as many times as he wants, at any
integer minutes. The air conditioner is off initially.

Each customer is characterized by three values: 𝑡𝑖
— the time (in minutes) when the 𝑖
-th customer visits the restaurant, 𝑙𝑖
— the lower bound of their preferred temperature range, and ℎ𝑖
— the upper bound of their preferred temperature range.

A customer is satisfied if the temperature is within the preferred range at the instant they visit the restaurant.
Formally, the 𝑖
-th customer is satisfied if and only if the temperature is between 𝑙𝑖
and ℎ𝑖
(inclusive) in the 𝑡𝑖
-th minute.

Given the initial temperature, the list of reserved customers' visit times and their preferred temperature ranges,
you're going to help him find if it's possible to satisfy all customers.