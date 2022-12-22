import operator
import re
import time
from dataclasses import dataclass
from functools import reduce
from typing import List, Callable

FILE = open("input.txt")


def divide_chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i + n]


@dataclass
class Monkey:
    items: List[int]
    op: str
    num1: int
    num2: int
    divisible: int
    inspected_items: int = 0


OP = {"+": operator.add, "*": operator.mul}


def largest_prime_factor(n):
    i = 2
    while i * i <= n:
        if n % i:
            i += 1
        else:
            n //= i
    return n


def main(file):
    lines = [line.rstrip() for line in file]
    monkeys = []
    r = list(divide_chunks(lines, 7))
    for rr in r:
        num1 = int(re.findall(r"(\d+)", rr[4])[0])
        num2 = int(re.findall(r"(\d+)", rr[5])[0])
        divisible = int(re.findall(r"(\d+)", rr[3])[0])
        monkeys.append(Monkey(
            items=list(map(int, re.findall(r"(\d+)", rr[1]))),
            op=re.findall(r"(new = old [\*\-\+\/] (\d+|\w+))", rr[2])[0][0],
            num1=num1,
            num2=num2,
            divisible=divisible,
        ))
    factor = reduce(operator.mul, [m.divisible for m in monkeys])
    round = 1
    for _ in range(10000):
        for i in range(len(monkeys)):
            # print(monkeys[i].items)
            while len(monkeys[i].items) != 0:
                item = monkeys[i].items.pop()
                monkeys[i].inspected_items += 1
                op, second = monkeys[i].op.split()[-2:]
                if second == "old":
                    new = OP[op](item, item) % factor
                else:
                    new = OP[op](item, int(second))

                # Monkey gets bored with item. Worry level is divided by 3 to 500.
                # new = int(new / 3)
                m = monkeys[i]
                if new % m.divisible == 0:
                    new_monkey_index = m.num1
                else:
                    new_monkey_index = m.num2
                monkeys[new_monkey_index].items.append(new)
        round += 1
    res = []
    for m in monkeys:
        print(m.inspected_items)
        res.append(m.inspected_items)
    return sorted(res, reverse=True)[:2]


# start = time.time()
a, b = main(FILE)
print(a * b)
# end = time.time()
# print(end - start)
