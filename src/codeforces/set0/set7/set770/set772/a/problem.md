You have n devices that you want to use simultaneously.

The i-th device uses ai units of power per second. This usage is continuous. That is, in λ seconds, the device will use
λ·ai units of power. The i-th device currently has bi units of power stored. All devices can store an arbitrary amount
of power.

You have a single charger that can plug to any single device. The charger will add p units of power per second to a
device. This charging is continuous. That is, if you plug in a device for λ seconds, it will gain λ·p units of power.
You can switch which device is charging at any arbitrary unit of time (including real numbers), and the time it takes to
switch is negligible.

You are wondering, what is the maximum amount of time you can use the devices until one of them hits 0 units of power.

If you can use the devices indefinitely, print -1. Otherwise, print the maximum amount of time before any one device
hits 0 power.

### ideas

1. binary search expect
2. 如果共用了m分钟，那么对于每一个device (a, b) b + p * x - a * expect >= 0
3. => x >= (a * expect - b) / p
4. 