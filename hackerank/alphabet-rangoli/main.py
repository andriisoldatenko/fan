# ----c----
# --c-b-c--
# c-b-a-b-c
# --c-b-c--
# ----c----

import string
import pprint


def print_rangoli(n):
    result = []
    alpha = string.ascii_lowercase
    for i in range(n):
        s = "-".join(alpha[i:n])
        result.append(((s[::-1] + s[1:]).center(4 * n - 3, "-")))
    return result


t = print_rangoli(22)
print('\n'.join(t[:0:-1]+t))

