import os
from pathlib import Path
from typing import Final

ROOT_DIR_NM: Final[str] = "advent-challenges"

def find_root_path() -> Path:
    current_dir = Path(os.getcwd())
    dirs = os.listdir(current_dir)

    while ROOT_DIR_NM not in dirs:
        current_dir = current_dir.parent
        dirs = os.listdir(current_dir)
 
    root_dir = current_dir.joinpath(ROOT_DIR_NM)
    return root_dir.absolute()

def init_display(day: str, problem: str, input_file: str):
    print("")
    print(f"Solution: Day {day} Problem {problem}") 
    print("==============================")
    print(f"Data File: {input_file}")
    print("")
