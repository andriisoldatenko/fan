import sys
from array import array

FILE = sys.stdin

test_cases = int(input())
FILE.readline()


for t in range(test_cases):
    n = int(input())
    x = array('I')
    for _ in range(n):
        n1, n2 = [ord(x)-48 for x in input().split()]
        x.append(n1 + n2)
    for i in range(n-1, 0, -1):
        if x[i] >= 10:
            x[i-1] += 1
            x[i] -= 10
    print(''.join(map(str, x)))
    if t+1 != test_cases:
        input()
        print()
