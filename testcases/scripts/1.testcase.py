import math
import random

max = int(math.pow(10, 9))
min = -max
target = random.randint(2*min, 2*max)
r1 = random.randint(min, max)
r2 = target - r1
if r2 <= min or r2 >= max :
    r2 = target + r1

nums = [r1,r2]
mem = {r1,r2}

for i in range(10000-2):
    r = random.randint(min, max)
    while target - r in mem:
        r = random.randint(min, max)
    mem.add(r)
    nums.append(r)

random.shuffle(nums)
print(nums)
print(target)