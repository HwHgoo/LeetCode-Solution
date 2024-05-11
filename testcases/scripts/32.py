import random
import time

N = int(3 * 1e4)
print(N)

s = ""
for i in range(N):
    random.seed(time.time_ns())
    ri = random.randint(1,2)
    if ri == 1:
        s += "("
    else:
        s += ")"

print(s)