#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'hourglassSum' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY arr as parameter.
#


def hourglassSum(arr):
    print(arr)
    print("----")
    slice_x = 3
    slice_y = 3

    width = len(arr[0])
    height = len(arr)
    max_sum = 0
    for i in range(0, height - slice_y + 1):  # 4
        for j in range(0, width - slice_x + 1):  # 4
            total_sum = 0
            for a in range(i, i + slice_y):
                row_sum = 0
                for b in range(j, j + slice_x):
                    # element with index arr[1][0], arr[1][2]
                    row_sum += arr[a][b]
                total_sum += row_sum
            total_sum = total_sum - arr[i+1][j] - arr[i+1][j+2]
            print(total_sum)
            max_sum = max(max_sum, total_sum)
    return max_sum


if __name__ == "__main__":
    fptr = sys.stdout

    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))
    # print(arr)
    result = hourglassSum(arr)

    fptr.write(str(result) + "\n")

    fptr.close()
