# Day 7: Bridge Repair

## Problem Overview
In this puzzle, you're helping engineers repair a rope bridge by solving calibration equations. Young elephants have stolen the operators from these equations, and your task is to determine which test values could possibly be produced by placing different combinations of operators between the numbers.

### Part 1
In the first part, you work with two types of operators:
- Addition (+)
- Multiplication (*)

Key rules:
- Operators are evaluated left-to-right (no operator precedence)
- Numbers must be used in the given order (no rearrangement allowed)

For example:
```
190: 10 19      # Valid: 10 * 19 = 190
3267: 81 40 27  # Valid: 81 + 40 * 27 or 81 * 40 + 27 = 3267
292: 11 6 16 20 # Valid: 11 + 6 * 16 + 20 = 292
```

The solution sums up all test values from equations that can be made true using these operators.

### Part 2
Part 2 introduces a third operator:
- Concatenation (||) - combines digits from left and right (e.g., 15 || 6 = 156)

Additional examples become valid with concatenation:
```
156: 15 6       # Valid: 15 || 6 = 156
7290: 6 8 6 15  # Valid: 6 * 8 || 6 * 15 = 7290
192: 17 8 14    # Valid: 17 || 8 + 14 = 192
```

### Implementation Details

The solution uses several key components:

1. **Data Structure**
   - `Equation` struct stores:
     - TestValue: the target value to achieve
     - Numbers: slice of integers to be combined

2. **Key Functions**
   - `cartesianProduct`: Generates all possible combinations of operations efficiently
   - `IsTrueEquation`: Tests if a valid equation can be formed using the allowed operations
   - `solvePart1`: Processes equations with only + and * operators
   - `solvePart2`: Processes equations with +, *, and || operators

3. **Optimization**
   - Short-circuit evaluation when intermediate results exceed target value
   - Efficient cartesian product generation without materializing full list
   - Single backing array for permutations to minimize memory usage

## Results
- Part 1: 1298103531759
- Part 2: 140575048428831

## Usage
Run the program with:
```bash
go run main.go
```

The program expects input in the format:
```
target: num1 num2 num3 ...
```
where target is the value to achieve and num1, num2, etc. are the numbers to be combined.

Each line represents a calibration equation, and the program will determine which equations can be made true using the allowed operators for each part.