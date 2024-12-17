# Advent of Code Solutions

This repository contains my solutions for [Advent of Code](https://adventofcode.com/), an annual coding challenge that runs from December 1st to December 25th.

## What is Advent of Code?

Advent of Code is a series of programming puzzles released daily during December, created by Eric Wastl. Each day presents two related puzzles that can be solved using any programming language. The puzzles cover various computer science concepts, algorithms, and data structures, becoming progressively more challenging as the month goes on.

Key features:
- One puzzle released each day at midnight EST (UTC-5)
- Each day has two parts, with the second part unlocked after completing the first
- Input data is unique for each participant
- Solutions can be written in any programming language
- Global leaderboard for fastest solutions
- Private leaderboards for friends or organizations

## Setting Up Your Cookie

To automatically fetch puzzle inputs, you'll need to set up your Advent of Code session cookie:

1. Log in to [Advent of Code](https://adventofcode.com/)
2. Open your browser's developer tools (usually F12 or right-click -> Inspect)
3. Go to the "Application" or "Storage" tab
4. Under "Cookies", find the cookie named `session`
5. Copy the value of this cookie
6. Create a `.env` file in the root directory of this project
7. Add your cookie value:
   ```
   AOC_COOKIE=your_cookie_value_here
   ```

**Note**: Keep your session cookie private and never commit it to version control.

## Using the Justfile

This project uses [just](https://github.com/casey/just) as a command runner. Here are the available commands:

### Generate a New Day's Solution

```bash
just gen YEAR DAY
```

This command:
- Creates a new directory for the specified year and day
- Copies the Go template files from `templates/go/`
- Downloads the input data using your AOC cookie
- Opens the solution in Windsurf IDE

Example:
```bash
just gen 2023 1
```

### Run a Solution

```bash
just run YEAR DAY
```

This command:
- Changes to the specified year/day directory
- Runs the Go solution

Example:
```bash
just run 2023 1
```

## Project Structure

```
.
├── YEAR/
│   └── dayN/
│       ├── main.go
│       └── input.txt
├── templates/
│   └── go/
├── utils/
├── .env
├── justfile
└── README.md
