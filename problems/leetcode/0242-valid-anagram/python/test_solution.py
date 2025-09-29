import pytest
import json
from solution import Solution


def load_test_cases(filename="../test_cases.json"):
    with open(filename) as f:
        return json.load(f)


solution = Solution()


@pytest.mark.parametrize("case", load_test_cases())
def test_is_anagram(case):
    s, t, expected = case["s"], case["t"], case["expected"]
    result = solution.isAnagram(s, t)
    assert result == expected, f"Input: s={s}, t={t} | Got {result}, Expected {expected}"
