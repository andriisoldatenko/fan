import json

FILE = open("input.txt")

def compare(x, y):
    while len(x) > 0:
        ll = x.pop(0)
        rr = y.pop(0)
    if isinstance(ll, int) and isinstance(rr, int) and ll < rr:
        print("in the right order")
        break


def main(file):
    lines = [line.strip() for line in file if line != "\n"]
    index = 1
    results = []
    for left, right in [lines[i:i + 2] for i in range(0, len(lines), 2)]:
        left_p, right_p = json.loads(left), json.loads(right)
        while len(left_p) > 0:
            ll = left_p.pop(0)
            rr = right_p.pop(0)
            if isinstance(ll, int) and isinstance(rr, int) and ll < rr:
                print("in the right order")
                break

            if isinstance(ll, list)

# start = time.time()
print(main(FILE))
# end = time.time()
# print(end - start)
