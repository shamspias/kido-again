### Notes

**Approach:**
I converted the list into a set and compared their lengths. If the set’s length is smaller, it means duplicates exist. This works because building a set from a list in Python takes **O(n)** time, and comparing lengths is **O(1)**.

**What went wrong:**
Initially, I overthought the problem and considered more complex checks. But no specific implementation errors occurred once I tried the set approach.

**What finally worked:**
Using `len(nums) != len(set(nums))` provided a clean and efficient solution.

```python
from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        return len(nums) != len(set(nums))
```