# The Challenge
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false
 

Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.

# My Approach

## Step 1: Implicit edge cases

- If the provided string had an odd length then its invalid

- The first appearance of a bracket type must be the opening brace ie '[', '{', or '('

## Step 2: Tease out requirements
`Open brackets must be closed by the same type of brackets.` - Pairs of brackets need to be tracked 

`Open brackets must be closed in the correct order`- I need to remmeber 

## Step 3: Construct an ETL pipeline

