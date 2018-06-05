import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

test_cases = int(FILE.readline().strip())

for t in range(test_cases):
    a = sorted(map(int, FILE.readline().strip().split()))
    print('Case {}: {}'.format(t+1, a[1]))
