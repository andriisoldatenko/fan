import sys

test_case = 1
FILE = sys.stdin
# FILE = open('sample.in')
test_cases = int(FILE.readline().strip())

for _ in range(1, test_cases+1):
    r, *s = map(int, FILE.readline().strip().split())
    ss = sorted(s, reverse=True)
    total_diff = 0
    middle = ss[len(ss)//2]
    for i in range(len(ss)):
        if i == len(ss)//2:
            continue
        total_diff += abs(ss[i] - middle)
    print(total_diff)
