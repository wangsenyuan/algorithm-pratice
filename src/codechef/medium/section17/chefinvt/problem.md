## Understanding the Problem
**Problem:**
* Chef is trying to find the correct gas out of N gases to create a new light bulb.
* He uses a modulo-based search algorithm with a divisor K.
* We know the correct gas is at index p.
* We need to find how many days it will take Chef to find the correct gas.

**Algorithm:**
* Chef tests gases in groups based on their index modulo K.
* First, he tests all gases with index i such that i % K == 0.
* Then, all gases with index i such that i % K == 1, and so on.

**Goal:**
* Determine the day on which the gas at index p will be tested.

## Solution Approach
1. **Calculate the group index:**
   * Find the group to which gas p belongs by calculating `p % K`.
2. **Calculate the number of gases in previous groups:**
   * Determine how many gases have been tested before the group containing gas p.
3. **Calculate the position of p in its group:**
   * Find the position of gas p within its group.
4. **Calculate total days:**
   * Sum the days to test previous groups and the position of p in its group.

## Code Implementation (Python)
```python
def time_to_invention(N, p, K):
  """Calculates the number of days to find the correct gas.

  Args:
    N: Total number of gases.
    p: Index of the correct gas.
    K: Divisor for the modulo-based search.

  Returns:
    The number of days to find the correct gas.
  """

  group_index = p % K
  gases_before_group = group_index * (N // K)
  position_in_group = p // K + 1  # Position starts from 1
  total_days = gases_before_group + position_in_group
  return total_days
```

**Explanation:**
* `group_index`: Determines the group to which gas p belongs.
* `gases_before_group`: Calculates the number of gases tested before the group of gas p.
* `position_in_group`: Finds the position of gas p within its group.
* `total_days`: Adds the days for previous groups and the position of p.

This code effectively calculates the number of days required for Chef to find the correct gas using the given modulo-based search algorithm.
 
**Would you like to try some examples?** 



### ideas
1. 分段考虑