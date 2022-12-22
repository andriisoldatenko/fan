import time

FILE = open("input.txt")


def print_m(A, B):
    for i in range(len(A)):
        for j in range(len(A[0])):
            print(A[i][j], end="")
        print(end="    ")
        for j in range(len(B[0])):
            print(B[i][j], end="")
        print()


def need_move(h_i, h_j, t_i, t_j):
    # h_i, h_j 4, 4
    # t_i, t_j = 4, 3
    if abs(h_i - t_i) > 1 or abs(h_j - t_j) > 1:
        return True
    return False


def column(matrix, i):
    return [row[i] for row in matrix]


def move(A, B, d, h_i, h_j, t_i, t_j):
    if d == "R":
        A[h_i][h_j + 1] = "H"
        A[h_i][h_j] = "."
        if need_move(h_i, h_j+1, t_i, t_j):
            A[h_i][h_j+1-1] = "T"
            B[h_i][h_j+1-1] = "#"
            A[t_i][t_j] = "."
            t_i, t_j = h_i - 1 + 1, h_j
        return h_i, h_j + 1, t_i, t_j
    elif d == "U":
        A[h_i - 1][h_j] = "H"
        A[h_i][h_j] = "."
        if need_move(h_i - 1, h_j, t_i, t_j):
            A[h_i - 1 + 1][h_j] = "T"
            B[h_i - 1 + 1][h_j] = "#"
            A[t_i][t_j] = "."
            t_i, t_j = h_i - 1 + 1, h_j
        return h_i - 1, h_j, t_i, t_j
    elif d == "L":
        A[h_i][h_j-1] = "H"
        A[h_i][h_j] = "."
        if need_move(h_i, h_j-1, t_i, t_j):
            A[h_i][h_j-1+1] = "T"
            B[h_i][h_j-1+1] = "#"
            A[t_i][t_j] = "."
            t_i, t_j = h_i, h_j - 1 + 1
        return h_i, h_j - 1, t_i, t_j
    elif d == "D":
        A[h_i+1][h_j] = "H"
        A[h_i][h_j] = "."
        if need_move(h_i+1, h_j, t_i, t_j):
            A[h_i+1-1][h_j] = "T"
            B[h_i+1-1][h_j] = "#"
            A[t_i][t_j] = "."
            t_i, t_j = h_i+1-1, h_j
        return h_i+1, h_j, t_i, t_j


def main(file):
    size = 1000
    w, h = size, size
    A = [["." for x in range(w)] for y in range(h)]
    B = [["." for x in range(w)] for y in range(h)]
    results = set()
    h_i, h_j = size//2, size//2
    t_i, t_j = size//2, size//2
    A[h_i][h_j] = "H"
    A[t_i][t_j] = "T"
    B[t_i][t_j] = "#"
    # print_m(A, B)
    for line in file:
        direction, step = line.rstrip().split()
        step_ = int(step)
        # print(f"{direction * 10}")
        index = 0
        for _ in range(step_):
            print(index)
            results.add((t_i, t_j))
            h_i, h_j, t_i, t_j = move(A, B, direction, h_i, h_j, t_i, t_j)
            # print_m(A, B)
            index += 1
    # print(len(results))
    print()
    # print_m(A, B)
    r = 0
    for i in B:
        r += i.count("#")
    print(r)
    return


start = time.time()
print(main(FILE))
end = time.time()
print(end - start)
