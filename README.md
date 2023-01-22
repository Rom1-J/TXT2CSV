# TXT2CSV

[![Github release action](https://github.com/Rom1-J/TXT2CSV/workflows/Release/badge.svg)](https://github.com/Rom1-J/TXT2CSV/actions?query=workflow%3ARelease)
[![Github commit action](https://github.com/Rom1-J/TXT2CSV/workflows/Building/badge.svg)](https://github.com/Rom1-J/TXT2CSV/actions?query=workflow%3AGo)

Script to convert large `.txt` files (or any other format) to `.csv` via a regular expression.

---

## Installation 

### From Sources

```bash
$ git clone https://github.com/Rom1-J/TXT2CSV
$ cd TXT2CSV
$ make build  # assuming you already have go installed on your system
```

Then you can find the executable inside `dist/` directory.

### From Builds

- [Download the latest release](https://github.com/Rom1-J/TXT2CSV/releases/latest) compatible with your system 

---

## Usage

```bash
$ ./txt2csv -h
# Usage of ./main:
#   -input string
#         Input file
#   -output string
#         Output file (default "output.csv")
#   -regex string
#         Regex to use
#   -threads int
#         Number of threads to use (default 12)
```

---

## Examples

```bash
$ time ./main -input=example/input.txt -regex="(?P<uuid_a>(?:[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12})):(?P<random>(?:\w|\s|\:)+):(?P<uuid_b>(?:[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}))" -threads=48 -output=example/result.csv
# CSV header: [uuid_a random uuid_b garbage]
# Regex: (?P<uuid_a>(?:[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12})):(?P<random>(?:\w|\s|\:)+):(?P<uuid_b>(?:[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}))
# Threads: 48
# Done!
# ./main -input=example/input.txt  -threads=48 -output=example/result.csv  0.06s user 0.00s system 450% cpu 0.015 total

$ cd example
$ python verify.py
# Test passed!
```

---

## Performances tests

Specs:
+ CPU: `Intel i7-9750H (12) @ 4.500GHz`
+ Disk: `NVMe`
+ Memory: `32GB`

Sample:
+ Size: `~890MB`
+ Lines: `15,271,670`
+ Regex: `(?P<value_a>.*):(?P<value_b>[\w.-]+@[\w.-]+):(?P<value_c>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}):(?P<value_d>.*)`

Runs:
+ `135.09s user 3.39s system 1080% cpu 12.820 total`
+ `136.42s user 3.61s system 1087% cpu 12.879 total`
+ `136.58s user 3.48s system 1083% cpu 12.927 total`
