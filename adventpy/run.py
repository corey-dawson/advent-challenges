import importlib
import sys
import argparse
from pathlib import Path
from enum import Enum
from puzzles import find_root_path, init_display
from typing import Final


def main():
    parser = argparse.ArgumentParser(description="Run a specific puzzle solution.")
    parser.add_argument(
        "day", 
        type=str,
        default="01",
        help="The day number (e.g., 01, 02)"
    )
    parser.add_argument(
        "--part",
        type=str,
        default="sol1",
        choices=["sol1", "sol2"],
        help="The solution function to run (sol1, sol2, default: sol1)"
    )
    parser.add_argument(
        "--data",
        type=str,
        default="sample",
        choices=["sample", "input1"],
        help="The data file to use (default: sample)"
    )
    args = parser.parse_args()

    # Normalize day format (e.g., "1" -> "01")
    day_num = args.day.zfill(2)
    module_path = f"puzzles.day_{day_num}.solution"

    # setup paths for job
    proj_root_dir: Final[Path] = find_root_path()

    class Paths(Enum):
        value: Path
        ROOT = proj_root_dir
        DATA = proj_root_dir.joinpath("data", day_num)

    data_file: Final[Path] = Paths.DATA.value.joinpath(args.data)

    init_display(day_num, data_file)

    # try:
    #     # Dynamically import the module
    #     problem_module = importlib.import_module(module_path)
        
    #     # Get the solution function
    #     if hasattr(problem_module, args.part):
    #         solution_func = getattr(problem_module, args.part)
    #         print(f"--- Running Day {day_num}, {args.part} ---")
    #         solution_func(args.data)
    #     else:
    #         print(f"Error: Function '{args.part}' not found in {module_path}")
    #         available = [attr for attr in dir(problem_module) if not attr.startswith('_')]
    #         print(f"Available functions: {available}")
    # except ModuleNotFoundError as e:
    #     print(f"Error: Could not find solution for day '{day_num}' at {module_path}")
    #     print(f"Details: {e}")
    # except Exception as e:
    #     print(f"Error running solution: {e}")
    #     import traceback
    #     traceback.print_exc()


if __name__ == "__main__":
    main()