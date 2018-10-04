import sys
from pprint import pprint

# FILE = sys.stdin
FILE = open('sample.in')

AVAILABLE_COMMANDS = {'I', 'C', 'L', 'V', 'H', 'K', 'F', 'S', 'X'}

def create_matrix(m, n):
    return [['0' for _ in range(m)]
             for _ in range(n)]

def draws_vertical_segment(x, y1, y2, c, img):
    for i in range(len(img[0])):
        for j in range(len(img)):
            if j == x-1 and y1-1 <= i <= y2-1:
                img[i][j] = c
    return img


def draws_horizontal_segment(x1, x2, y, c, img):
    for i in range(len(img[0])):
        for j in range(len(img)):
            if i == y-1 and x1-1 <= j <= x2-1:
                img[i][j] = c
    return img

def draws_filled_rectangle(x1, y1, x2, y2, c):
    pass

image = None
for line in FILE:
    cmd = line.strip().split()
    if cmd[0] not in AVAILABLE_COMMANDS:
        continue
    if cmd[0] == 'I':
        image = create_matrix(int(cmd[1]), int(cmd[2]))
        # pprint(image)
    if cmd[0] == 'C':
        image = create_matrix(len(image), len(image[0]))
        # pprint(image)
    if cmd[0] == 'L':
        image[int(cmd[1])][int(cmd[2])] = cmd[3]
    if cmd[0] == 'S':
        print(cmd[1])
    if cmd[0] == 'H':
        image = draws_horizontal_segment(int(cmd[1]), int(cmd[2]), int(cmd[3]),
                                         cmd[4], image)
        pprint(image)
    if cmd[0] == 'X':
        break
    print(cmd)
