import sys

FILE = sys.stdin

for line in FILE:
    if line.strip() == '0':
        break
    n = int(line)
    nums = map(int, FILE.readline().strip().split())
    print(' '.join(map(str, sorted(nums))))
