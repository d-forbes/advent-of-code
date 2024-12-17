set dotenv-load

gen YEAR DAY:
  #!/bin/bash
  echo "aoc cookie is $AOC_COOKIE"
  cp -R templates/go/ {{YEAR}}/day{{DAY}}
  curl --cookie "session=$AOC_COOKIE" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o {{YEAR}}/day{{DAY}}/input.txt

  windsurf {{YEAR}}/day{{DAY}}

run YEAR DAY:
  #!/bin/bash
  cd {{YEAR}}/day{{DAY}}/
  go run main.go