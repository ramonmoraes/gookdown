# GookDown

Partially create longs texts using markdown


# Example
Given the summary file at `summary.md`:
```md
# 1. Intro
- [Intro](./readme/intro.md)
# 2. Development
- [Development](./readme/development.md)
```
GookDown will compile the content of `Intro` and `Development` file into the new `README.md` file

## Options

To change the input and output files, pass the it's flags
```
gookdown -input="./foo" -output="./bar"
```

# Development

## Runing

`$ go run main.go`
