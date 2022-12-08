FILE = open("input.txt")


def column(matrix, i):
    return [row[i] for row in matrix]


def scenic_score(vec, target):
    pos = 0
    for item in vec:
        if target > item:
            pos += 1
        elif target <= item:
            pos += 1
            return pos
    return pos


def main(file):
    A = []
    results = []
    for line in file:
        A.append(list(map(int, list(line.rstrip()))))
    N = len(A)
    for i in range(0, N):
        for j in range(0, N):
            # 1, 1
            # 30373
            # 25512
            # 65332
            # 33549
            # 35390

            v1, v2 = A[i][:j], A[i][j + 1:]
            h1, h2 = column(A, j)[:i], column(A, j)[i + 1:]

            total = (scenic_score(list(reversed(v1)), A[i][j]) *
                     scenic_score(v2, A[i][j]) *
                     scenic_score(list(reversed(h1)), A[i][j]) *
                     scenic_score(h2, A[i][j]))
            results.append(total)

    return max(results)

import time
start = time.time()
print(main(FILE))
end = time.time()
print(end - start)
