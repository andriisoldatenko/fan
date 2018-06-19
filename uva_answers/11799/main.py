import sys

FILE = sys.stdin
# FILE = open('sample.in')
# inpt = FILE.readline().strip()

test_case = 1
test_cases = range(1, int(input())+1)

for tc in test_cases:
    students, *speed = list(map(int, FILE.readline().strip().split()))
    print('Case {}: {}'.format(tc, max(speed)))
