import sys

test_case = 1
FILE = sys.stdin

while True:
    n = int(FILE.readline().strip())
    if n == 0:
        break
    events = list(map(int, FILE.readline().strip().split()))
    zeros = events.count(0)
    print('Case %s: %s' % (test_case, len(events) - 2 * zeros))
    test_case +=1
