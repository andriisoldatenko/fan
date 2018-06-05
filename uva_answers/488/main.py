import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

test_cases = int(FILE.readline().strip())


def print_wave(n):
    for i in range(1, n+1):
        print(str(i)*i)
    for i in reversed(range(1, n)):
        print(str(i)*i)
for t in range(test_cases):
    FILE.readline().strip()
    a = int(FILE.readline().strip())
    f = int(FILE.readline().strip())
    for i in range(f):
        print_wave(a)
        if (t == test_cases - 1) and (i == f-1):
            continue
        print()
