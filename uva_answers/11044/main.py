import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

test_cases = range(int(FILE.readline()))
#import ipdb; ipdb.set_trace()
for tc in test_cases:
    n, m = map(int, FILE.readline().strip().split())
    print((n//3) * (m//3))
