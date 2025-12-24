# advent-challenges
advent challenges solved in python and go

## Python Setup

### Install the package (for proper imports)
```bash
cd adventpy
uv pip install -e .
```

## Running Python Solutions

### Via terminal
```bash
cd adventpy
# Run via run.py script
python run.py 01 --part sol1 --data sample

# Or run as module
python -m puzzles.day_01.solution
```

### Via VS Code Debugger

The project includes debugger configurations in `.vscode/launch.json`:

1. **Debug: Run Day 01 (via run.py)** - Debug through run.py with sample data
2. **Debug: Run Day 01 (via run.py) - Input** - Debug through run.py with input1 data
3. **Debug: Day 01 Solution (direct)** - Debug solution.py directly
4. **Debug: Day 01 Solution (as module)** - Debug as a Python module
5. **Debug: Current Python File** - Debug the currently open file

**To use the debugger:**
- Set breakpoints in your solution files (e.g., `solution.py`)
- Select a debug configuration from the VS Code debug panel (F5)
- The debugger will stop at breakpoints even when called from `run.py`

**Key features:**
- `PYTHONPATH` is automatically set to include `adventpy/`
- Working directory is set to `adventpy/` for proper path resolution
- `justMyCode: false` allows debugging into imported modules
