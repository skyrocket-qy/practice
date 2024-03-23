# Explanination

## Implementation 1

Using slice to store all inpus, and check substring in each stored input

### Complexity

- Space: O(total characters in all input strings)
  
#### Injest

- Time: O(1)

#### Appearance

- Time: O(total characters in all input strings)

## Implementation 2

- Using Prefix tree with token as node to store all inputs -> reduce duplicate token storage
- Using tokenMap to store token's node for each searching entry point
- Storing extra information(start, count, end) in node for counting purpose

### Complexity

- Space: O(num of all duplicate tokens)

#### Injest

- Time: O(num of input tokens)

#### Appearance

- Time: O(num of input tokens)
