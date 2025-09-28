# Kido Again 🥲

Hi, I’m a human (sometimes).  
Once upon a time, I could code without crying. Then autocomplete, AI, and corporate slavery made me lazy.  
Now I’m here again, forcing myself to solve problems one by one.  

This repo is my **rehab center** for coding.  
Expect bad code, too many languages, and sarcastic notes.  

---

## ⚡ Platforms I Fight On
- [LeetCode](https://leetcode.com) — the daily dose of ego destruction  
- [HackerRank](https://www.hackerrank.com) — green checkmarks for corporate slaves  
- Others (because pain has no boundaries)  

---

## 📂 Structure
Every problem lives in its own folder. Inside you’ll find:  

```

problems/<platform>/<id-or-name>/
├── link.md      # the link to the problem
├── notes.md     # short thoughts, mistakes, fixes
├── python/      # solution in Python
├── rust/        # solution in Rust
├── cpp/         # solution in C++
└── java/        # solution in Java

```

Example:

```

problems/leetcode/0001-two-sum/
link.md
notes.md
python/solution.py
rust/solution.rs

````

---

## 🛠 Languages I Use
- Python 🐍 — quick and lazy  
- Rust 🦀 — slow and painful  
- C++ 💣 — because suffering builds character  
- Java ☕ — the corporate overlord’s choice  

---

## 📝 Workflow

### Option 1: Manual
1. Pick a problem.  
2. Make a folder: `problems/<platform>/<id-name>/`  
3. Add `link.md` with the problem URL.  
4. Add `notes.md` with my random thoughts / how I solved it / what went wrong.  
5. Throw in solutions in whatever language I feel like.  
6. Commit with a sarcastic message:  
   ```bash
   git commit -m "feat: solved two-sum after crying in rust"
   ```

### Option 2: With Script (Recommended)

Use the helper script `new_problem.sh` to auto-generate the structure.

```bash
./new_problem.sh <platform> <id> <name-with-dashes> <url>
```

Example:

```bash
./new_problem.sh leetcode 0001 two-sum https://leetcode.com/problems/two-sum/
```

This creates:

```
problems/leetcode/0001-two-sum/
    link.md
    notes.md
    python/
    rust/
    cpp/
    java/
```

Then just add your code inside the language folders.

---

## 🤡 Why "Kido Again"?

Because this is me acting like a coding kid again,
except with more coffee, less patience, and questionable life choices.

---

## 📜 License

Public domain of suffering.
Use it to laugh at my code, or to feel better about yours.
