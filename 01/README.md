# Day 1
For all of these save the input you get as `input.txt`.

## Part 1

This is a fairly simple scanner for the input written in Go

```sh
go run main.go
```

## Part 2

This is a bit funky. It has an LR(1) parser using GNU Bison and Flex to generate the scanner and parser,
and is written in C. However, it has a bug related to input such as `oneight` which should be parsed as
`1 8` instead of `1`. The fix is cursed, run `dirty.sh` and find replace all numbers such as `one` to
the digit and the last letter, i.e: `one` to `1e`, this is stupid, but might work.


```sh
make # builds the stuff
bash dirty.sh # runs the cursed sed commands
./part2 # generates the output
```
