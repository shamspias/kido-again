import pytest
import json
from solution import Solution


def load_test_cases(filename="../test_cases.json"):
    with open(filename) as f:
        return json.load(f)


solution = Solution()


@pytest.mark.parametrize("case", load_test_cases())
def test_contains_duplicate(case):
    nums, expected = case["nums"], case["expected"]
    result = solution.containsDuplicate(nums)
    assert result == expected, f"Input: nums={nums} | Got {result}, Expected {expected}"
