import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

test_cases = int(FILE.readline().strip())

def gen_odds(n, m):
    results = []
    for x in range(n, m+1):
        if x % 2 != 0:
            results.append(x)
    return results

for t in range(test_cases):
    a = int(FILE.readline().strip())
    b = int(FILE.readline().strip())
    print('Case {}: {}'.format(t+1, sum(gen_odds(a, b))))
