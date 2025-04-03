Bob is about to take a hot bath.

There are two taps to fill the bath: a hot water tap and a cold water tap. The cold water's temperature is t1, and the hot water's temperature is t2. The cold water tap can transmit any integer number of water units per second from 0 to x1, inclusive. Similarly, the hot water tap can transmit from 0 to x2 water units per second.

If y1 water units per second flow through the first tap and y2 water units per second flow through the second tap, then the resulting bath water temperature will be:


Bob wants to open both taps so that the bath water temperature was not less than t0. However, the temperature should be as close as possible to this value. If there are several optimal variants, Bob chooses the one that lets fill the bath in the quickest way possible.

Determine how much each tap should be opened so that Bob was pleased with the result in the end.

### ideas
1. 假设答案是 y1, y2, 
2. t = (t1 * y1 + t2 * y2) / (y1 + y2) >= t0
3. 在给定y1的情况下，计算出y2, 
4. 要求y1+y2越大越好