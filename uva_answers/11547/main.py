import sys

FILE = sys.stdin
# FILE = open('sample.in')

test_cases = range(int(FILE.readline()))

for tc in test_cases:
    n = int(FILE.readline().strip())
    res = ((n * 63 + 7492) * 5 - 498)
    print(str(res)[-2])
