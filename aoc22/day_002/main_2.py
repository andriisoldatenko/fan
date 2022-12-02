import sys
import heapq

# FILE = sys.stdin
FILE = open("input.txt")

# A Rock
# X lose,
# Y in a draw
# Z win

score = {"A": 1, "B": 2, "C": 3}


def main(file):
    result = 0
    for line in file:
        cl_line = line.strip()
        left, right_ = cl_line.split()
        if right_ == "Y":
            result += 3 + score[left]
        elif right_ == "X":
            if left == "A":
                result += 3
            else:
                result += score[left] - 1
        elif right_ == "Z":
            if left == "C":
                result += 1 + 6
            else:
                result += score[left] + 1 + 6

    return result


print(main(FILE))
