#!/usr/bin/env bash

# Simple script to create a new problem folder structure
# Usage: ./new_problem.sh <platform> <id> <name-with-dashes> <url>
# Example: ./new_problem.sh leetcode 0001 two-sum https://leetcode.com/problems/two-sum/

set -e

if [ "$#" -lt 4 ]; then
    echo "Usage: $0 <platform> <id> <name-with-dashes> <url>"
    echo "Example: $0 leetcode 0001 two-sum https://leetcode.com/problems/two-sum/"
    exit 1
fi

PLATFORM=$1
PROBLEM_ID=$2
PROBLEM_NAME=$3
URL=$4

FOLDER_NAME="${PROBLEM_ID}-${PROBLEM_NAME}"
BASE_PATH="problems/${PLATFORM}/${FOLDER_NAME}"

# Create directories
mkdir -p "${BASE_PATH}"
mkdir -p "${BASE_PATH}/python" "${BASE_PATH}/rust" "${BASE_PATH}/cpp" "${BASE_PATH}/java"

# Create link.md
cat > "${BASE_PATH}/link.md" <<EOL
# Problem Link

${URL}
EOL

# Create notes.md
cat > "${BASE_PATH}/notes.md" <<EOL
# Notes

- Approach:
- What went wrong:
- What finally worked:
EOL

# Create test_cases.json
cat > "${BASE_PATH}/test_cases.json" <<EOL
{
  "samples": [
    {
      "input": "TODO: sample input",
      "output": "TODO: expected output"
    }
  ]
}
EOL

echo "✅ Created new problem at: ${BASE_PATH}"
