import sys

FILE = sys.stdin
# FILE = open('sample.in')


for line in FILE:
    num_students = int(line.strip())
    if num_students == 0:
        break
    expenses = []
    for _ in range(num_students):
        d, c = map(int, FILE.readline().strip().split('.'))
        expenses.append(d * 100 + c)
    s = sum(expenses)
    N = len(expenses)
    low_avg = s // N
    high_avg = low_avg + (1 if (s % N) else 0)

    sum_above = 0
    for e in expenses:
        if e > high_avg:
            sum_above += e - high_avg

    sum_below = 0
    for e in expenses:
        if e < low_avg:
            sum_below += low_avg - e

    used_sum = max(sum_above, sum_below)
    print('${}.{:02d}'.format(used_sum // 100, used_sum % 100))
