import sys
from collections import OrderedDict

FILE = sys.stdin
# FILE = open('sample.in')
# inpt = FILE.readline().strip()

test_case = 1
test_cases = range(1, int(input())+1)

for tc in test_cases:
    pages = OrderedDict()
    for site in range(10):
        l, r = FILE.readline().strip().split()
        pages[l] = int(r)
    max_rank = max(pages.values())
    print('Case #{}:'.format(tc))
    for k,v in pages.items():
        if v == max_rank:
            print(k)
