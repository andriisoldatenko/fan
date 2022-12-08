FILE = open("input.txt")


def column(matrix, i):
    return [row[i] for row in matrix]


def scenic_score(vec, target):
    pos = 0
    for item in vec:
        if item > target:
            pos += 1
        elif item == target:
            pos += 1
            return pos


def main(file):
    A = []
    for line in file:
        A.append(list(map(int, list(line.rstrip()))))
    count = 0
    N = len(A)
    for i in range(0, N):
        for j in range(0, N):
            if i == 4 and j == 2:
                1 == 1
            v1, v2 = A[i][:j], A[i][j + 1:]
            h1, h2 = column(A, j)[:i], column(A, j)[i + 1:]
            if i in [0, N - 1] or j in [0, N - 1]:
                count += 1
                continue

            if (A[i][j] > max(v1) or
                    A[i][j] > max(v2) or
                    A[i][j] > max(h1) or
                    A[i][j] > max(h2)
            ):
                print(i, j)
                count += 1

    return count


import time

start = time.time()
print(main(FILE))
end = time.time()
print(end - start)
