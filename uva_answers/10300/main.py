import pprint
import sys

FILE = sys.stdin
# FILE = open('sample.in')


test_cases = int(FILE.readline().strip())
for tc in range(test_cases):
    farmers = int(FILE.readline().strip())
    total_sum = 0
    for farmer in range(farmers):
        s, a, c = map(int, FILE.readline().strip().split())
        total_sum += s * c
    print(total_sum)
