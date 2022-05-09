from typing import List


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        extra = 0
        N = len(digits) - 1
        ones = [0] * N
        ones.append(1)
        i = N
        while i >= 0:
            curr = digits[i] + extra + ones[i]
            if curr >= 10:
                extra = 1
                digits[i] = curr % 10
            else:
                extra = 0
                digits[i] = curr
            i = i - 1
        if extra:
            digits.insert(0, extra)
        return digits
