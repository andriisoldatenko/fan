import time

FILE = open("input.txt")

cycles = {
    40,
    80,
    120,
    160,
    200,
    240
}


def main(file):
    X = 1
    cycle = 0

    crt = ""
    for line in file:

        cl_line = line.rstrip()
        try:
            op, value = cl_line.split()
        except:
            op, value = "noop", "0"

        if op == "noop":
            if X - 1 <= cycle % 40 <= X + 1:
                crt += "*"
            else:
                crt += " "
            cycle += 1
            if cycle in cycles:
                print(crt)
                crt = ""

        elif op == "addx":
            for _ in range(2):
                if X - 1 <= cycle % 40 <= X + 1:
                    crt += "*"
                else:
                    crt += " "
                cycle += 1
                if cycle in cycles:
                    print(crt)
                    crt = ""

            X += int(value)
        else:
            pass
        # print(crt)
    # print(s)


# start = time.time()
print(main(FILE))
# end = time.time()
# print(end - start)
