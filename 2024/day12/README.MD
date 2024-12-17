# Day 12: Garden Groups

## Problem Overview
The problem involves calculating the total cost of fencing required for different regions in a garden plot. Each region is represented by connected plots containing the same type of plant (denoted by the same letter). The plots are connected horizontally or vertically.

## Part 1: Perimeter-Based Fencing Cost
In Part 1, the cost of fencing for each region is calculated by multiplying:
- Area: The number of plots in the region
- Perimeter: The number of sides that don't touch another plot of the same type

### Solution
The solution for Part 1 uses a depth-first search (DFS) approach:

1. A visited array keeps track of processed plots
2. For each unvisited plot:
   - Calculate the area by counting connected plots of the same type
   - Calculate the perimeter by checking adjacent cells
   - Multiply area × perimeter for the region's cost
3. Sum up all region costs for the final answer

Key function: `calculateRegion()` performs the DFS and calculates both area and perimeter simultaneously.

## Part 2: Straight Sections Fencing Cost
Part 2 introduces a "bulk discount" where instead of using the perimeter, we count the number of distinct sides a region has. A side counts as one unit regardless of its length. This changes how we calculate the cost:
- Area: Still the number of plots in the region
- Sides: The number of distinct straight sections of fence, regardless of length

### Solution
The solution for Part 2 uses a breadth-first search (BFS) approach with a more sophisticated side counting mechanism:

1. First identifies all cells in a region using BFS
2. Creates a padded grid for easier edge handling
3. Counts unique sides by:
   - Checking each cell's edges
   - Using a counting system to avoid double-counting shared edges
   - Tracking horizontal and vertical edges separately

Key function: `calculateRegionSides()` handles the BFS and complex side counting logic.