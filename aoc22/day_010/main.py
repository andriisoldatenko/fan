import time

FILE = open("input.txt")

cycles = {
    20, 
    60, 
    100, 
    140, 
    180, 
    220
}

def main(file):
    X = 1
    cycle = 0
    sum = 0
    results = []
    for line in file:
        cl_line = line.rstrip()
        try:
            op, value = cl_line.split()
        except:
            op, value = "noop", "0"
        
        if op == "noop":
            # print(cycle, X)
            cycle += 1
            if cycle in cycles:
                print(X, cycle, X * cycle)

                # results.append((X, cycle, X * cycle))
                sum += (X * cycle)
        elif op == "addx":
            for _ in range(2):
                # print(cycle, X)
                cycle += 1
                if cycle in cycles:
                    print(X, cycle, X * cycle)
                    sum += (X * cycle)
            X += int(value)
        else:
            pass
        # if cycle in cycles:
        #     results.append((X, cycle, X * cycle))
            # sum += (X * cycle)
    return sum


# start = time.time()
print(main(FILE))
# end = time.time()
# print(end - start)
