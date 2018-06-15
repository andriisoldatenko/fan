import sys

FILE = sys.stdin
# FILE = open('sample.in')

test_cases = range(int(FILE.readline()))

for tc in test_cases:
    # number of stores to visit
    n = int(FILE.readline().strip())

    # positions on Long street
    x = list(map(int, FILE.readline().strip().split()))
    print(2 * (max(x) - min(x)))
