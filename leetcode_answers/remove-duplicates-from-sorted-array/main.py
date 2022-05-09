from typing import List

class Solution:
    """
    [0,0,1,1,1,2,2,3,3,4]
    """
    def removeDuplicates(self, nums: List[int]) -> int:
        prev = -1000
        N = len(nums)
        i = 0
        result = 0
        while i<N:
            curr = nums.pop(0)
            if curr != prev:
                result += 1
                nums.append(curr)
            i += 1
            prev = curr
        return result

