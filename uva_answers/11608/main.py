import sys

FILE = sys.stdin
# FILE = open('sample.in')
test_case = 1

while True:
    s = int(FILE.readline().strip())
    if s < 0:
        break
    problems_created = [s] + list(map(int, FILE.readline().strip().split()))[:-1]
    problems_required = list(map(int, FILE.readline().strip().split()))
    balance = 0
    print('Case {}:'.format(test_case))
    for c, r in zip(problems_created, problems_required):
        if (c + balance) >= r:
            print('No problem! :D')
            balance += (c - r)
        else:
            print('No problem. :(')
            balance += c
    test_case += 1
