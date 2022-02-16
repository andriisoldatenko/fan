from typing import List
from collections import Counter


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        return Counter(nums).most_common()[:-2:-1][0][0]
