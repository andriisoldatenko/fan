import pprint
import sys

FILE = sys.stdin
# FILE = open('sample.in')

test_case = 1
while True:
    n = int(FILE.readline().strip())
    if n == 0:
        break
    nums = list(map(int, FILE.readline().strip().split()))
    avg = sum(nums)//len(nums)
    results = sum(abs(avg - x) for x in nums)
    print("Set #{}".format(test_case))
    print("The minimum number of moves is {}.".format(results//2))
    print()
    test_case += 1
