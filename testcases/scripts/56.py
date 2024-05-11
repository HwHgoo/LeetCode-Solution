import random

def generate_pair() -> list[int]:
    pair = []
    start = random.randint(0, 10000)
    end = random.randint(start, 10000)
    pair.append(start)
    pair.append(end)
    return pair

N = 10000
def genarate_intervals() -> list[list[int]]:
    intervals = []
    for i in range(N):
        intervals.append(generate_pair())
    
    return intervals

intervals = genarate_intervals()
print(intervals)