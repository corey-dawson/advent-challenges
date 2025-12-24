from enum import Enum
from pathlib import Path
from typing import Final
from argparse import ArgumentParser
from adventpy.utils import find_root_path, init_display

class Puzzle(Enum):
    DAY = "01"
    PROBLEM = "01"

data_dir: Final[str] = f"day_{Puzzle.DAY.value}"
root_dir: Final[Path] = find_root_path()
data_dir: Final[Path] = root_dir.joinpath("data", data_dir)

def main(input_file: str) -> int:

    # setup data to read and read in data
    data_path: Final[Path] = data_dir.joinpath(input_file)
    with open(data_path, "r") as file:
        data = file.read()

    # split into 2 lists
    l1, l2 = [], []
    for row in data.split("\n"):
        chars = row.split(" ")
        l1.append(chars[0])
        l2.append(chars[-1]) 

    # sort lists
    l1.sort()
    l2.sort()

    # iterate and calculate distance
    diffs = []
    for n1, n2 in zip(l1, l2):
        diffs.append(abs(int(n1) - int(n2)))
    
    return sum(diffs)


if __name__ == "__main__":
    parser = ArgumentParser(description="Run via specified input file.")
    parser.add_argument(
        "input_file", 
        type=str,
        nargs='?',
        default="sample",
        help="The input file (e.g. sample, input1, input2)"
    )
    args = parser.parse_args()
    init_display(Puzzle.DAY.value, Puzzle.PROBLEM.value, args.input_file)
    answer = main(args.input_file)
    print(f"Answer: {answer}")
