import sys

# FILE = sys.stdin
# FILE = open('sample.in')

def fopen(filename, default=sys.stdin):
    """
    This function helps you do not see errors during upload solutions
    because you forget to switch back to sys.stdin and also you can easily
    debug your code with ipdb or another python debuggers
    """
    try:
        f = open(filename)
    except FileNotFoundError:
        f = default
    return f

# In this algorithm, a variation

# of selection sort, we bring the largest pancake not yet sorted
# to the top with one flip; take it down to its final position with one more flip;
# and repeat this process for the remaining pancakes.
def is_sorted(arr):
    for i in range(len(arr) - 1):
        if arr[i] >= arr[i+1]:
            return False
    return True

def partial_rotate(arr, start, end):
    first_part = arr[start:end+1][::-1]
    second_part = arr[end+1:]
    return first_part + second_part


def flip(arr):
    flips = []
    if is_sorted(arr):
        flips.append(0)
        return flips
    start = 0
    for i in reversed(range(len(arr))):
        index_max = arr.index(max(arr[:i+1]))
        if index_max == i:
            continue
        elif index_max == start:
            arr = partial_rotate(arr, 0, i)
            flips.append(len(arr) - i)
        else:
            arr = partial_rotate(arr, 0, index_max)
            flips.append(len(arr) - index_max)
            arr = partial_rotate(arr, 0, i)
            flips.append(len(arr)-i)

    if is_sorted(arr):
        flips.append(0)

    return flips

for line in fopen('sample1.in'):
    pancakes = list(map(int, line.strip().split()))
    print(' '.join(map(str, pancakes)))
    print(' '.join(map(str, flip(pancakes))))
