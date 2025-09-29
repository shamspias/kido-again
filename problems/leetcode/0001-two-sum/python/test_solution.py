import pytest
import json
from solution import Solution


def load_test_cases(filename="../test_cases.json"):
    with open(filename) as f:
        return json.load(f)


solution = Solution()


@pytest.mark.parametrize("case", load_test_cases())
def test_two_sum(case):
    nums, target, expected = case["nums"], case["target"], case["expected"]
    result = solution.twoSum(nums, target)
    assert result == expected, f"Input: {nums}, target={target} | Got {result}, Expected {expected}"
