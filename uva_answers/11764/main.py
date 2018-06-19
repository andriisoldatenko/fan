import sys

FILE = sys.stdin
# FILE = open('sample.in')
test_case = 1
total_test_cases = range(1, int(input())+1)

for tc in total_test_cases:
    num_walls = range(int(input()))
    walls = list(map(int, FILE.readline().strip().split()))
    high_jumps = 0
    low_jumps = 0
    if len(walls) == 1:
        print('Case {}: {} {}'.format(tc, high_jumps, low_jumps))
    else:
        for index, w in enumerate(walls[1:], start=1):
            if walls[index] > walls[index-1]:
                high_jumps += 1
            elif walls[index] < walls[index-1]:
                low_jumps += 1
        print('Case {}: {} {}'.format(tc, high_jumps, low_jumps))
