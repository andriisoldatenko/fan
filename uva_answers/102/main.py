import functools
import sys

@functools.lru_cache(maxsize=1000000, typed=True)
def cycle_length(n):
    if n == 1:
        return 1
    cnt = 1
    while n != 1:
        if n & 1 == 1:
            n = n + n + n + 1
        else:
            n = n >> 1
        cnt += 1
    return cnt

for line in sys.stdin:
    cl_line = line.strip()
    a, b = map(int, cl_line.split())
    max_cycle_length = 0
    for x in range(min(a, b), max(a, b) + 1):
        max_cycle_length = max(cycle_length(x), max_cycle_length)
    print('{} {} {}\n'.format(a, b, max_cycle_length), end='')
