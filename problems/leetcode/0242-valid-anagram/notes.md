## 1. Sorting Approach

```python
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return sorted(s) == sorted(t)
```

**Why it works:**

* Sorting puts both strings in lexicographical order, so if they are anagrams, they’ll become identical.

**Why it’s bad:**

* Sorting takes **O(n log n)** time, which is slower than necessary.
* Works fine for small strings, but not optimal for very large inputs.

**Time Complexity:**

* `sorted(s)` → O(n log n)
* Comparison → O(n)
* **Total = O(n log n)**

---

## 2. Map with Deque

```python
from collections import defaultdict, deque

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        map_s = defaultdict(deque)
        for idx, char in enumerate(s):
            map_s[char].append(idx)

        for ch in t:
            if ch in map_s and map_s[ch]:
                map_s[ch].popleft()
                if not map_s[ch]:  
                    del map_s[ch]
            else:
                return False

        return not map_s
```

**Why it works:**

* Tracks positions of each character in `s` using a deque, and removes matches while scanning `t`.

**Why it’s bad:**

* Overengineered, adds extra bookkeeping (indices, deques).
* Not better than sorting or counting.
* Still **O(n)** but with higher constant factors (memory + deque operations).

**Time Complexity:**

* Building map → O(n)
* Checking/removing → O(m)
* **Total = O(n + m)**
  (but inefficient in practice).

---

## 3. Counter Approach (Best for Unicode)

```python
from collections import Counter

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return Counter(s) == Counter(t)
```

**Why it works:**

* `Counter` counts frequency of each character.
* Comparing two counters checks if both have identical frequency distributions.
* Works for **Unicode** and any character set.

**Why it’s good:**

* Very clean and Pythonic.
* Handles all characters, not just `a–z`.

**Time Complexity:**

* Building counters → O(n + m)
* Comparing → O(k), where k = number of unique characters.
* **Total = O(n)** (if k is small compared to n).

---

## 4. Fixed-Size Array (Fastest for Lowercase Only)

```python
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        counts = [0] * 26
        for ch1, ch2 in zip(s, t):
            counts[ord(ch1) - ord('a')] += 1
            counts[ord(ch2) - ord('a')] -= 1

        return all(c == 0 for c in counts)
```

**Why it works:**

* Uses an array of size 26 (for `a–z`) to track counts.
* Adding while scanning `s` and subtracting while scanning `t` keeps balance.

**Why it’s limited:**

* Only works if input is guaranteed lowercase English letters.
* Breaks for Unicode or symbols.

**Time Complexity:**

* Counting → O(n)
* Checking → O(26) ≈ O(1)
* **Total = O(n)** (fastest possible).

---

**My simple mind**

* **Sorting:** Easy but **O(n log n)** → ❌ too slow for big input.
* **Deque map:** Works but overkill → ❌ unnecessary overhead.
* **Counter:** Clean, general, Unicode-safe → ✅ best practical solution.
* **Fixed array:** Fastest for `a–z` → ✅ but not general.
