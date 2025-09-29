# Notes

- **Approach:**
  - First, I tried the brute-force solution with two nested loops:
    ```python
    for i, num in enumerate(nums):
        for j, other_num in enumerate(nums):
            if j == 0:
                continue
            if num + other_num == target:
                return [i, j]
    return []
    ```
    This works but is **O(n²)** because of the double loop.

  - Then I thought about improving it. Sorting came to mind first, since sorting is **O(n log n)**, and after that maybe use two pointers or binary search. But that only reduces part of the problem and still doesn’t give the best lookup efficiency.

  - The real bottleneck is the lookup for the "other number" (the complement). If I can speed that up, I can avoid the second loop.

- **What went wrong:**
  - With brute force, the solution was too slow (**O(n²)**).
  - With sorting, it got a bit better (**O(n log n)**), but I realized it complicates the problem since I also need the original indices, not just the numbers.

- **What finally worked:**
  - Using a hash map (dictionary in Python) for constant-time lookups.
  - For each number, calculate the complement (`target - num`).  
    - If the complement is already in the map, return the pair of indices.  
    - If not, store the number with its index in the map.  
  - This reduces the time complexity to **O(n)**:
    ```python
    lookup = {}
    for i, num in enumerate(nums):
        other_num = target - num
        if other_num in lookup:
            return [lookup[other_num], i]
        lookup[num] = i
    return []
    ```