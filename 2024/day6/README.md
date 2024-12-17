# Day 6: Guard Gallivant

This solution implements a simulation of a guard's patrol pattern in a manufacturing lab, with two different challenges:

## Part 1: Guard Path Tracking

The first part simulates a guard's movement through a lab following strict patrol rules:
1. If there's an obstacle in front, turn 90 degrees right
2. If the path is clear, move forward one step

The solution uses the following approach:
- Loads the grid from input file, representing:
  - `^` as the guard's starting position
  - `#` as obstacles
  - `.` as empty spaces
- Tracks the guard's movement until they exit the grid
- Marks each visited position with 'X'
- Returns the total number of spaces the guard visited

Key components:
- `Grid` struct: Manages the lab layout and tracks visited positions
- `Guard` struct: Maintains the guard's current position and direction
- `Direction` type: Handles the four cardinal directions and rotation logic

## Part 2: Loop Detection

The second part analyzes how placing additional obstacles affects the guard's patrol pattern. It:
1. Tries placing a block at each possible position in the grid
2. For each placement, simulates the guard's movement
3. Detects if the guard's path forms a loop (visits the same position and direction twice)
4. Counts how many different block placements create loops

The solution uses:
- A map to track visited states (position + direction combinations)
- Grid copying to test each block placement independently
- Early termination when the guard exits the grid

## Implementation Details

The code uses several helper types and functions:
- `XY`: Represents 2D coordinates
- `Direction`: Enum for Up, Right, Down, Left with rotation logic
- `Guard`: Combines position and direction with movement methods
- Grid operations for getting/setting positions and marking traveled spots

The solution prioritizes clean separation of concerns and efficient grid operations while maintaining readability.
