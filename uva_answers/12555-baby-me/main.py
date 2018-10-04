#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys
import re

test_case = 1
FILE = sys.stdin
test_cases = int(input())

for test_case in range(1, test_cases+1):
    chars = re.findall(r'\d+', FILE.readline())
    total = int(chars[0]) * 50
    if len(chars) == 2:
        total += int(chars[1]) * 5
    t_100 = (total/100)
    t_int = int(t_100) if t_100.is_integer() else t_100
    print('Case {}: {}'.format(test_case, t_int))
