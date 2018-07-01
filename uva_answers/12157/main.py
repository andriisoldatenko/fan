import sys

FILE = sys.stdin
# FILE = open('sample.in')
# inpt = FILE.readline().strip()

test_case = 1
test_cases = range(1, int(input())+1)

for tc in test_cases:
    n = int(input())
    costs = list(map(int, FILE.readline().strip().split()))
    total_costs = 0
    m = sum([(c//30 + 1)*10 for c in costs])
    j = sum([(c//60 + 1)*15 for c in costs])

    if m > j:
        print('Case {}: Juice {}'.format(tc, j))
    elif m < j:
        print('Case {}: Mile {}'.format(tc, m))
    else:
        print('Case {}: Mile Juice {}'.format(tc, m))
