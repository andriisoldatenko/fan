#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the jumpingOnClouds function below.
def jumpingOnClouds(c):
    i = 0
    current_position = 0
    total_jumps = 0
    while i < len(c)-1:
        if i+1 == len(c)-1:
            total_jumps += 1
            break
        # print(i)
        if c[i+2] == 0:
            i += 2
        else:
            i += 1
        total_jumps += 1
    return total_jumps

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    c = list(map(int, input().rstrip().split()))

    result = jumpingOnClouds(c)

    fptr.write(str(result) + '\n')

    fptr.close()

