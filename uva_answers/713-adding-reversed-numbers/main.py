import sys


FILE = sys.stdin
test_cases = int(input())


def reverse_num(num):
    return int(''.join(list(reversed(str(num)))))


for test_case in range(test_cases):
    a, b = map(int, FILE.readline().strip().split())
    print(reverse_num(reverse_num(a) + reverse_num(b)))
