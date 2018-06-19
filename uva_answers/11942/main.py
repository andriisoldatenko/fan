import sys

FILE = sys.stdin
# FILE = open('sample.in')
# inpt = FILE.readline().strip()

test_case = 1
test_cases = range(1, int(input())+1)

print('Lumberjacks:')

for tc in test_cases:
    l = list(map(int, FILE.readline().strip().split()))
    if all((l[i] <= l[i+1])for i in range(len(l)-1)) or all(
            (l[i] >= l[i+1])for i in range(len(l)-1)):
        print('Ordered')
    else:
        print('Unordered')
