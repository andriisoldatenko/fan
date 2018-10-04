
def get_nth_ugly_number(n):
    ugly = [0 for _ in range(n)]
    ugly[1] = 0

    i2 = [2 * x for x in range(n)]
    i3 = [3 * x for x in range(n)]
    i5 = [5 * x for x in range(n)]


print('The 1500\'th ugly number is {}.'.format(i))
