import time

FILE = open("input.txt")
# W, H = 6, 5
# X, Y = 4, 0

W, H = 100, 100
X, Y = 50, 50


# def compute():



def print_m(numbers):
    A = [["." for x in range(W)] for y in range(H)]
    for index, num in enumerate(numbers):
        if A[num[0]][num[1]] == ".":
            if index == 0:
                A[num[0]][num[1]] = "H"
            else:
                A[num[0]][num[1]] = f"{index}"
    for i in range(len(A)):
        for j in range(len(A[0])):
            print(A[i][j], end="")
        print()

def print_r(numbers):
    A = [["." for x in range(W)] for y in range(H)]
    for i in range(len(A)):
        for j in range(len(A[0])):
            if (i, j) in numbers:
                print("#", end="")
            else:
                print(".", end="")
        print()



def column(matrix, i):
    return [row[i] for row in matrix]


def move(d, numbers):
    if d == "R":
        new_numbers = []
        first_i, first_j = numbers[0][0], numbers[0][1]
        h_i, h_j = first_i, first_j + 1
        new_numbers.append((h_i, h_j))

        for i in numbers[1:]:
            t_i, t_j = i[0], i[1]
            if abs(h_j - t_j) > 1 and abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1

                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1
            elif abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1
                t_j = h_j
            elif abs(h_j - t_j) > 1:
                t_i = h_i
                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1
            h_i, h_j = t_i, t_j
            new_numbers.append((t_i, t_j))
        return new_numbers
    elif d == "U":
        new_numbers = []
        first_i, first_j = numbers[0][0], numbers[0][1]
        h_i, h_j = first_i - 1, first_j
        new_numbers.append((h_i, h_j))

        for i in numbers[1:]:
            t_i, t_j = i[0], i[1]
            if abs(h_j - t_j) > 1 and abs(h_i - t_i) > 1:
                # t_i, t_j = h_i+1, h_j - 1
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1

                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1

            elif abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1
                t_j = h_j
            elif abs(h_j - t_j) > 1:
                t_i = h_i
                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1
            h_i, h_j = t_i, t_j
            new_numbers.append((t_i, t_j))
        return new_numbers
    elif d == "L":
        new_numbers = []
        first_i, first_j = numbers[0][0], numbers[0][1]
        h_i, h_j = first_i, first_j - 1
        new_numbers.append((h_i, h_j))

        for i in numbers[1:]:
            t_i, t_j = i[0], i[1]
            if abs(h_j - t_j) > 1 and abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1

                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1

            elif abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1
                t_j = h_j
            elif abs(h_j - t_j) > 1:
                t_i = h_i
                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1
            h_i, h_j = t_i, t_j
            new_numbers.append((t_i, t_j))
        return new_numbers
    elif d == "D":
        new_numbers = []
        first_i, first_j = numbers[0][0], numbers[0][1]
        h_i, h_j = first_i + 1, first_j
        new_numbers.append((h_i, h_j))
        for i in numbers[1:]:
            t_i, t_j = i[0], i[1]
            if abs(h_j - t_j) > 1 and abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1

                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1
            elif abs(h_i - t_i) > 1:
                if h_i > t_i:
                    t_i = h_i - 1
                else:
                    t_i = h_i + 1
                t_j = h_j
            elif abs(h_j - t_j) > 1:
                if h_j > t_j:
                    t_j = h_j-1
                else:
                    t_j = h_j+1
                t_i = h_i
            h_i, h_j = t_i, t_j
            new_numbers.append((t_i, t_j))
        return new_numbers


def main(file):
    # w, h = 26, 21
    A = [["." for x in range(W)] for y in range(H)]
    B = [["." for x in range(W)] for y in range(H)]
    start_i, start_j = 0, 0
    result = set()
    numbers = [
        (X, Y) for _ in range(10)
    ]
    # print_m(A, B)
    for line in file:
        direction, step = line.rstrip().split()
        step_ = int(step)
        # print(f"{direction * 10}")
        index = 0
        for _ in range(step_):
            numbers = move(direction, numbers)
            # print_m(numbers)
            print(f"{direction * 10}")
            result.add(numbers[-1])
            index += 1

    print_r(result)
    return len(result)


start = time.time()
print(main(FILE))
end = time.time()
print(end - start)
