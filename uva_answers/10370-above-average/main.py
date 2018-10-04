import statistics
import sys

FILE = sys.stdin
tc = int(input())

for line in FILE:
    nums = list(map(int, line.strip().split()))
    n, *arr = nums
    avg = statistics.mean(arr)
    print('{0:.3f}%'.format(round(len([1 for x in arr if x > avg]) / n * 100, 3)))
