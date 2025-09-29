from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        lookup = {}
        for i, num in enumerate(nums):
            complement = target - num
            if complement in lookup:
                return [lookup[complement], i]
            lookup[num] = i
        return []

# if __name__ == "__main__":
#     solution = Solution()
#     print(solution.twoSum([3, 2, 4], 6))
