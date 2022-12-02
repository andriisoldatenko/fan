import sys
import heapq

# FILE = sys.stdin
FILE = open('input.txt')


curr_max = 0
total = 0
results = list()
for line in FILE:
    cl_line = line.strip()
    if line != '\n':
        total += int(line)
    else:
        results.append(total)
        total = 0
print(sum(heapq.nlargest(3, results)))
