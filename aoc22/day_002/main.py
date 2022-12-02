import sys
import heapq

# FILE = sys.stdin
FILE = open("input.txt")

# A Rock
# B paper
# C clipe
mapping = {"X": "A", "Y": "B", "Z": "C"}
score = {"A": 1, "B": 2, "C": 3}


def main(file):
    result = 0
    for line in file:
        cl_line = line.strip()
        left, right_ = cl_line.split()
        right = mapping[right_]
        if left == right:
            result += 3 + score[left]
        elif right == "A" and left == "C":
            result += 6 + 1
        elif right == "B" and left == "A":
            result += 6 + 2
        elif right == "C" and left == "B":
            result += 6 + 3
        else:
            result += score[right]

    return result


print(main(FILE))
