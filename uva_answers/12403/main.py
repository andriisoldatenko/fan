import sys

FILE = sys.stdin

n = input()
total_amount = 0

for line in FILE:
    if line.startswith('donate'):
        _, amount = line.strip().split()
        total_amount += int(amount)
    else:
        print(total_amount)
