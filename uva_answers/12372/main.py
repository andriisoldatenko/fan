import sys

test_case = 1
FILE = sys.stdin

n = input()

for line in FILE:
    l, w, h = map(int, line.strip().split())
    output = 'good'
    if l > 20 or w > 20 or h > 20:
        output = 'bad'
    print('Case {}: {}'.format(test_case, output))
    test_case +=1